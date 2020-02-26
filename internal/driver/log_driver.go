package driver

import (
	"net/http"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogDriver ...
type LogDriver struct {
	level       logrus.Level
	sugarLogger *zap.SugaredLogger
}

// NewLogDriver ...
func NewLogDriver(filename string, level logrus.Level) *LogDriver {
	writerSyncer := getLogWriter(filename)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	return &LogDriver{
		level:       level,
		sugarLogger: logger.Sugar(),
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

func (l *LogDriver) SimpleHttpGet(url string) {
	l.sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		l.sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		l.sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
