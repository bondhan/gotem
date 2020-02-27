package driver

import (
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogDriver ...
type LogDriver struct {
	filename    string
	level       logrus.Level
	sugarLogger *zap.SugaredLogger
}

// NewLogDriver ...
func NewLogDriver(filename string, level logrus.Level) *LogDriver {
	return &LogDriver{
		filename: filename,
		level:    level,
	}
}

// InitLog ...
func (l *LogDriver) InitLog() {
	logrus.SetLevel(l.level)
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	writerSyncer := getLogWriter(l.filename)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	l.sugarLogger = zap.New(core, zap.AddCaller()).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logFName string) zapcore.WriteSyncer {
	dt := time.Now()
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "logs/" + dt.Format("20060102") + "_" + logFName,
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (l *LogDriver) Debug(err error) {
	l.Debug(err)

	// post to db or es
	// if fail then write to file
	if err != nil {
		l.sugarLogger.Debug(err)
	}
}

func (l *LogDriver) Error(err error) {
	l.Error(err)

	// post to db or es
	// if fail then write to file
	if err != nil {
		l.sugarLogger.Error(err)
	}
}

func (l *LogDriver) Warn(err error) {
	l.Warn(err)

	// post to db or es
	// if fail then write to file
	if err != nil {
		l.sugarLogger.Warn(err)
	}
}
