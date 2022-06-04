// Created by East.Zhang at 2022.06.04

// Package debug 测试时候用的包
package debug

import (
	"fmt"
	"time"
)

type clock struct {
	name string
	startAt time.Time
	endAt time.Time
}

func (c clock) duration() time.Duration {
	return c.endAt.Sub(c.startAt)
}

var rs map[string]clock

func init() {
	rs = make(map[string]clock)
}

func Tick(name string)  {
	if len(name) < 1 {
		fmt.Println("你必须传入一个非空字符串以标识这个计时器哦")
		return
	}
	rs[name] = clock{
		name: name,
		startAt: time.Now(),
	}
}

func Tok(name string)  {
	v, e := rs[name]
	if !e {
		fmt.Println("你必须先调用 debug.Tik(name)，然后才能调用 debug.Tock(name)，必须配套使用哦！")
		return
	}
	v.endAt = time.Now()
	duration := v.duration()
	ms := duration.Milliseconds()
	if ms > 1000 {
		fmt.Printf("Clock[%s] cost %.3fs\n", name, duration.Seconds())
	} else if ms < 1 {
		fmt.Printf("Clock[%s] cost %dns\n", name, duration.Nanoseconds())
	} else {
		fmt.Printf("Clock[%s] cost %dms\n", name, duration.Milliseconds())
	}
	delete(rs, name)
}