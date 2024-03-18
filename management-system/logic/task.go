package logic

import "time"

func StartTimer(f func()) {
	go func() {
		for {
			// 取得当前时间
			t := time.Now()
			// 计算下一个8:30的时间
			next := time.Date(t.Year(), t.Month(), t.Day(), 8, 30, 0, 0, t.Location())
			if t.After(next) {
				// 如果当前时间已经过了今天的8:30，就计算下一个8:30
				next = next.Add(24 * time.Hour)
			}
			// 计算时间差
			duration := next.Sub(t)
			// 设置定时器
			t2 := time.NewTicker(duration)
			<-t2.C
			// 到达8:30，执行任务
			f()
		}
	}()
}
