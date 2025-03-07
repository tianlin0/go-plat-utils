package utils_test

import (
	"fmt"
	"github.com/tianlin0/go-plat-utils/utils"
	"testing"
)

type AAA struct {
	Name string `json:"bName" db:"aName"`
	Age  string `json:"age"`
	Key  string `db:"key"`
	Key1 string `db:"key1" json:"bKey1"`
	Key2 string
}

func TestTag(t *testing.T) {
	mm, err := utils.GetFieldNamesByTag(AAA{}, "json", "db")
	fmt.Println(mm, err)

	mm, err = utils.GetFieldNamesByTag(AAA{}, "db")
	fmt.Println(mm, err)

	mm, err = utils.GetFieldNamesByTag(AAA{}, "json")
	fmt.Println(mm, err)

	mm, err = utils.GetFieldNamesByTag(AAA{})
	fmt.Println(mm, err)
}
