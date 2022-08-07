package v2ray

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
	util.HandleError(err, accessLogPath+" is open failed")
	lastOffset, _ := f.Seek(offset, io.SeekStart)
	line := bufio.NewReader(f)
	loop := true
	var multLineValStr []string
	startUnix := time.Now().Unix()
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
				lineValStr := string(lineVal)
				if checkLineVal(lineValStr) {
					multLineValStr = append(multLineValStr, lineValStr)
					lastOffset, _ = f.Seek(0, io.SeekCurrent)
				}
			}
			endUnix := time.Now().Unix()
			if len(multLineValStr) > 50 || endUnix-startUnix > 2 {
				startUnix = endUnix
				saveMultiLineVal(multLineValStr)
				multLineValStr = []string{}
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
	f, err := os.OpenFile("tmp/v2ray_accessLog_lastoffset.id", os.O_WRONLY|os.O_CREATE, 0666)
	util.HandleError(err, "saveLastOffSet failed")
	defer f.Close()
	write := bufio.NewWriter(f)
	_, _ = write.WriteString(strconv.FormatInt(lastOffset, 10))
	_ = write.Flush()
}

func checkLineVal(lineVal string) bool {
	if lineVal == "" {
		return false
	}
	lineStrs := strings.Split(lineVal, " ")
	if len(lineStrs) < 4 {
		return false
	}
	return true
}

func saveMultiLineVal(multiLineVal []string) {
	var v2rayAccessLog model.V2rayAccessLog
	v2rayIpCountMap := make(map[string]int)
	for _, lineVal := range multiLineVal {
		v2rayAccessLog = buildV2rayAccessLogByLineStrs(lineVal)
		_ = v2rayAccessLog.SaveV2rayAccessLog(v2rayAccessLog)
		BuildV2rayIpCountMap(v2rayIpCountMap, v2rayAccessLog)
	}
	SaveIpInfo(v2rayIpCountMap)
}

func BuildV2rayIpCountMap(v2rayIpCountMap map[string]int, v2rayAccessLog model.V2rayAccessLog) {
	if count, ok := v2rayIpCountMap[v2rayAccessLog.Ip]; ok {
		v2rayIpCountMap[v2rayAccessLog.Ip] = count + 1
	} else {
		v2rayIpCountMap[v2rayAccessLog.Ip] = 1
	}
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
		reg := regexp.MustCompile(`([0-9.]+):([0-9]+)`)
		reason := reg.FindAllString(v2rayAccessLog.Reason, -1)
		if len(reason) > 0 {
			v2rayAccessLog.RemoteAdr = strings.Split(reason[0], ":")[0]
			v2rayAccessLog.RemoteAdrPort = strings.Split(reason[0], ":")[1]
		}
	}
	return v2rayAccessLog
}
