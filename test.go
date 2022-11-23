package main

import (
	"github.com/candy44777/dbcore-sdk/examples"
)

type Dept struct {
	DeptId   int    `mapstructure:"depts_id"`
	DeptName string `mapstructure:"depts_name"`
}

type UserAddResult struct {
	UserID       int `mapstructure:"user_id"`
	RowsAffected int `mapstructure:"_rowsAffected"`
}

func main() {
	//conn, err := grpc.DialContext(context.Background(),
	//	"localhost:9999",
	//	grpc.WithInsecure(),
	//)
	//
	//if err != nil {
	//	log.Fatalln()
	//}
	//
	//client := pbfiles.NewDBServiceClient(conn)
	//
	////paramBuilder := builder.NewParamBuilder().Add("id", 1)
	////api := builder.NewApiBuilder("deptlist", builder.ApiTypeQuery)
	//
	////depts := make([]*Dept, 0)
	////if err := api.Invoke(context.Background(), paramBuilder, client, &depts); err != nil {
	////	log.Fatal(err)
	////}
	////for _, dept := range depts {
	////	fmt.Println(dept)
	////}
	//paramBuilder := builder.NewParamBuilder().
	//	Add("username", "xiaohua").
	//	Add("userPass", "123456").
	//	Add("addTime", "2020-10-10")
	//
	//api := builder.NewApiBuilder("adduser", builder.ApiTypeExec)
	//ret := &UserAddResult{}
	//if err := api.Invoke(context.Background(), paramBuilder, client, ret); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(ret)

	examples.QueryTest()
}
