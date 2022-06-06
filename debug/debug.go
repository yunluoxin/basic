// Created by East.Zhang at 2022.06.04

// Package debug 测试时候用的包
package debug

import (
	"fmt"
	"sync"
	"time"
)

type clock struct {
	name    string
	startAt time.Time
	endAt   time.Time
}

func (c clock) duration() time.Duration {
	return c.endAt.Sub(c.startAt)
}

var rs map[string]clock
var lock sync.Mutex

func init() {
	rs = make(map[string]clock)
}

// Tick 开始计时
func Tick(name string) {
	if len(name) < 1 {
		fmt.Println("你必须传入一个非空字符串以标识这个计时器哦")
		return
	}
	lock.Lock()
	rs[name] = clock{
		name:    name,
		startAt: time.Now(),
	}
	lock.Unlock()
}

// Tok 结束计时并输出耗时
func Tok(name string) {
	lock.Lock()
	defer lock.Unlock()
	v, e := rs[name]
	if !e {
		fmt.Println("你必须先调用 debug.Tik(name)，然后才能调用 debug.Tock(name)，必须配套使用哦！")
		return
	}
	lock.Unlock()

	v.endAt = time.Now()
	duration := v.duration()
	ms := duration.Milliseconds()
	if ms > 1000 {
		fmt.Printf("🛵🛵🛵Clock[%s] cost %.3f s\n", name, duration.Seconds())
	} else if ms < 1 {
		fmt.Printf("🚀🚀🚀Clock[%s] cost %d ns\n", name, duration.Nanoseconds())
	} else {
		fmt.Printf("🚗🚗🚗Clock[%s] cost %d ms\n", name, duration.Milliseconds())
	}
	lock.Lock()
	delete(rs, name)
}

// TickTok 执行传入的方法，计算并输出耗时
func TickTok(code func()) {
	Tick("_")
	code()
	Tok("_")
}
