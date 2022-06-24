package monitor

import (
	"bufio"
	"hh_tool/model"
	"hh_tool/service"
	"hh_tool/util"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const StatusAccepted = "accepted"
const StatusRejected = "rejected"

type V2rayProcessor struct {
	service.BaseProcessor
}

func NewV2rayProcessor() (*V2rayProcessor, error) {
	return &V2rayProcessor{}, nil
}

func (p *V2rayProcessor) Monitor() {
	v2rayOpen := viper.GetBool("monitor.v2ray.open")
	if !v2rayOpen {
		logrus.Info("v2ray monitor is not open")
		return
	}
	go p.monitorAccessLog()
}

// 监控accessLog
func (p *V2rayProcessor) monitorAccessLog() {
	accessLogOpen := viper.GetBool("monitor.v2ray.accessLog.open")
	if !accessLogOpen {
		logrus.Info("v2ray accessLog monitor is not open")
		return
	}
	lastOffset := getLastOffSet()
	readAccessLog(lastOffset)
}

func readAccessLog(offset int64) {
	accessLogPath := viper.GetString("monitor.v2ray.accessLog.path")
	f, err := os.OpenFile(accessLogPath, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		logrus.Error(accessLogPath + " is open failed")
	}
	lastOffset, _ := f.Seek(offset, io.SeekStart)
	line := bufio.NewReader(f)
	loop := true
	for loop {
		select {
		case _, ok := <-util.StopSignal:
			if !ok {
				saveLastOffSet(lastOffset)
				loop = false
			}
		default:
			lineVal, _, err := line.ReadLine()
			if err != nil {
				time.Sleep(time.Duration(viper.GetInt64("monitor.v2ray.accessLog.sleep")) * time.Second)
			} else {
				lastOffset, _ = f.Seek(0, io.SeekCurrent)
				lineValStr := string(lineVal)
				if checkLineVal(lineValStr) {
					saveLineVal(lineValStr)
				}
			}
		}
	}
}

func getLastOffSet() int64 {
	f, err := os.OpenFile("tmp/v2ray_accessLog_lastoffset.id", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		logrus.Error("v2ray_accessLog_lastoffset.id open failed")
		return 0
	}
	defer f.Close()
	line := bufio.NewReader(f)
	lineVal, _, _ := line.ReadLine()
	lastOffset, _ := strconv.ParseInt(string(lineVal), 10, 64)
	return lastOffset
}

func saveLastOffSet(lastOffset int64) {
	f, _ := os.OpenFile("tmp/v2ray_accessLog_lastoffset.id", os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	write := bufio.NewWriter(f)
	write.WriteString(strconv.FormatInt(lastOffset, 10))
	write.Flush()
}

func checkLineVal(lineVal string) bool {
	if lineVal == "" {
		return false
	}
	return true
}

func saveLineVal(lineVal string) {
	v2rayAccessLog := buildV2rayAccessLogByLineStrs(lineVal)
	v2rayAccessLog.SaveV2rayAccessLog(v2rayAccessLog)
}

func buildV2rayAccessLogByLineStrs(lineVal string) (v2rayAccessLog model.V2rayAccessLog) {
	lineStrs := strings.Split(lineVal, " ")
	v2rayAccessLog.Dt = strings.Replace(lineStrs[0], "/", "", -1)
	v2rayAccessLog.Time = lineStrs[1]
	v2rayAccessLog.Ip = strings.Split(lineStrs[2], ":")[0]
	v2rayAccessLog.Port = strings.Split(lineStrs[2], ":")[1]
	v2rayAccessLog.Status = lineStrs[3]
	if v2rayAccessLog.Status == StatusAccepted {
		v2rayAccessLog.Type = strings.Split(lineStrs[4], ":")[0]
		v2rayAccessLog.RemoteAdr = strings.Split(lineStrs[4], ":")[1]
		v2rayAccessLog.RemoteAdrPort = strings.Split(lineStrs[4], ":")[2]
	}
	if v2rayAccessLog.Status == StatusRejected {
		v2rayAccessLog.Reason = strings.Trim(strings.Split(lineVal, StatusRejected)[1], " ")
		reg := regexp.MustCompile(`([0-9.]+)`)
		reason := reg.FindAllString((v2rayAccessLog.Reason), -1)
		v2rayAccessLog.RemoteAdr = reason[0]
		v2rayAccessLog.RemoteAdrPort = reason[1]
	}
	return v2rayAccessLog
}
