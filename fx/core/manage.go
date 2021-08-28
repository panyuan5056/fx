package core

import (
	"fmt"
	"fx/core/xdb"
	"fx/core/xfile"
	"fx/core/xstrategy"
	"fx/pkg/rule"
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

type Manage struct {
}

func (st *Manage) pf(config []*xstrategy.Strategy, schemas []string) map[int]int {
	result := make(map[int]int)
	for key, value := range schemas {
		for key2, value2 := range config {
			forwards := strings.Split(value2.Forward, ",")
			for _, forward := range forwards {
				if strings.ToLower(forward) == strings.ToLower(value) {
					result[key] = key2
				}
			}
		}
	}
	return result
}

func (st *Manage) matched1(content interface{}, plan string) bool {
	if match, err := regexp.MatchString(plan, cast.ToString(content)); err == nil {
		return match
	}
	return false
}

func (st *Manage) matched2(content interface{}, plan string) bool {
	fmt.Println("matched")
	result := rule.NewRule(content, plan)
	if cast.ToString(result["status"]) == "1" {
		return true
	}
	return false
}

func (st *Manage) matched3(content interface{}, plan string) bool {
	fmt.Println("matched")
	for _, p := range strings.Split(plan, ",") {
		if cast.ToString(content) == p {
			return true
		}
	}
	return false
}

func (st *Manage) custom(content interface{}, plan string) bool {
	if cast.ToString(content) == plan {
		return true
	}
	return false
}

func (st *Manage) swith(conf *xstrategy.Strategy, name string, data []map[string]interface{}) int {
	step := 0
	for _, item := range data {
		if content, ok := item[name]; ok {
			switch conf.Category {
			//正则
			case "1":
				if ok := st.matched1(content, conf.Plan); ok {
					step += 1
				}
			//rule
			case "2":
				if ok := st.matched2(content, conf.Plan); ok {
					step += 1
				}
			//in
			case "3":
				if ok := st.matched3(content, conf.Plan); ok {
					step += 1
				}
				//eq
			default:
				if ok := st.custom(content, conf.Plan); ok {
					step += 1
				}
			}
		}
	}
	return step
}

func (st *Manage) Xrun(schemas []string, data []map[string]interface{}, dept float32) []map[string]string {
	//获取规则里的数据
	strategys := xstrategy.Find()
	xforward := st.pf(strategys, schemas)

	//循环表结构用字段比较
	result := []map[string]string{}
	for index, schema := range schemas {
		//判断比较是否存在，优先跑存在的
		conf_index := xforward[index]
		conf := strategys[conf_index]
		//获取到表里的数据
		if step := st.swith(conf, schema, data); step > 0 {
			row := map[string]string{"name": schema, "total": cast.ToString(len(data)),
				"step": cast.ToString(step), "config": conf.Name, "class1": conf.Class1, "class2": conf.Class2, "class3": conf.Class3}
			result = append(result, row)

		} else {
			for _, conf := range strategys {
				if step2 := st.swith(conf, schema, data); step2 > 0 {
					row := map[string]string{"name": schema, "total": cast.ToString(len(data)),
						"step": cast.ToString(step2), "config": conf.Name, "class1": conf.Class1, "class2": conf.Class2, "class3": conf.Class3}
					result = append(result, row)
				}
			}
		}
	}
	return result
}

func XFind(category string, config map[string]string) map[string][]map[string]string {
	var (
		tables  []string
		schemas map[string][]string
		datas   map[string][]map[string]interface{}
	)
	if category == "1" {
		tables, schemas, datas, _ = xdb.DbFind(config)
	} else {
		tables, schemas, datas, _ = xfile.FileFind(config)
	}
	result := map[string][]map[string]string{}
	manage := Manage{}
	for _, table := range tables {
		schema := schemas[table]
		data := datas[table]
		rows := manage.Xrun(schema, data, 0.8)
		result[table] = rows
	}
	return result
}
