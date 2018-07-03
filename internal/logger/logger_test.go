package logger

import (
	"os"
	"sync"
	"testing"

	logging "github.com/op/go-logging"
)

func TestLogger(t *testing.T) {
	var log = logging.MustGetLogger("example")
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend2Leveled := logging.AddModuleLevel(backend2)
	logging.SetBackend(backend2Leveled, backend2Formatter)
	log.Info("abc hello world")

	l := NewMyLog("log.txt")
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
	b.Run("myBufferedLog", func(b *testing.B) {
		benchmarkLogger(b, NewMyBufferedLog("log.txt"))
	})
	b.Run("nanolog", func(b *testing.B) {
		benchmarkLogger(b, NewNanoLog("log.txt"))
	})
	b.Run("stdlog", func(b *testing.B) {
		benchmarkLogger(b, NewStdLog("log.txt"))
	})
	b.Run("logrus", func(b *testing.B) {
		benchmarkLogger(b, NewLogrusLog("log.txt"))
	})
	b.Run("opLoggingLog", func(b *testing.B) {
		benchmarkLogger(b, NewOpLoggingLog("log.txt"))
	})
	b.Run("zaplog", func(b *testing.B) {
		benchmarkLogger(b, NewZapLog("log.txt"))
	})
	b.Run("mylog", func(b *testing.B) {
		benchmarkLogger(b, NewMyLog("log.txt"))
	})
	b.Run("seelog", func(b *testing.B) {
		benchmarkLogger(b, NewSeeLog("log.txt"))
	})

}
