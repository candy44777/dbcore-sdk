package examples

import (
	"context"
	"fmt"
	"github.com/candy44777/dbcore-sdk/pkg/builder"
	"google.golang.org/grpc"
	"log"
	"time"
)

func QueryTest() {

	// 客户端构建器
	client, _ := builder.NewClientBuilder("localhost:9999").WithOption(grpc.WithInsecure()).Build()
	// 参数构建器
	paramBuilder := builder.NewParamBuilder().Add("id", 1)
	// api 构建器
	api := builder.NewApiBuilder("deptlist", builder.ApiTypeQuery)

	// 查询结果集
	depts := make([]*Dept, 0)

	//执行 和调用  API
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	err := api.Invoke(ctx, paramBuilder, client, &depts)
	if err != nil {
		log.Fatal(err)
	}
	for _, dept := range depts {
		fmt.Println(dept)
	}
}
