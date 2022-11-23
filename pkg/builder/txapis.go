package builder

import (
	"context"
	"fmt"
	"gitee.com/candy44/dbcore-sdk/pbfiles"
	"github.com/mitchellh/mapstructure"

	"google.golang.org/grpc"
	"log"
)

// TxApi 事务API对象
type TxApi struct {
	ctx    context.Context
	cancel context.CancelFunc
	client pbfiles.DBService_TxClient
}

func NewTxApi(ctx context.Context, client pbfiles.DBServiceClient, opts ...grpc.CallOption) *TxApi {
	apiCtx, cancel := context.WithCancel(ctx)
	txClient, err := client.Tx(apiCtx, opts...)
	if err != nil {
		panic(err)
	}
	return &TxApi{
		ctx:    ctx,
		cancel: cancel,
		client: txClient,
	}
}

func (t *TxApi) Exec(apiName string, paramBuilder *ParamBuilder, out interface{}) error {
	// 对于exec ,如果不出错， 会返回一个map，其中key=exec ,  值是一个interface切片 ，包含了两项
	// 1、受影响的行  2 、selectkey(如果有的话)
	err := t.client.Send(&pbfiles.TxRequest{
		Api:    apiName,
		Params: paramBuilder.Build(),
		Type:   "exec",
	})
	if err != nil {
		return err
	}
	// 接收消息
	rsp, err := t.client.Recv()
	if err != nil {
		return err
	}

	if out != nil {
		if execRet, ok := rsp.Result.AsMap()["exec"]; ok { //execRet 是一个 切片 []interface{受影响的行，select}  .select 可能是nil
			if execRet.([]interface{})[1] != nil { //代表select 有值
				m := execRet.([]interface{})[1].(map[string]interface{})
				m["_RowsAffected"] = execRet.([]interface{})[0]
				return mapstructure.WeakDecode(m, out)
			} else { //没有select 情况 直接塞一个_RowsAffected 返回
				m := map[string]interface{}{"_RowsAffected": execRet.([]interface{})[0]}
				return mapstructure.WeakDecode(m, out)
			}
		}
	}
	return nil
}

func (t *TxApi) Query(apiName string, paramBuilder *ParamBuilder, out interface{}) error {
	err := t.client.Send(&pbfiles.TxRequest{Api: apiName, Params: paramBuilder.Build(), Type: "query"})
	// 对于查询，如果不出错，会返回一个map   其中key=query    值是查询结果
	if err != nil {
		return err
	}
	rsp, err := t.client.Recv()
	if err != nil {
		return err
	}
	if out != nil {
		if queryRet, ok := rsp.Result.AsMap()["query"]; ok {
			return mapstructure.WeakDecode(queryRet, out)
		} else {
			return fmt.Errorf("error query result ")
		}
	}
	return nil
}

func (t *TxApi) QueryForModel(apiName string, paramBuilder *ParamBuilder, out interface{}) error {
	err := t.client.Send(&pbfiles.TxRequest{Api: apiName, Params: paramBuilder.Build(), Type: "query"})
	// 对于查询，如果不出错，会返回一个map   其中key=query    值是查询结果
	if err != nil {
		return err
	}
	rsp, err := t.client.Recv()
	if err != nil {
		return err
	}
	if out != nil {
		if queryRet, ok := rsp.Result.AsMap()["query"]; ok {
			// queryRet 是 []interface类型。具体看 dbcore的一层转换--service.go #102
			if retForMap, ok := queryRet.([]interface{}); ok && len(retForMap) == 1 {
				return mapstructure.WeakDecode(retForMap[0], out)
			} else {
				return fmt.Errorf("error query model: no result ")
			}
		} else {
			return fmt.Errorf("error query result ")
		}
	}
	return nil
}

// Tx 模仿 gorm
func (t *TxApi) Tx(fn func(tx *TxApi) error) error {
	err := fn(t)
	if err != nil {
		log.Println("tx error:", err)
		t.cancel() //取消
		return err
	}
	return t.client.CloseSend() //协程不安全
}
