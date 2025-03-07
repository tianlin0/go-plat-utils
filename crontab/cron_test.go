package crontab_test

import (
	"github.com/tianlin0/go-plat-utils/crontab"
	"testing"
)
import "fmt"

func TestCrontab(t *testing.T) {

	crontab.StartJobs(map[string]func(){
		"*/2 * * * *": func() {
			fmt.Println("1分钟1")
		},
		"0 02 17 * *": func() {
			fmt.Println("定点1")
		},
	})

	select {}
}
func TestCrontabLockKey(t *testing.T) {
	crontab.StartJobs(map[string]func(){
		"* * * * *": func() {
			fmt.Println("1分钟1")
		},
	})
	crontab.StartJobs(map[string]func(){
		"* * * * *": func() {
			fmt.Println("1分钟2")
		},
	})

	select {}
}
