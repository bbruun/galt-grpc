package logging

import "github.com/bbruun/galt/server"

var Logger *LoggerStruct

func NewLogger() *LoggerStruct {
	l := LoggerStruct{
		LogLevel:  server.Configuration.Logger.LogLevel,
		Directory: server.Configuration.Logger.Directory,
		Filename:  server.Configuration.Logger.Filename,
		Format:    server.Configuration.Logger.Format,
	}
	l.Init()
	return &l
}

func Debug(text ...interface{}) {
	Logger.Debug(text)
}
