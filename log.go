package zlog

import (
	"flag"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var sugarLogger *zap.SugaredLogger
var sugarLoggerWithOnceCallerAdd *zap.SugaredLogger
var (
	logLevel, logPath, logFileName string
	logSize                        int
	logBackups                     int
	logAge                         int
	isCallerVisible                bool
)

func init() {
	flag.StringVar(&logLevel, "log.level", "info", "log levels:debug/info/warn/error/dpanic/panic/fatal")
	flag.StringVar(&logPath, "log.path", "/tmp", "log save path")
	flag.IntVar(&logSize, "log.size", 10,
		"MaxSize is the maximum size in megabytes of the log file before it gets rotated. It defaults to 10 megabytes.")
	flag.IntVar(&logBackups, "log.backups", 5,
		"MaxBackups is the maximum number of old log files to retain.")
	flag.IntVar(&logAge, "log.age", 7,
		"MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename.")
	flag.StringVar(&logFileName, "log.filename", "default", "log file name")
	flag.BoolVar(&isCallerVisible, "log.caller", true, "log the caller or not")

}

func InitLogger() {

	writeSyncer := getLogWriter()
	encoder := getEncoder()

	var realLogLevel zapcore.Level
	if err := realLogLevel.Set(logLevel); err != nil {
		panic(err)
	}
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), realLogLevel)

	if isCallerVisible {
		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		sugarLogger = logger.Sugar()
		sugarLoggerWithOnceCallerAdd = zap.New(core, zap.AddCaller()).Sugar()
	} else {
		sugarLogger = zap.New(core).Sugar()
		sugarLoggerWithOnceCallerAdd = sugarLogger
	}

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func randomSuffix() string {
	rand.Seed(time.Now().UnixNano())
	suffix := "_"
	for range make([]int, 5) {
		suffix += strconv.FormatInt(int64(rand.Intn(16)), 16)
	}
	return suffix
}

func getLogWriter() zapcore.WriteSyncer {

	if len(os.Args) > 0 && logFileName == "default" {
		exec := os.Args[0]

		paths := strings.Split(exec, "/")
		if len(paths) > 0 {
			logFileName = paths[len(paths)-1]
			logFileName += randomSuffix()
		}
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", logPath, logFileName),
		MaxSize:    logSize,
		MaxBackups: logBackups,
		MaxAge:     logAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args)
}
func Error(args ...interface{}) {
	sugarLogger.Error(args)
}
func Debug(args ...interface{}) {
	sugarLogger.Debug(args)
}
func Fatal(args ...interface{}) {
	sugarLogger.Fatal(args)
}
func Warn(args ...interface{}) {
	sugarLogger.Warn(args)
}

func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}
func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

func With(args ...interface{}) *zap.SugaredLogger {
	return sugarLoggerWithOnceCallerAdd.With(args...)
}

func WithField(args ...interface{}) *zap.SugaredLogger {
	return sugarLoggerWithOnceCallerAdd.With(args...)
}
