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

// BenchmarkLogger/myBufferedLog-8         	  100000	     12633 ns/op	    2561 B/op	      40 allocs/op
// BenchmarkLogger/nanolog-8               	  100000	     17313 ns/op	    5460 B/op	     100 allocs/op
// BenchmarkLogger/stdlog-8                	  100000	     19336 ns/op	    2562 B/op	      40 allocs/op
// BenchmarkLogger/logrus-8                	   50000	     41596 ns/op	    8989 B/op	     220 allocs/op
// BenchmarkLogger/opLoggingLog-8          	   20000	     93461 ns/op	   39569 B/op	     580 allocs/op
// BenchmarkLogger/zaplog-8                	    5000	    286435 ns/op	    3539 B/op	      60 allocs/op
// BenchmarkLogger/mylog-8                 	    3000	    480025 ns/op	    4809 B/op	      60 allocs/op
// BenchmarkLogger/seelog-8                	   20000	     79896 ns/op	   12819 B/op	     200 allocs/op
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
