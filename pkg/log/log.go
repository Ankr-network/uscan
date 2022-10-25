package log

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "[uscan] INF ", log.Ldate|log.Ltime|log.Lmsgprefix)
	errorLogger = log.New(os.Stdout, "[uscan] ERR ", log.Ldate|log.Ltime|log.Lmsgprefix)
	fatalLogger = log.New(os.Stdout, "[uscan] FTL ", log.Ldate|log.Ltime|log.Lmsgprefix)
)

func Info(msg ...any) {
	infoLogger.Println(msg...)
}

func Infof(format string, msg ...any) {
	infoLogger.Printf(format, msg...)
}
func Error(msg ...any) {
	errorLogger.Println(msg...)
}

func Errorf(format string, msg ...any) {
	errorLogger.Printf(format, msg...)
}

func Fatal(msg ...any) {
	fatalLogger.Fatalln(msg...)
}

func Fatalf(format string, msg ...any) {
	fatalLogger.Fatalf(format, msg...)
}
