package base
import (
	"fmt"
	"github.com/op/go-logging"
	"os"
)



var format = logging.MustStringFormatter(`%{time:15:04:05.000} %{shortfunc} > %{level:.4s}  %{message}`,)

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}


func LoggerSetting() *logging.Logger{
	var log = logging.MustGetLogger("example")
	logFile, err := os.OpenFile("log/access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
	if err != nil{
		fmt.Println(err)
		os.Exit(10)
	}
	backend := logging.NewLogBackend(logFile, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backend2Formatter)
	return log
}
