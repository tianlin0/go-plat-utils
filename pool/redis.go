package pool

import "github.com/gomodule/redigo/redis"

type RedisPoolFactory[P redis.Pool] struct {
}

func (p *RedisPoolFactory[P]) NewPool(param *redis.Pool) (*RedisPool, error) {
	res := &redis.Pool{}
	return &RedisPool{
		p: res,
	}, nil
}

type RedisPool struct {
	p *redis.Pool
}

func (p *RedisPool) Get() (*RedisObject, error) {
	return &RedisObject{
		conn: p.p.Get(),
	}, nil
}
func (p *RedisPool) Put(v *RedisObject) error {

	return nil
}

type RedisObject struct {
	conn redis.Conn
}
