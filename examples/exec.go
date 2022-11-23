package examples

import (
	"context"
	"fmt"
	"gitee.com/candy44/dbcore-sdk/pkg/builder"

	"google.golang.org/grpc"
	"log"
)

type UserAddResult struct {
	UserID       int `mapstructure:"user_id"`
	RowsAffected int `mapstructure:"_RowsAffected"`
}

func ExecTest() {
	// 客户端构建器
	client, _ := builder.NewClientBuilder("localhost:9999").WithOption(grpc.WithInsecure()).Build()

	//构建 参数
	paramBuilder := builder.NewParamBuilder().
		Add("username", "xiaoba").
		Add("userPass", "123456").
		Add("addTime", "2020-10-10")

	//构建API对象
	api := builder.NewApiBuilder("adduser", builder.ApiTypeExec)

	//构建一个 结果集对象----必须是 地址
	ret := &UserAddResult{}

	err := api.Invoke(context.Background(), paramBuilder, client, ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)

}
