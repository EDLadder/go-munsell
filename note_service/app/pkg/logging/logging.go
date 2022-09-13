package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0755)

	if err != nil || os.IsExist(err) {
		panic("can't create log dir. no configured logging to files")
	} else {
		logFile, err := os.OpenFile("logs/service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(fmt.Sprintf("[Message]: %s", err))
		}

		l.SetOutput(ioutil.Discard)

		l.AddHook(&writer.Hook{
			Writer:    io.Writer(logFile),
			LogLevels: logrus.AllLevels,
		})

		l.AddHook(&writer.Hook{
			Writer:    os.Stdout,
			LogLevels: logrus.AllLevels,
		})
	}

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
