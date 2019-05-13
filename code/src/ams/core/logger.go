package core

import (
	"os"

	logging "github.com/op/go-logging"
)

var Logger = logging.MustGetLogger("ams")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} %{level:.4s} %{id:03x}%{color:reset} # %{message}`,
)

func SetupLogger() {
	filepath := os.Getenv("AMS_LOG_FILE")
	if "" == filepath {
		filepath = "/var/log/ams.log"
	}
	loglevel := os.Getenv("AMS_LOG_LEVEL")
	if "" == loglevel {
		loglevel = "INFO"
	}

	os.Remove(filepath)

	logfile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if nil != err {
		panic(err)
	}

	// 标准输出
	stdBackend := logging.NewLogBackend(os.Stderr, "", 0)
	stdBackendFormatter := logging.NewBackendFormatter(stdBackend, format)
	stdBackendLeveled := logging.AddModuleLevel(stdBackendFormatter)
	stdBackendLeveled.SetLevel(logging.DEBUG, "")

	// 文件输出
	fileBackend := logging.NewLogBackend(logfile, "", 0)
	fileBackendFormatter := logging.NewBackendFormatter(fileBackend, format)
	fileBackendLeveled := logging.AddModuleLevel(fileBackendFormatter)
	fileLoglevel, err := logging.LogLevel(loglevel)
	if nil != err {
		fileBackendLeveled.SetLevel(logging.INFO, "")
	} else {
		fileBackendLeveled.SetLevel(fileLoglevel, "")
	}

	logging.SetBackend(stdBackendLeveled, fileBackendLeveled)
}
