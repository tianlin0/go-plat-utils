package id

import (
	"fmt"
	"github.com/sony/sonyflake"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	generator *sonyflake.Sonyflake
	maxTimes  = 10
	m         sync.Mutex
)

// initGenerator 初始化,避免用init
func initGenerator() {
	if generator != nil {
		return
	}
	if maxTimes <= 0 {
		log.Panic("id generator is nil")
		return
	}

	t, _ := time.Parse("2006-01-02 15:04:05", "2025-02-14 00:00:00")

	m.Lock()
	generator = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: nil,
	})
	m.Unlock()

	if generator == nil {
		m.Lock()
		maxTimes--
		m.Unlock()
		//循环检测
		initGenerator()
	}
}

// Generator id 生成器，接受一个进制值，转换为对应进制的 id
func Generator(base int) (string, error) {
	initGenerator()

	if base < 2 || base > 64 {
		return "", fmt.Errorf("生成唯一ID失败,base参数要求在 2～64 之间")
	}

	id, err := generator.NextID()
	if err != nil {
		// 函数内部重试一次
		id, err = generator.NextID()
		if err != nil {
			return "", fmt.Errorf("生成唯一ID失败: %w", err)
		}
	}

	return strconv.FormatInt(int64(id), base), nil
}

// GeneratorBase32 generator 的 base=32 参数科里化后的方法
func GeneratorBase32() (string, error) {
	return Generator(32)
}

// getMachineID 给 sonyFlake 的 machineID 方法赋值
func getMachineID() (uint16, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return 0, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if ip == nil {
			continue
		}

		return uint16(ip[2])<<8 + uint16(ip[3]), nil
	}

	return 0, fmt.Errorf("get nil ip")
}
