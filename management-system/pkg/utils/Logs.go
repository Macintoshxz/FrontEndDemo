package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	INFO  = "[Info]"
	Warn  = "[Warn]"
	ERROR = "[Error]"
)

type Logs struct{}

func LogInfo(content string) {
	writeServiceLog(INFO, content)
}

func LogWarn(content string) {
	writeServiceLog(Warn, content)
}

func LogError(content string) {
	writeServiceLog(ERROR, content)
}

func LogUploadInfo(uid, content, file, ip string) {
	writeServiceLog(INFO, "uid "+uid+" upload "+content+" "+file+" from ip "+ip)
}

/********************************************************************************
* @功    能：写服务日志到文件内
* @输入参数：logType 日志类型；content 日志内容
* @返 回 值：无
* @日    期：2022年5月22日18:00:00
* @版    本：1.0
********************************************************************************/
func writeServiceLog(logType string, content string) {
	getwd, _ := os.Getwd()
	path := filepath.Dir(getwd)
	file := path + "/logs/" + time.Now().Format("20060102") + ".log"
	logContent := logType + " " + time.Now().Format("2006-01-02 15:04:05") + " " + content
	writeLogsToFile(file, logContent)
}

/********************************************************************************
* @功    能：写日志到文件内
* @输入参数：fileName 文件名称；content 日志内容
* @返 回 值：无
* @日    期：2022年5月22日17:59:57
* @版    本：1.0
********************************************************************************/
func writeLogsToFile(fileName string, content string) {
	var f1 *os.File
	var err error

	if exist, _ := pathExists(fileName); exist {
		//如果文件存在,打开文件
		f1, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Println(ERROR, " Log file opening failed "+err.Error())
			return
		}
	} else {
		//如果文件不存在，创建文件
		f1, err = os.Create(fileName)
		if err != nil {
			log.Println(ERROR, " Log file creation failed "+err.Error())
			return
		}
	}
	_, err = io.WriteString(f1, content+"\n") //写入文件
	if err != nil {
		log.Println(ERROR, " Write log error "+err.Error())
	}
}

/**
* @功    能：错误写入日志抛出
* @输入参数：err , 错误描述
* @返 回 值：返回值描述
* @日    期：2023年08月25日18:33:25
* @版    本：版本号
 */
func CheckError(err error, content string) {
	if err != nil {
		//TODO 捕获不到异常
		// LogError(content + "-- err:" + err.Error())
		panic(err)
	}
}
