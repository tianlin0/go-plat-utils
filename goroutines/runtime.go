package goroutines

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/tianlin0/go-plat-utils/internal"
	"github.com/timandy/routine"
	"log"
	"runtime"
	"strings"
)

var (
	defaultAntsPool    *ants.Pool
	defaultPanicHandle func(err error, retRecover any) //全局Panic后的处理方法
	oncePool           = &internal.Once{}
)

// SetDefaultPanicHandle panic的方法
func SetDefaultPanicHandle(c func(err error, retRecover any)) {
	if c != nil {
		defaultPanicHandle = c
	}
}

// OpenRoutinePool 启动一个全局的goroutine的协程池，只会执行一次
func OpenRoutinePool(nums int) *ants.Pool {
	if defaultAntsPool == nil {
		err := oncePool.Do(func() error {
			if nums == 0 {
				nums = ants.DefaultAntsPoolSize
			}
			newPool, err := ants.NewPool(nums)
			if err == nil {
				defaultAntsPool = newPool
				return nil
			}
			return err
		})
		if err != nil {
			log.Println("OpenRoutinePool error:", err)
			return nil
		}
	}
	return defaultAntsPool
}

// CloseRoutinePool 不用了，关闭池
func CloseRoutinePool() {
	if defaultAntsPool != nil {
		if !defaultAntsPool.IsClosed() {
			defaultAntsPool.Release()
			defaultAntsPool = nil
		}
	}
}

// GoSync 同步方法
func GoSync(task func(params ...any), params ...any) {
	defer func() {
		if err := recover(); err != nil {
			//打印调用栈信息
			buf := make([]byte, 2048)
			n := runtime.Stack(buf, false)
			stackInfo := fmt.Sprintf("%s", buf[:n])
			stackInfo = strings.ReplaceAll(stackInfo, "\n", "|")
			errStr := fmt.Sprintf("panic_stack_info: %s ### %s", err, stackInfo)
			if defaultPanicHandle != nil {
				defaultPanicHandle(fmt.Errorf(errStr), err)
			} else {
				log.Println(errStr)
			}
			return
		}
	}()
	task(params...)
}

// GoAsync 异步方法
func GoAsync(task func(params ...any), params ...any) {
	fun := func() {
		func(tempParams ...interface{}) {
			GoSync(task, tempParams...)
		}(params...)
	}
	taskFun := routine.WrapTask(fun)
	if defaultAntsPool == nil {
		go taskFun.Run()
	} else {
		defaultAntsPool.Submit(taskFun.Run)
	}
}
