package builder

import (
	"gitee.com/candy44/dbcore-sdk/pbfiles"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

// ParamBuilder 参数构建器
type ParamBuilder struct {
	param map[string]interface{}
}

func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{
		param: make(map[string]interface{}),
	}
}

func (p *ParamBuilder) Add(name string, value interface{}) *ParamBuilder {
	p.param[name] = value
	return p
}

func (p *ParamBuilder) Build() *pbfiles.SimpleParams {
	paramStruct, err := structpb.NewStruct(p.param)
	if err != nil {
		log.Println(err)
	}

	return &pbfiles.SimpleParams{Params: paramStruct}
}
