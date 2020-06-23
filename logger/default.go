package logger

import (
	"bufio"
	"github.com/jessevdk/go-flags"
	"github.com/joyllee/blocks/config"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"time"
)

var opts struct {
	ConfigFile string `short:"c" required:"true" name:"config file"`
}

var logger = logrus.New()

func InitLogger() *logrus.Logger {
	return logger
}

func init() {
	// init configs
	if _, err := flags.Parse(&opts); err != nil {
		panic(err)
	}
	loadConfigFile(opts.ConfigFile)

	var logFormat logrus.Formatter
	if config.ServerConfig.Logger.LogFormat == "json" {
		logFormat = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	} else {
		logFormat = &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
	}
	logger.SetFormatter(logFormat)

	baseLogPath := path.Join(config.ServerConfig.Logger.LogDir,
		config.ServerConfig.Logger.LogFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Println("config local file system logger errors")
	}

	switch config.ServerConfig.Logger.LogLevel {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
		logger.SetOutput(os.Stderr)
		// 显示行号等信息
		logger.SetReportCaller(true)
	case "info":
		setNull()
		logger.SetLevel(logrus.InfoLevel)
		logger.SetOutput(os.Stderr)
		logger.SetReportCaller(true)
	case "warn":
		setNull()
		logger.SetLevel(logrus.WarnLevel)
	case "errors":
		setNull()
		logger.SetLevel(logrus.ErrorLevel)
		logger.SetReportCaller(true)
	default:
		setNull()
		logger.SetLevel(logrus.InfoLevel)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, logFormat)
	logger.AddHook(lfHook)
}

func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logrus.Info("err", err)
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}

func loadConfigFile(filePath string) {
	logrus.Debug("load config %s", filePath)
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("fail to read config", err)
		panic(err)
	}
	if err := viper.Unmarshal(&config.ServerConfig); err != nil {
		logrus.Fatal("fail to unmarshal config", err)
		panic(err)
	}
}
