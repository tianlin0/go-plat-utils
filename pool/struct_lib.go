package pool

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

/**
 * 对象Pool
 */
var (
	poolRegister       = sync.Map{} //存储不同对象的pool
	currentNodeIdMutex = sync.Mutex{}
	currentNodeId      = "" //不同机器的不同节点
)

// getCurrentKey 获取不同的对象所对应的唯一key
func getCurrentKey(id string) string {
	tempId := func() string {
		if currentNodeId != "" {
			return currentNodeId
		}
		currentNodeIdMutex.Lock()
		defer currentNodeIdMutex.Unlock()
		if currentNodeId == "" {
			currentNodeId = uuid.New().String()
		}
		return currentNodeId
	}()
	typeKey := fmt.Sprintf("{%s}%s", tempId, id)
	return typeKey
}

// registerPool 注册一个pool
func registerPool[V any](po ObjectBuilder[V]) error {
	if po.ID() != po.ID() {
		return fmt.Errorf("id ret error, must be same")
	}

	typeKey := getCurrentKey(po.ID())
	if _, ok := poolRegister.Load(typeKey); !ok {
		poolRegister.Store(typeKey, &sync.Pool{
			New: func() any {
				return po.New()
			},
		})
	}
	return nil
}

// get 获取一个对象
func get[V any](po ObjectBuilder[V]) (vv V) {
	typeKey := getCurrentKey(po.ID())
	if onePool, ok := poolRegister.Load(typeKey); ok {
		obj := onePool.(*sync.Pool).Get()
		if v, ok := obj.(V); ok {
			return v
		}
	}
	return vv
}

// close 归还一个对象，需要reset一下
func close[V any](po ObjectBuilder[V], vv V) error {
	typeKey := getCurrentKey(po.ID())
	if onePool, ok := poolRegister.Load(typeKey); ok {
		//需要reset一下
		onePool.(*sync.Pool).Put(po.Reset(vv))
		return nil
	}
	return fmt.Errorf("no pool found")
}

// close 归还一个对象，需要reset一下
func closePool[V any](po ObjectBuilder[V]) {
	typeKey := getCurrentKey(po.ID())
	poolRegister.Delete(typeKey)
}
