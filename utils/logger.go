package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// 全局日志实例
var Logger *log.Logger

// 全局变量用于持有日志文件引用
var logfile *os.File

func init() {
	var output io.Writer = os.Stdout
	logFilePath := os.Getenv("GOU_TEST_PLG_LOG")

	if logFilePath == "" {
		logFilePath = "./logs" // 假设相对路径正确
		// 确保日志目录存在
		err := os.MkdirAll(logFilePath, 0755)
		if err != nil {
			log.Fatalf("创建日志目录失败: %v", err)
		}
	}

	// 创建或打开日志文件
	logFileFullPath := filepath.Join(logFilePath, "plugin.log")
	var err error
	logfile, err = os.OpenFile(logFileFullPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("无法打开日志文件：%v，将继续使用标准输出\n", err)
	} else {
		output = logfile
	}

	Logger = log.New(output, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)

	Logger.Println("日志实例创建成功")
}

// 在程序的其他部分，确保在程序退出前关闭日志文件
func CloseLog() {
	if logfile != nil {
		Logger.Println("日志实例关闭")
		logfile.Close()
	}
}

// Info 记录信息级别的日志
func Log(v ...interface{}) {
	Logger.Println(v...)
}

// Infof 记录带有格式的信息级别日志
func Logf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}
