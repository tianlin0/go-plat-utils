package ruleengine_test

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tianlin0/go-plat-utils/templates/ruleengine"
	"testing"
)

func TestCheckOneRule(t *testing.T) {
	dataMap := map[string]interface{}{
		"name": "jacky",
		"age":  20,
	}
	rule := "name"

	ruleEngine := ruleengine.NewEngineLogic()
	ok, err := ruleEngine.RunOneRuleString(rule, dataMap)
	fmt.Println(ok, err)
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
