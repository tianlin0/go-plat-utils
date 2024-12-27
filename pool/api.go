package pool

// PooledFactory 创建一个一类池的工厂
type PooledFactory[PO PooledObject[V], V any, P any] interface {
	NewPool(param P) (OnePool[PO, V], error)
}

// OnePool 一个池得实现接口
type OnePool[PO PooledObject[V], V any] interface {
	Get() (PO, error)
	Put(v PO) error
	Close() error
}

// PooledObject 需要放进池中的对象
type PooledObject[V any] interface {
	New() (V, error) //新建一个具体对象
	Reset()          //放进池时，需要进行重置
}
