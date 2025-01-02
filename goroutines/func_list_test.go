package goroutines_test

import (
	"fmt"
	"github.com/tianlin0/go-plat-utils/goroutines"
	"testing"
	"time"
)

func TestAsyncExecuteDataList(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i < 100; i++ {
		arr = append(arr, i+1)
	}
	num := 0
	ret, err := goroutines.AsyncExecuteDataList(1*time.Second, arr, func(key int, value int) (breakFlag bool, err error) {
		fmt.Println("key=", key, "; value=", value)
		num++
		time.Sleep(1 * time.Second)
		return true, fmt.Errorf("3333")
	})
	fmt.Println(num, ret, err)
}
