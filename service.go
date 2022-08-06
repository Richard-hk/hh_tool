package main

import (
	"fmt"
	"hh_tool/monitor/app/v2ray"
	"hh_tool/monitor/web/bit_domain"
	"hh_tool/util"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	go receiveSignal()
	go Monitor()
}

func receiveSignal() {
	signal.Notify(util.StopSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-util.StopSignal
	close(util.StopSignal)
	fmt.Println("receive stop signal after 3s will shut down ")
	time.Sleep(3 * time.Second)
	os.Exit(0)
}

func Monitor() {
	for _, app := range util.MonitorApp {
		go func(app string) {
			MonitorApp(app)
		}(app)
	}
	for _, web := range util.MonitorWeb {
		go func(web string) {
			MonitorWeb(web)
		}(web)
	}
}

// 监控app
func MonitorApp(app string) {
	switch app {
	case util.V2RAY:
		v2rayProcessor, _ := v2ray.NewV2rayProcessor()
		v2rayProcessor.Monitor()
	}
}

// 监控app
func MonitorWeb(app string) {
	switch app {
	case util.BitDomain:
		bit_domain, _ := bit_domain.NewBitDomainProcessor()
		bit_domain.Monitor()
	}
}
