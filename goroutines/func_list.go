package goroutines

import (
	"math"
	"time"
)

// AsyncExecuteFunction 异步执行方法，返回是否全部执行
func AsyncExecuteFunction(timeout time.Duration, calls ...func() (bool, error)) (complete bool, errExec error) {
	if calls == nil || len(calls) == 0 {
		return true, nil
	}
	dataList := make([]func() (bool, error), 0)
	for _, call := range calls {
		dataList = append(dataList, call)
	}
	callback := func(key int, value func() (bool, error)) (bool, error) {
		return value()
	}
	return AsyncExecuteDataList(timeout, dataList, callback)
}

// AsyncExecuteDataList 异步执行数据列表，返回是否全部执行
// return: bool 是否完成循环。  error  执行过程中是否有错误
func AsyncExecuteDataList[T any](timeout time.Duration, dataList []T,
	callback func(key int, value T) (breakFlag bool, err error)) (complete bool, errExec error) {
	if dataList == nil || len(dataList) == 0 {
		return true, nil
	}
	waitGroupTemp := waitGroup(timeout)
	waitGroupTemp.add(len(dataList))

	breakDataList := false
	var errTotal error

	//如果dataList太长，这样会并发很多也不合理，所以分为二维数组会更合理一些，每50一组
	stepNum := 50
	timesFloat := float64(len(dataList)) / float64(stepNum)
	asyncTimeInt := int(math.Ceil(timesFloat)) //合并后的次数

	i := 0
	for j := 0; j < asyncTimeInt; j++ {
		for ; i < len(dataList); i++ {
			if i >= (j+1)*stepNum {
				break
			}

			if breakDataList {
				//如果循环中有跳出的指令以后，则后续的循环都直接全部完成
				waitGroupTemp.done()
				continue
			}
			GoAsync(func(params ...interface{}) {
				oneKeyTemp, ok0 := params[0].(int)
				oneValTemp, ok1 := params[1].(T)
				if ok0 && ok1 {
					breakFlag, err := callback(oneKeyTemp, oneValTemp)
					if err != nil {
						errTotal = err
					}
					if breakFlag {
						//表示需要跳出后续循环
						breakDataList = true
					}
					//如果完成了，才能关闭
					waitGroupTemp.done()
				}
			}, i, dataList[i])

		}
	}

	err := waitGroupTemp.wait()
	if err != nil {
		//超时没有完成
		return false, err
	}
	if errTotal != nil {
		return true, errTotal
	}
	return true, nil
}