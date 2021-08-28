package rule

import (
	"fx/pkg/logging"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

func NewRule(content interface{}, rule1 string) map[string]interface{} {
	dataContext := context.NewDataContext()
	//注入初始化的结构体
	data := map[string]interface{}{
		"request":  content,
		"response": map[string]interface{}{},
	}
	dataContext.Add("Data", data)
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	err := ruleBuilder.BuildRuleFromString(rule1) //string(bs)
	if err != nil {
		logging.Error(err.Error())
	} else {
		eng := engine.NewGengine()
		err := eng.Execute(ruleBuilder, true)
		if err != nil {
			logging.Error(err)
		}
	}
	return data
}
