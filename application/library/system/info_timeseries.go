/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package system

import (
	"context"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/net"
	"github.com/webx-top/com"
	"github.com/webx-top/echo/defaults"
	"github.com/webx-top/echo/param"

	"github.com/admpub/log"
	"github.com/admpub/nging/application/library/cron"
	"github.com/admpub/nging/application/library/msgbox"
	"github.com/admpub/nging/application/registry/alert"
)

var (
	mutext         sync.Mutex
	realTimeStatus *RealTimeStatus
	// CancelRealTimeStatusCollection 取消实时状态搜集
	CancelRealTimeStatusCollection = func() {}
)

func init() {
	alert.Topics.Add(`systemStatus`, `系统状态`)
}

// RealTimeStatusObject 实时状态
func RealTimeStatusObject(n ...int) *RealTimeStatus {
	if len(n) == 0 || n[0] <= 0 {
		return realTimeStatus
	}
	r := &RealTimeStatus{
		CPU: TimeSeries{},
		Mem: TimeSeries{},
		Net: NewNetIOTimeSeries(),
	}
	max := n[0]
	if max < len(realTimeStatus.CPU) {
		r.CPU = realTimeStatus.CPU[len(realTimeStatus.CPU)-max:]
	} else {
		r.CPU = realTimeStatus.CPU
	}
	if max < len(realTimeStatus.Mem) {
		r.Mem = realTimeStatus.Mem[len(realTimeStatus.Mem)-max:]
	} else {
		r.Mem = realTimeStatus.Mem
	}
	if max < len(realTimeStatus.Net.BytesSent) {
		r.Net.BytesSent = realTimeStatus.Net.BytesSent[len(realTimeStatus.Net.BytesSent)-max:]
	} else {
		r.Net.BytesSent = realTimeStatus.Net.BytesSent
	}
	if max < len(realTimeStatus.Net.BytesRecv) {
		r.Net.BytesRecv = realTimeStatus.Net.BytesRecv[len(realTimeStatus.Net.BytesRecv)-max:]
	} else {
		r.Net.BytesRecv = realTimeStatus.Net.BytesRecv
	}
	if max < len(realTimeStatus.Net.PacketsSent) {
		r.Net.PacketsSent = realTimeStatus.Net.PacketsSent[len(realTimeStatus.Net.PacketsSent)-max:]
	} else {
		r.Net.PacketsSent = realTimeStatus.Net.PacketsSent
	}
	if max < len(realTimeStatus.Net.PacketsRecv) {
		r.Net.PacketsRecv = realTimeStatus.Net.PacketsRecv[len(realTimeStatus.Net.PacketsRecv)-max:]
	} else {
		r.Net.PacketsRecv = realTimeStatus.Net.PacketsRecv
	}
	return r
}

// RealTimeStatusIsListening 是否正在监听实时状态
func RealTimeStatusIsListening() bool {
	return realTimeStatus != nil && realTimeStatus.status == `started`
}

// ListenRealTimeStatus 监听实时状态
func ListenRealTimeStatus(cfg *Settings) {
	mutext.Lock()
	defer mutext.Unlock()
	interval := time.Second * 2
	max := 80
	if RealTimeStatusIsListening() {
		CancelRealTimeStatusCollection()
		realTimeStatus.SetSettings(cfg, interval, max)
	} else {
		realTimeStatus = NewRealTimeStatus(cfg, interval, max)
	}

	msgbox.Info(`System Monitor`, `Starting collect server status`)

	ctx, cancel := context.WithCancel(context.Background())
	go realTimeStatus.Listen(ctx)
	CancelRealTimeStatusCollection = func() {
		if RealTimeStatusIsListening() {
			cancel()
		}
	}
}

// NewRealTimeStatus 创建实时状态数据结构
func NewRealTimeStatus(cfg *Settings, interval time.Duration, maxSize int) *RealTimeStatus {
	r := &RealTimeStatus{
		max:      maxSize,
		interval: interval,
		CPU:      TimeSeries{},
		Mem:      TimeSeries{},
		Net:      NewNetIOTimeSeries(),
	}
	return r.SetSettings(cfg, interval, maxSize)
}

// NewNetIOTimeSeries 创建网络IO时序数据结构
func NewNetIOTimeSeries() NetIOTimeSeries {
	return NetIOTimeSeries{
		lastBytesSent:   LastTimeValue{},
		lastBytesRecv:   LastTimeValue{},
		lastPacketsSent: LastTimeValue{},
		lastPacketsRecv: LastTimeValue{},
		BytesSent:       TimeSeries{},
		BytesRecv:       TimeSeries{},
		PacketsSent:     TimeSeries{},
		PacketsRecv:     TimeSeries{},
	}
}

// LastTimeValue 上次时间的状态值
type LastTimeValue struct {
	Time  time.Time
	Value float64
}

// NetIOTimeSeries 网络IO时序数据结构
type NetIOTimeSeries struct {
	lastBytesSent   LastTimeValue
	lastBytesRecv   LastTimeValue
	lastPacketsSent LastTimeValue
	lastPacketsRecv LastTimeValue

	BytesSent   TimeSeries
	BytesRecv   TimeSeries
	PacketsSent TimeSeries
	PacketsRecv TimeSeries
}

// RealTimeStatus 实时状态数据结构
type RealTimeStatus struct {
	max         int
	interval    time.Duration
	CPU         TimeSeries
	Mem         TimeSeries
	Net         NetIOTimeSeries
	settings    *Settings
	reportEmail []string
	reportTime  time.Time
	status      string
	lock        sync.RWMutex
}

// Listen 监听
func (r *RealTimeStatus) Listen(ctx context.Context) *RealTimeStatus {
	r.status = `started`
	info := &DynamicInformation{}
	t := time.NewTicker(r.interval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			msgbox.Warn(`System Monitor`, `Exit server real-time status collection`)
			r.status = `stoped`
			return r
		case <-t.C:
			info.NetMemoryCPU()
			if len(info.CPUPercent) > 0 {
				r.CPUAdd(info.CPUPercent[0])
			} else {
				r.CPUAdd(0)
			}
			r.MemAdd(info.Memory.Virtual.UsedPercent)
			if len(info.NetIO) > 0 {
				r.NetAdd(info.NetIO[0])
			}
			//log.Info(`Collect server status`)
		}
	}
}

var emptyTime = time.Time{}

func checkAndSendAlarm(r *RealTimeStatus, value float64, typ string) {
	if r == nil || r.settings == nil {
		return
	}
	if !r.settings.AlarmOn {
		return
	}
	switch typ {
	case `CPU`:
		if r.settings.AlarmThreshold.CPU > 0 && r.settings.AlarmThreshold.CPU < value {
			r.sendAlarm(r.settings.AlarmThreshold.CPU, value, typ)
			return
		}
	case `Mem`:
		if r.settings.AlarmThreshold.Memory > 0 && r.settings.AlarmThreshold.Memory < value {
			r.sendAlarm(r.settings.AlarmThreshold.Memory, value, typ)
			return
		}
	}
	if !r.reportTime.IsZero() {
		r.reportTime = emptyTime
		return
	}
}

type alarmContent struct {
	title    string
	hostname string
	typeName string
	value    string
}

func (a *alarmContent) EmailContent(params param.Store) []byte {
	return com.Str2bytes(`<h1>` + a.title + `</h1><p>主机名: ` + a.hostname + `<br />` + a.typeName + `使用率: ` + a.value + `%<br />时间: ` + time.Now().Format(time.RFC3339) + `<br /></p>`)
}

func (a *alarmContent) MarkdownContent(params param.Store) []byte {
	return com.Str2bytes(`### ` + a.title + "\n" + `**主机名**: ` + a.hostname + "\n**" + a.typeName + `使用率**: ` + a.value + `%` + "\n" + `**时间**: ` + time.Now().Format(time.RFC3339) + "\n")
}

func (r *RealTimeStatus) sendAlarm(alarmThreshold, value float64, typ string) *RealTimeStatus {
	now := time.Now()
	if r.reportTime.IsZero() || now.Sub(r.reportTime) < time.Minute*5 { // 连续5分钟达到阀值时发邮件告警
		return nil
	}
	var typeName string
	switch typ {
	case `CPU`:
		typeName = `CPU`
	case `Mem`:
		typeName = `内存`
	default:
		return nil
	}
	hostname, _ := os.Hostname()
	title := fmt.Sprintf(`【`+hostname+`】`+typeName+`使用率超出%v%%`, alarmThreshold)
	ct := alarmContent{
		title:    title,
		hostname: hostname,
		typeName: typeName,
		value:    fmt.Sprint(value),
	}
	params := param.Store{
		`title`:   title,
		`content`: ct,
	}
	ctx := defaults.NewMockContext()
	alert.SendTopic(ctx, `systemStatus`, params)
	if len(r.reportEmail) == 0 {
		return r
	}
	content := ct.EmailContent(params)
	var cc []string
	if len(r.reportEmail) > 1 {
		cc = r.reportEmail[1:]
	}
	err := cron.SendMail(r.reportEmail[0], `administrator`, title, content, cc...)
	if err != nil {
		log.Error(err)
	}
	r.reportTime = now
	return r
}

func (r *RealTimeStatus) SetSettings(c *Settings, interval time.Duration, max int) *RealTimeStatus {
	r.settings = c
	var reportEmail []string
	if c != nil {
		if len(c.ReportEmail) > 0 {
			for _, email := range strings.Split(c.ReportEmail, "\n") {
				email = strings.TrimSpace(email)
				if len(email) == 0 {
					continue
				}
				reportEmail = append(reportEmail, email)
			}
		}
	}
	r.reportEmail = reportEmail
	r.interval = interval
	r.max = max
	return r
}

func (r *RealTimeStatus) CPUAdd(y float64) *RealTimeStatus {
	if r.max <= 0 {
		return r
	}
	r.lock.Lock()
	checkAndSendAlarm(r, y, `CPU`)
	l := len(r.CPU)
	if l >= r.max {
		r.CPU = r.CPU[1+l-r.max:]
	}
	r.CPU = append(r.CPU, NewXY(y))
	r.lock.Unlock()
	return r
}

func (r *RealTimeStatus) MemAdd(y float64) *RealTimeStatus {
	if r.max <= 0 {
		return r
	}
	r.lock.Lock()
	checkAndSendAlarm(r, y, `Mem`)
	l := len(r.Mem)
	if l >= r.max {
		r.Mem = r.Mem[1+l-r.max:]
	}
	r.Mem = append(r.Mem, NewXY(y))
	r.lock.Unlock()
	return r
}

func (r *RealTimeStatus) NetAdd(stat net.IOCountersStat) *RealTimeStatus {
	if r.max <= 0 {
		return r
	}
	r.lock.Lock()
	now := time.Now()
	l := len(r.Net.BytesRecv)
	if l >= r.max {
		r.Net.BytesRecv = r.Net.BytesRecv[1+l-r.max:]
	}
	n := float64(stat.BytesRecv)
	var speed float64
	if !r.Net.lastBytesRecv.Time.IsZero() {
		speed = (n - r.Net.lastBytesRecv.Value) / now.Sub(r.Net.lastBytesRecv.Time).Seconds()
		speed = math.Ceil(speed)
	} else {
		speed = 0
	}
	r.Net.BytesRecv = append(r.Net.BytesRecv, NewXY(speed))
	r.Net.lastBytesRecv.Time = now
	r.Net.lastBytesRecv.Value = n

	l = len(r.Net.BytesSent)
	if l >= r.max {
		r.Net.BytesSent = r.Net.BytesSent[1+l-r.max:]
	}
	n = float64(stat.BytesSent)
	if !r.Net.lastBytesSent.Time.IsZero() {
		speed = (n - r.Net.lastBytesSent.Value) / now.Sub(r.Net.lastBytesSent.Time).Seconds()
		speed = math.Ceil(speed)
	} else {
		speed = 0
	}
	r.Net.BytesSent = append(r.Net.BytesSent, NewXY(speed))
	r.Net.lastBytesSent.Time = now
	r.Net.lastBytesSent.Value = n

	l = len(r.Net.PacketsRecv)
	if l >= r.max {
		r.Net.PacketsRecv = r.Net.PacketsRecv[1+l-r.max:]
	}
	n = float64(stat.PacketsRecv)
	if !r.Net.lastPacketsRecv.Time.IsZero() {
		speed = (n - r.Net.lastPacketsRecv.Value) / now.Sub(r.Net.lastPacketsRecv.Time).Seconds()
		speed = math.Ceil(speed)
	} else {
		speed = 0
	}
	r.Net.PacketsRecv = append(r.Net.PacketsRecv, NewXY(speed))
	r.Net.lastPacketsRecv.Time = now
	r.Net.lastPacketsRecv.Value = n

	l = len(r.Net.PacketsSent)
	if l >= r.max {
		r.Net.PacketsSent = r.Net.PacketsSent[1+l-r.max:]
	}
	n = float64(stat.PacketsSent)
	if !r.Net.lastPacketsSent.Time.IsZero() {
		speed = (n - r.Net.lastPacketsSent.Value) / now.Sub(r.Net.lastPacketsSent.Time).Seconds()
		speed = math.Ceil(speed)
	} else {
		speed = 0
	}
	r.Net.PacketsSent = append(r.Net.PacketsSent, NewXY(speed))
	r.Net.lastPacketsSent.Time = now
	r.Net.lastPacketsSent.Value = n
	r.lock.Unlock()
	return r
}

type (
	TimeSeries []XY
	XY         [2]interface{}
)

func NewXY(y float64) XY {
	x := time.Now().UnixNano() / 1e6 //毫秒
	return XY{x, y}
}
