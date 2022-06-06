package debug

import "time"

// Wait 阻塞式等待，但不占用CPU.
// 只适合debug使用哦！
func Wait() {
	for true {
		time.Sleep(time.Hour * 24 * 30 * 365)
	}
}

// Sleep 休眠多少s. 精度到毫秒
func Sleep(seconds float64) {
	time.Sleep(time.Millisecond * time.Duration(seconds*1000))
}
