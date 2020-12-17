package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	// \033[31m:设置字体颜色   	\033[0m : 关闭所有属性
	// log.LstdFlags: 显示日期	log.Lshortfile:显示文件及行数
	infoLog = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers = []*log.Logger{errorLog, infoLog}
	mu      sync.Mutex
	//sync.Mutex: 互斥锁
)

// log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf // "%s%v" 格式化输出
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
		//ioutil.Discard 不执行任何操作，永远返回成功
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
