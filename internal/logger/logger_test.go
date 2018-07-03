package logger

import (
	"sync"
	"testing"
)

func TestLogger(t *testing.T) {
	l := NewNanoLog("log.txt")
	l.Info(1, 2, 3)
	l.Info(1, 2, 3, 4)
	t.Error()
}

func benchmarkLogger(b *testing.B, logger Logger) {
	// for i := 0; i < b.N; i++ {
	// 	logger.Info("12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	// }
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
	// b.Run("stdlog", func(b *testing.B) {
	// 	benchmarkLogger(b, NewStdLog("log.txt"))
	// })
	b.Run("seelog", func(b *testing.B) {
		benchmarkLogger(b, NewSeeLog("log.txt"))
	})
	b.Run("logrus", func(b *testing.B) {
		benchmarkLogger(b, NewLogrusLog("log.txt"))
	})
	b.Run("zaplog", func(b *testing.B) {
		benchmarkLogger(b, NewZapLog("log.txt"))
	})
	b.Run("nanolog", func(b *testing.B) {
		benchmarkLogger(b, NewNanoLog("log.txt"))
	})
}
