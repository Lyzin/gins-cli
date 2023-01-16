package zaplogger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	dpath "{{.ProjectName}}/pkg/path_handler"
	"time"
)

var SugarLogger *zap.SugaredLogger

func InitLogger() {
	encoder := getEncoder()
	
	logFileName := time.Now().Format("2006-01-02")
	// 全量日志文件名，例如：logs/2023-01-16.log
	fullLogFIleName := logFileName
	
	// 错误日志文件名，例如：logs/2023-01-16.err.log
	errLogFileName := logFileName + ".err"
	
	// 全量日志的core
	fullLogCore := zapcore.NewCore(encoder, getLogWriter(fullLogFIleName), zapcore.DebugLevel)
	
	// 错误日志的core
	errLogCore := zapcore.NewCore(encoder, getLogWriter(errLogFileName), zapcore.ErrorLevel)
	
	// 合并全量日志code和错误日志core
	core := zapcore.NewTee(fullLogCore, errLogCore)
	
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func logFilePathHandler(logFileName string) string {
	currentPath, _ := os.Getwd()
	logsPath := currentPath + "/logs"
	
	if !dpath.CheckDirOrFileExists(logsPath) {
		// create log dir
		_ = os.Mkdir(logsPath, os.ModePerm)
	}
	
	logFilePath := logsPath + fmt.Sprintf("/%v.log", logFileName)
	return logFilePath
}

func getLogWriter(logFileName string) zapcore.WriteSyncer {
	logFilePath := logFilePathHandler(logFileName)
	file, _ := os.Create(logFilePath)
	return zapcore.AddSync(file)
}
