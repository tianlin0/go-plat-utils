package pool

type BasePool[PO PooledObject[V], V any] struct {
}

// Get 从BasePool中获取一个对象。
func (p *BasePool[PO, V]) Get() (v V, err error) {
	//return p.pool.Get(), nil
	return v, nil
}

// Put 将一个对象返回给BasePool。
func (p *BasePool[PO, V]) Put(po PO) {
	po.Reset()
	//p.pool.Put(x)
}

// Close 关闭BasePool，释放所有对象。
func (p *BasePool[PO, V]) Close() error {
	// sync.Pool没有提供直接关闭的方式，但是可以通过设置New为nil来防止Get创建新对象。
	//p.pool.New = nil
	return nil
}
