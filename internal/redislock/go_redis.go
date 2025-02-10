package lredis

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var (
	oneSleep = 100 * time.Millisecond
)

// redisLock
type redisLock struct {
	redisClient *redis.Client

	key         string
	value       string
	expiration  time.Duration
	tries       int // 重试次数
	mu          sync.Mutex
	isLocked    bool
	count       int
	renewCtx    context.Context
	renewCancel context.CancelFunc
}

// NewRedisLock 新的锁
func NewRedisLock(redisClient *redis.Client, key string, expiration time.Duration) (*redisLock, error) {
	if redisClient == nil {
		return nil, fmt.Errorf("redis client is nil")
	}

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	v := base64.StdEncoding.EncodeToString(b)

	//redis时间不能太短，避免大量的redis操作
	if expiration < DefaultExpireTime {
		expiration = DefaultExpireTime
	}

	times := int(expiration/oneSleep) + 1

	return &redisLock{
		redisClient: redisClient,
		key:         getLockerKeyName(key),
		value:       v,
		tries:       times,
		expiration:  expiration,
	}, nil
}

// Lock 上锁
func (l *redisLock) Lock(ctx context.Context) (bool, error) {
	return l.lockContext(ctx, l.tries)
}

// UnLock 解锁
func (l *redisLock) UnLock(ctx context.Context) (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if !l.isLocked {
		return true, nil
	}

	l.count--
	if l.count > 0 {
		return true, nil
	}

	unlockScript := `
        local key = KEYS[1]
        local identifier = ARGV[1]

        if redis.call('GET', key) == identifier then
            return redis.call('DEL', key)
        else
            return 0
        end
    `
	result, err := l.redisClient.Eval(ctx, unlockScript, []string{l.key}, l.value).Result()
	if err != nil {
		return false, err
	}
	retSuccess := result.(int64) == 1
	if retSuccess {
		l.mu.Lock()
		defer l.mu.Unlock()
		l.isLocked = false
		l.count = 0
		l.renewCancel()
	}

	return retSuccess, nil
}

// TryLock 尝试加锁
func (l *redisLock) TryLock(ctx context.Context) (bool, error) {
	return l.lockContext(ctx, 1)
}

func (l *redisLock) lockContext(ctx context.Context, tries int) (bool, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.isLocked {
		l.count++
		return true, nil
	}

	var timer *time.Timer
	for i := 0; i < tries; i++ {
		if i != 0 {
			if timer == nil {
				timer = time.NewTimer(oneSleep)
			} else {
				timer.Reset(oneSleep)
			}

			select {
			case <-ctx.Done():
				timer.Stop()
				// Exit early if the context is done.
				return false, ctx.Err()
			case <-timer.C:
				// Fall-through when the delay timer completes.
			}
		}

		ok, err := l.redisClient.SetNX(ctx, l.key, l.value, l.expiration).Result()
		if i == tries-1 && err != nil { //最后一次才会返回错误
			return false, err
		}
		if ok {
			l.isLocked = true
			l.count++
			l.renewCtx, l.renewCancel = context.WithCancel(context.Background())
			return true, nil
		}
	}

	return false, nil
}
