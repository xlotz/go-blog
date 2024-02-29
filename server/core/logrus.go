package core

import (
	"bytes"
	"fmt"
	rotatelog "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"go-blog/global"
	"io/fs"
	"os"
	"path"
	"time"
)

type LogFormatter struct {
}

/**
自定义日志格式
*/

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	log := global.Config.Logger
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件格式
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "%s [%s] [%s] %s %s %s\n", log.Prefix, timestamp, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s [%s] [%s] %s\n", log.Prefix, timestamp, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	log := global.Config.Logger
	// 设置输出到文件
	_, err := os.ReadDir(log.Director)
	if err != nil {
		os.MkdirAll(log.Director, fs.ModePerm)
	}
	//fileName := path.Join(log.Director, log.LogFileName) + ".log"
	//file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	mLog := logrus.New()      //新建一个实例
	mLog.SetOutput(os.Stdout) // 设置输出类型

	mLog.SetReportCaller(log.ShowLine) // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) // 设置自定义格式
	level, err := logrus.ParseLevel(log.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) // 设置最低的级别
	InitDefaultLogger()  //调用全局

	//// 写日志文件
	//writers := []io.Writer{
	//	LogFileSplit(fileName),
	//	os.Stdout,
	//}
	//fileAndStdoutWriter := io.MultiWriter(writers...)
	//if err == nil {
	//	mLog.SetOutput(fileAndStdoutWriter)
	//} else {
	//	mLog.SetOutput(file)
	//}

	return mLog
}

/**
初始化默认日志
*/

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)                           // 设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{})                  // 设置自定义格式
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) // 设置最低的级别
}

/**
分割日志
*/

func LogFileSplit(fileName string) *rotatelog.RotateLogs {
	log := global.Config.Logger
	logier, err := rotatelog.New(
		fileName+".%Y%m%d",               // 切割后的日志文件名
		rotatelog.WithLinkName(fileName), // 为最新的日志创建软链接
		//rotatelog.WithMaxAge(time.Duration(log.LogMaxAge)*24*time.Hour), // 文件最大保存时间
		rotatelog.WithRotationTime(24*time.Hour),   // 文件切割时间间隔
		rotatelog.WithRotationCount(log.LogMaxAge), // 最大保存份数 和最大保存时间两个只生效一个
	)

	if err != nil {
		panic(err)
	}
	return logier
}
