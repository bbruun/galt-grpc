package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/bbruun/galt/server"
)

var LogLevels []string = []string{"debug", "info", "warn", "error", "fatal", "trace"}

type LoggerStruct struct {
	Text      string
	Level     int
	LogLevel  string
	Directory string
	Filename  string
	Compress  bool
	Rotate    time.Duration
	Format    string
}

func (l *LoggerStruct) Init() {
	if l.Directory == "" {
		l.Directory = "logs"
	}

	if _, err := os.Stat(l.Directory); err != nil {
		fmt.Printf("- dig not find log directory: %s\n", l.Directory)

		if err := os.Mkdir(l.Directory, 0755); err != nil {
			fmt.Printf("Failed to access logs directory: %v\n", err)
			os.Exit(1)
		}
	}

	logFile := fmt.Sprintf("%s/%s", l.Directory, l.Filename)
	if _, err := os.Stat(logFile); err != nil {
		if _, err := os.Create(logFile); err != nil {
			fmt.Printf("Failed to create log file: %v\n", err)
			os.Exit(1)
		}
	}

	l.setLogLevel()
}

func (l *LoggerStruct) setLogLevel() {
	if l.LogLevel == "" {
		l.LogLevel = "info"
		l.Level = 1 // Default loglevel
	} else {
		for i, c := range LogLevels {
			if l.LogLevel == c {
				l.Level = i
			}
		}
	}
}

func (l *LoggerStruct) getTime() string {
	t := time.Now().UTC()
	ret := t.Format(server.Configuration.Logger.Format)
	fmt.Printf("%s time: %s\n", server.Configuration.Logger.Format, ret)
	return ret
}

func (l *LoggerStruct) Debug(text ...interface{}) {
	tmp := ""
	for x := range text {
		tmp += fmt.Sprintf("%v\n", text[x])
	}
	fmt.Printf("%s DEBUG %s", l.getTime(), tmp)
}
