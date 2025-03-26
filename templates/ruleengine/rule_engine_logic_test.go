package ruleengine_test

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/tianlin0/go-plat-utils/templates/ruleengine"
	"regexp"
	"testing"
)

func TestCheckOneRule(t *testing.T) {
	dataMap := map[string]interface{}{
		"name": "jacky",
		"age":  decimal.NewFromInt(20),
	}

	rule := "age>18"

	ruleEngine := ruleengine.NewEngineLogic()
	ok, err := ruleEngine.RunOneRuleString(rule, dataMap)
	fmt.Println(ok, err)
}

func TestCheckOneRule22(t *testing.T) {
	ruleEngine := ruleengine.NewEngineLogic()
	ruleEngine.SetDelimitedString("{{", "}}")

	kk, err := ruleEngine.Vars("age1 + age2")

	fmt.Println(kk, err)

}

// TestCheckRuleList 判断规则是否满足条件
func TestCheckRuleList(t *testing.T) {
	var f1 float64 = 0.1
	var f2 float64 = 0.2
	fmt.Println(f1 > f2)

	d1 := decimal.NewFromFloat(0.1)
	d2 := decimal.NewFromFloat(0.2)
	// 进行加法运算
	result1 := d1.Add(d2)

	fmt.Println(result1)

	ruleList := []*ruleengine.RuleInfo{
		{
			Key:        "1",
			RuleString: "name=='jack'",
		},
		{
			Key:        "4",
			RuleString: "age1 + age2",
		},
		{
			Key:        "5",
			RuleString: "RULE_1 && (RULE_4 > 45 || RULE_4 > 30) && MulByNumber(age1, age2) > 23",
		},
	}
	ruleEngine := ruleengine.NewEngineLogic()
	ruleEngine.SetRetRulePrefix("RULE_")
	ok, err := ruleEngine.CheckLastRuleByList(ruleList, map[string]interface{}{
		"name": "jack",
		"age1": 24,
		"age2": 20,
	})
	fmt.Println(ok, err)

	result, err := ruleEngine.RunOneRuleString("5 + MulByNumber(age1,AddByNumber(age1, age2))", map[string]interface{}{
		"age1": 0.1,
		"age2": 0.2,
	})
	fmt.Println(result, err)

	ruleList = []*ruleengine.RuleInfo{
		{
			Key:        "1",
			RuleString: "name=='jack'",
		},
		{
			Key:        "2",
			RuleString: "name",
		},
		{
			Key:        "5",
			RuleString: "MulByNumber(age1, age2) > 23",
		},
	}

	ok, err = ruleEngine.CheckAllRuleList(ruleList, "&&", map[string]interface{}{
		"name": "jack",
		"age1": 24,
		"age2": 20,
	})
	fmt.Println(ok, err)
}

// TestCheckRuleList 判断规则是否满足条件
func TestCheckRuleList1(t *testing.T) {
	condTypeCustomVarPattern := `{{([^}]+)}}`
	re, err := regexp.Compile(condTypeCustomVarPattern)
	if err != nil {
		return
	}
	// 执行匹配操作
	matches := re.FindAllStringSubmatch("{{namae}}==5 && {{kkkk}} == 6", -1)
	if len(matches) > 1 {
		// 提取捕获组中的值
		name := matches[1]
		fmt.Println("提取到的 name 值为:", name)
	} else {
		fmt.Println("未找到匹配的内容")
	}
	return
}

type AA struct {
	Name string
}

func TestCheckRuleList2(t *testing.T) {
	ppMap := map[string]*AA{
		"tian1": {
			Name: "tian1",
		},
		"tian2": {
			Name: "tian2",
		},
	}
	ppList := []*AA{
		{
			Name: "tian3",
		},
		{
			Name: "tian4",
		},
	}
	a := lo.Keys(ppMap)
	fmt.Println(a)
	b := lo.Values(ppMap)
	fmt.Println(b)

	c := lo.MapKeys(ppMap, func(value *AA, key string) string {
		return key + "aa"
	})
	fmt.Println(c)

	d := lo.MapValues(ppMap, func(value *AA, key string) string {
		return value.Name + "aa"
	})
	fmt.Println(d)

	lo.ForEach(ppList, func(item *AA, index int) {
		fmt.Println(item, index)
	})

	return
}
