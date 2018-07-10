package logger

import (
	"sync"
	"testing"
)

// 20 个协程一起跑
func benchmarkLogger(b *testing.B, logger Logger) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < b.N; i++ {
				logger.Info("12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkLogger(b *testing.B) {
	b.Run("mylog no buffer", func(b *testing.B) {
		benchmarkLogger(b, NewMyLog("log.txt"))
	})
	b.Run("mylog buffered", func(b *testing.B) {
		benchmarkLogger(b, NewMyBufferedLog("log.txt"))
	})
	b.Run("stdlog", func(b *testing.B) {
		benchmarkLogger(b, NewStdLog("log.txt"))
	})
	b.Run("logrus", func(b *testing.B) {
		benchmarkLogger(b, NewLogrusLog("log.txt"))
	})
	b.Run("seelog", func(b *testing.B) {
		benchmarkLogger(b, NewSeeLog("log.txt"))
	})
	b.Run("seelog no buffer", func(b *testing.B) {
		benchmarkLogger(b, NewSeeLogWithoutBuffer("log.txt"))
	})
	b.Run("oplogging", func(b *testing.B) {
		benchmarkLogger(b, NewOpLoggingLog("log.txt"))
	})
	b.Run("zaplog", func(b *testing.B) {
		benchmarkLogger(b, NewZapLog("log.txt"))
	})
}
