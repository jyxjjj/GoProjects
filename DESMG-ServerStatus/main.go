package main

import (
	"ServerStatus/Structs"
	"encoding/json"
	"flag"
	"os"
)

type MonitorData = Structs.MonitorData

var debug string

func init() {
	flag.StringVar(&debug, "d", "false", "Debug mode")
	flag.Parse()
	if debug == "false" || debug == "" {
		debug = os.Getenv("DEBUG")
		if debug == "false" || debug == "" {
			debug = "false"
		}
	}
}

func main() {
	data := run()
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	if debug == "true" {
		println(string(jsonData))
	} else {
		send("", data.DeviceID, jsonData)
	}
}

func run() MonitorData {
	iorw := GetIORW()
	data := MonitorData{
		DeviceID:   GetDeviceID(),
		CPUModel:   GetCPUModel(),
		CPUNum:     GetCPUNum(),
		CPUFreq:    GetCPUFreq(),
		CPUUsage:   GetCPUUsage(),
		MemSize:    GetMemSize(),
		MemUsed:    GetMemUsed(),
		NumProcess: GetNumProcess(),
		DiskName:   GetDiskName(),
		DiskUsage:  GetDiskUsage(),
		DiskSize:   GetDiskSize(),
		Uptime:     GetUptime(),
		IORead:     iorw[0],
		IOWrite:    iorw[1],
	}
	return data
}
