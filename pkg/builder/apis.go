package builder

import (
	"context"
	"github.com/candy44777/dbcore-sdk/heplers"
	"github.com/candy44777/dbcore-sdk/pbfiles"
	"github.com/mitchellh/mapstructure"
)

const (
	ApiTypeQuery = iota
	ApiTypeExec
)

type ApiBuilder struct {
	name    string // api 名称
	apiType int
}

func NewApiBuilder(name string, apiType int) *ApiBuilder {
	return &ApiBuilder{
		name:    name,
		apiType: apiType,
	}
}

// Invoke 不同执行，不是事务
func (a *ApiBuilder) Invoke(ctx context.Context, paramBuilder *ParamBuilder,
	client pbfiles.DBServiceClient, out interface{}) error {
	if a.apiType == ApiTypeQuery {
		req := &pbfiles.QueryRequest{Name: a.name, Params: paramBuilder.Build()}
		rsp, err := client.Query(ctx, req)
		if err != nil {
			return err
		}
		// 如果 out 没有 传值 不做转换
		if out != nil {
			return mapstructure.WeakDecode(heplers.PbStructsToMapList(rsp.GetResult()), out)
		}
		return nil
	} else {
		req := &pbfiles.ExecRequest{Name: a.name, Params: paramBuilder.Build()}
		rsp, err := client.Exec(ctx, req)
		if err != nil {
			return err
		}
		if out != nil {
			var m map[string]interface{}
			if rsp.Select != nil {
				m = rsp.Select.AsMap()
				m["_RowsAffected"] = rsp.RowsAffected
			} else {
				m = map[string]interface{}{
					"_RowsAffected": rsp.RowsAffected,
				}
			}
			return mapstructure.WeakDecode(m, out)
		}
		return nil
	}
}
