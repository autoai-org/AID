// Copyright 2019 The CVPM Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	CpuName         string `json:"cpu"`
	Memory          uint64 `json:"memory"`
	Platform        string `json:"platform"`
	Os              string `json:"os"`
	PlatformVersion string `json:"platformVersion"`
}

func getSystemInfo() SystemInfo {
	var systemInfo SystemInfo
	v, _ := mem.VirtualMemory()
	systemInfo.Memory = v.Total
	platform, _ := host.Info()
	systemInfo.Platform = platform.Platform
	systemInfo.PlatformVersion = platform.PlatformVersion
	systemInfo.Os = platform.OS

	// Handle CPU
	cpuInfo, _ := cpu.Info()
	systemInfo.CpuName = cpuInfo[0].ModelName
	return systemInfo
}
