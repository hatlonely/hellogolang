package logger

import "github.com/cihub/seelog"

func NewSeeLog(filename string) *SeeLog {
	testConfig := `<seelog>
	<outputs formatid="access">
		<buffered formatid="access" size="10000" flushperiod="100">
        	<rollingfile type="date" filename="` + filename + `" datepattern="2006-01-02-15" maxrolls="240"/>
		</buffered>
	</outputs>
    <formats>
        <format id="access" format="%Msg%n"/>
    </formats>
</seelog>`

	log, err := seelog.LoggerFromConfigAsBytes([]byte(testConfig))
	if err != nil {
		panic(err)
	}
	return &SeeLog{
		log: log,
	}
}

type SeeLog struct {
	log seelog.LoggerInterface
}

func (l *SeeLog) Info(v ...interface{}) {
	l.log.Info(v...)
}
