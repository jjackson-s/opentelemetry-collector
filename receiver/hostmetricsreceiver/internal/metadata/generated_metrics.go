// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "hostmetricsreceiver"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	ProcessCPUTime              MetricIntf
	ProcessDiskIo               MetricIntf
	ProcessMemoryPhysicalUsage  MetricIntf
	ProcessMemoryVirtualUsage   MetricIntf
	SystemCPULoadAverage15m     MetricIntf
	SystemCPULoadAverage1m      MetricIntf
	SystemCPULoadAverage5m      MetricIntf
	SystemCPUTime               MetricIntf
	SystemDiskIo                MetricIntf
	SystemDiskIoTime            MetricIntf
	SystemDiskMerged            MetricIntf
	SystemDiskOperationTime     MetricIntf
	SystemDiskOperations        MetricIntf
	SystemDiskPendingOperations MetricIntf
	SystemDiskWeightedIoTime    MetricIntf
	SystemFilesystemInodesUsage MetricIntf
	SystemFilesystemUsage       MetricIntf
	SystemMemoryUsage           MetricIntf
	SystemNetworkConnections    MetricIntf
	SystemNetworkDropped        MetricIntf
	SystemNetworkErrors         MetricIntf
	SystemNetworkIo             MetricIntf
	SystemNetworkPackets        MetricIntf
	SystemPagingFaults          MetricIntf
	SystemPagingOperations      MetricIntf
	SystemPagingUsage           MetricIntf
	SystemProcessesCount        MetricIntf
	SystemProcessesCreated      MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"process.cpu.time",
		"process.disk.io",
		"process.memory.physical_usage",
		"process.memory.virtual_usage",
		"system.cpu.load_average.15m",
		"system.cpu.load_average.1m",
		"system.cpu.load_average.5m",
		"system.cpu.time",
		"system.disk.io",
		"system.disk.io_time",
		"system.disk.merged",
		"system.disk.operation_time",
		"system.disk.operations",
		"system.disk.pending_operations",
		"system.disk.weighted_io_time",
		"system.filesystem.inodes.usage",
		"system.filesystem.usage",
		"system.memory.usage",
		"system.network.connections",
		"system.network.dropped",
		"system.network.errors",
		"system.network.io",
		"system.network.packets",
		"system.paging.faults",
		"system.paging.operations",
		"system.paging.usage",
		"system.processes.count",
		"system.processes.created",
	}
}

var metricsByName = map[string]MetricIntf{
	"process.cpu.time":               Metrics.ProcessCPUTime,
	"process.disk.io":                Metrics.ProcessDiskIo,
	"process.memory.physical_usage":  Metrics.ProcessMemoryPhysicalUsage,
	"process.memory.virtual_usage":   Metrics.ProcessMemoryVirtualUsage,
	"system.cpu.load_average.15m":    Metrics.SystemCPULoadAverage15m,
	"system.cpu.load_average.1m":     Metrics.SystemCPULoadAverage1m,
	"system.cpu.load_average.5m":     Metrics.SystemCPULoadAverage5m,
	"system.cpu.time":                Metrics.SystemCPUTime,
	"system.disk.io":                 Metrics.SystemDiskIo,
	"system.disk.io_time":            Metrics.SystemDiskIoTime,
	"system.disk.merged":             Metrics.SystemDiskMerged,
	"system.disk.operation_time":     Metrics.SystemDiskOperationTime,
	"system.disk.operations":         Metrics.SystemDiskOperations,
	"system.disk.pending_operations": Metrics.SystemDiskPendingOperations,
	"system.disk.weighted_io_time":   Metrics.SystemDiskWeightedIoTime,
	"system.filesystem.inodes.usage": Metrics.SystemFilesystemInodesUsage,
	"system.filesystem.usage":        Metrics.SystemFilesystemUsage,
	"system.memory.usage":            Metrics.SystemMemoryUsage,
	"system.network.connections":     Metrics.SystemNetworkConnections,
	"system.network.dropped":         Metrics.SystemNetworkDropped,
	"system.network.errors":          Metrics.SystemNetworkErrors,
	"system.network.io":              Metrics.SystemNetworkIo,
	"system.network.packets":         Metrics.SystemNetworkPackets,
	"system.paging.faults":           Metrics.SystemPagingFaults,
	"system.paging.operations":       Metrics.SystemPagingOperations,
	"system.paging.usage":            Metrics.SystemPagingUsage,
	"system.processes.count":         Metrics.SystemProcessesCount,
	"system.processes.created":       Metrics.SystemProcessesCreated,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

func (m *metricStruct) FactoriesByName() map[string]func(pdata.Metric) {
	return map[string]func(pdata.Metric){
		Metrics.ProcessCPUTime.Name():              Metrics.ProcessCPUTime.Init,
		Metrics.ProcessDiskIo.Name():               Metrics.ProcessDiskIo.Init,
		Metrics.ProcessMemoryPhysicalUsage.Name():  Metrics.ProcessMemoryPhysicalUsage.Init,
		Metrics.ProcessMemoryVirtualUsage.Name():   Metrics.ProcessMemoryVirtualUsage.Init,
		Metrics.SystemCPULoadAverage15m.Name():     Metrics.SystemCPULoadAverage15m.Init,
		Metrics.SystemCPULoadAverage1m.Name():      Metrics.SystemCPULoadAverage1m.Init,
		Metrics.SystemCPULoadAverage5m.Name():      Metrics.SystemCPULoadAverage5m.Init,
		Metrics.SystemCPUTime.Name():               Metrics.SystemCPUTime.Init,
		Metrics.SystemDiskIo.Name():                Metrics.SystemDiskIo.Init,
		Metrics.SystemDiskIoTime.Name():            Metrics.SystemDiskIoTime.Init,
		Metrics.SystemDiskMerged.Name():            Metrics.SystemDiskMerged.Init,
		Metrics.SystemDiskOperationTime.Name():     Metrics.SystemDiskOperationTime.Init,
		Metrics.SystemDiskOperations.Name():        Metrics.SystemDiskOperations.Init,
		Metrics.SystemDiskPendingOperations.Name(): Metrics.SystemDiskPendingOperations.Init,
		Metrics.SystemDiskWeightedIoTime.Name():    Metrics.SystemDiskWeightedIoTime.Init,
		Metrics.SystemFilesystemInodesUsage.Name(): Metrics.SystemFilesystemInodesUsage.Init,
		Metrics.SystemFilesystemUsage.Name():       Metrics.SystemFilesystemUsage.Init,
		Metrics.SystemMemoryUsage.Name():           Metrics.SystemMemoryUsage.Init,
		Metrics.SystemNetworkConnections.Name():    Metrics.SystemNetworkConnections.Init,
		Metrics.SystemNetworkDropped.Name():        Metrics.SystemNetworkDropped.Init,
		Metrics.SystemNetworkErrors.Name():         Metrics.SystemNetworkErrors.Init,
		Metrics.SystemNetworkIo.Name():             Metrics.SystemNetworkIo.Init,
		Metrics.SystemNetworkPackets.Name():        Metrics.SystemNetworkPackets.Init,
		Metrics.SystemPagingFaults.Name():          Metrics.SystemPagingFaults.Init,
		Metrics.SystemPagingOperations.Name():      Metrics.SystemPagingOperations.Init,
		Metrics.SystemPagingUsage.Name():           Metrics.SystemPagingUsage.Init,
		Metrics.SystemProcessesCount.Name():        Metrics.SystemProcessesCount.Init,
		Metrics.SystemProcessesCreated.Name():      Metrics.SystemProcessesCreated.Init,
	}
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"process.cpu.time",
		func(metric pdata.Metric) {
			metric.SetName("process.cpu.time")
			metric.SetDescription("Total CPU seconds broken down by different states.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"process.disk.io",
		func(metric pdata.Metric) {
			metric.SetName("process.disk.io")
			metric.SetDescription("Disk bytes transferred.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"process.memory.physical_usage",
		func(metric pdata.Metric) {
			metric.SetName("process.memory.physical_usage")
			metric.SetDescription("The amount of physical memory in use.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"process.memory.virtual_usage",
		func(metric pdata.Metric) {
			metric.SetName("process.memory.virtual_usage")
			metric.SetDescription("Virtual memory size.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.cpu.load_average.15m",
		func(metric pdata.Metric) {
			metric.SetName("system.cpu.load_average.15m")
			metric.SetDescription("Average CPU Load over 15 minutes.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"system.cpu.load_average.1m",
		func(metric pdata.Metric) {
			metric.SetName("system.cpu.load_average.1m")
			metric.SetDescription("Average CPU Load over 1 minute.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"system.cpu.load_average.5m",
		func(metric pdata.Metric) {
			metric.SetName("system.cpu.load_average.5m")
			metric.SetDescription("Average CPU Load over 5 minutes.")
			metric.SetUnit("1")
			metric.SetDataType(pdata.MetricDataTypeGauge)
		},
	},
	&metricImpl{
		"system.cpu.time",
		func(metric pdata.Metric) {
			metric.SetName("system.cpu.time")
			metric.SetDescription("Total CPU seconds broken down by different states.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.io",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.io")
			metric.SetDescription("Disk bytes transferred.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.io_time",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.io_time")
			metric.SetDescription("Time disk spent activated. On Windows, this is calculated as the inverse of disk idle time.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.merged",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.merged")
			metric.SetDescription("The number of disk reads merged into single physical disk access operations.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.operation_time",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.operation_time")
			metric.SetDescription("Time spent in disk operations.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.operations",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.operations")
			metric.SetDescription("Disk operations count.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.pending_operations",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.pending_operations")
			metric.SetDescription("The queue size of pending I/O operations.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.weighted_io_time",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.weighted_io_time")
			metric.SetDescription("Time disk spent activated multiplied by the queue length.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.filesystem.inodes.usage",
		func(metric pdata.Metric) {
			metric.SetName("system.filesystem.inodes.usage")
			metric.SetDescription("FileSystem inodes used.")
			metric.SetUnit("{inodes}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.filesystem.usage",
		func(metric pdata.Metric) {
			metric.SetName("system.filesystem.usage")
			metric.SetDescription("Filesystem bytes used.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.memory.usage",
		func(metric pdata.Metric) {
			metric.SetName("system.memory.usage")
			metric.SetDescription("Bytes of memory in use.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.connections",
		func(metric pdata.Metric) {
			metric.SetName("system.network.connections")
			metric.SetDescription("The number of connections.")
			metric.SetUnit("{connections}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.dropped",
		func(metric pdata.Metric) {
			metric.SetName("system.network.dropped")
			metric.SetDescription("The number of packets dropped.")
			metric.SetUnit("{packets}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.errors",
		func(metric pdata.Metric) {
			metric.SetName("system.network.errors")
			metric.SetDescription("The number of errors encountered.")
			metric.SetUnit("{errors}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.io",
		func(metric pdata.Metric) {
			metric.SetName("system.network.io")
			metric.SetDescription("The number of bytes transmitted and received.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.network.packets",
		func(metric pdata.Metric) {
			metric.SetName("system.network.packets")
			metric.SetDescription("The number of packets transferred.")
			metric.SetUnit("{packets}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.paging.faults",
		func(metric pdata.Metric) {
			metric.SetName("system.paging.faults")
			metric.SetDescription("The number of page faults.")
			metric.SetUnit("{faults}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.paging.operations",
		func(metric pdata.Metric) {
			metric.SetName("system.paging.operations")
			metric.SetDescription("The number of paging operations.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.paging.usage",
		func(metric pdata.Metric) {
			metric.SetName("system.paging.usage")
			metric.SetDescription("Swap (unix) or pagefile (windows) usage.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.processes.count",
		func(metric pdata.Metric) {
			metric.SetName("system.processes.count")
			metric.SetDescription("Total number of processes in each state.")
			metric.SetUnit("{processes}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.processes.created",
		func(metric pdata.Metric) {
			metric.SetName("system.processes.created")
			metric.SetDescription("Total number of created processes.")
			metric.SetUnit("{processes}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
	// Cpu (CPU number starting at 0.)
	Cpu string
	// CPUState (Breakdown of CPU usage by type.)
	CPUState string
	// DiskDevice (Name of the disk.)
	DiskDevice string
	// DiskDirection (Direction of flow of bytes/opertations (read or write).)
	DiskDirection string
	// FilesystemDevice (Identifier of the filesystem.)
	FilesystemDevice string
	// FilesystemMode (Mountpoint mode such "ro", "rw", etc.)
	FilesystemMode string
	// FilesystemMountpoint (Mountpoint path.)
	FilesystemMountpoint string
	// FilesystemState (Breakdown of filesystem usage by type.)
	FilesystemState string
	// FilesystemType (Filesystem type, such as, "ext4", "tmpfs", etc.)
	FilesystemType string
	// MemState (Breakdown of memory usage by type.)
	MemState string
	// NetworkDevice (Name of the network interface.)
	NetworkDevice string
	// NetworkDirection (Direction of flow of bytes/opertations (receive or transmit).)
	NetworkDirection string
	// NetworkProtocol (Network protocol, e.g. TCP or UDP.)
	NetworkProtocol string
	// NetworkState (State of the network connection.)
	NetworkState string
	// PagingDevice (Name of the page file.)
	PagingDevice string
	// PagingDirection (Page In or Page Out.)
	PagingDirection string
	// PagingState (Breakdown of paging usage by type.)
	PagingState string
	// PagingType (Type of fault.)
	PagingType string
	// ProcessDirection (Direction of flow of bytes (read or write).)
	ProcessDirection string
	// ProcessState (Breakdown of CPU usage by type.)
	ProcessState string
	// ProcessesStatus (Breakdown status of the processes.)
	ProcessesStatus string
}{
	"cpu",
	"state",
	"device",
	"direction",
	"device",
	"mode",
	"mountpoint",
	"state",
	"type",
	"state",
	"device",
	"direction",
	"protocol",
	"state",
	"device",
	"direction",
	"state",
	"type",
	"direction",
	"state",
	"status",
}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels

// LabelCPUState are the possible values that the label "cpu.state" can have.
var LabelCPUState = struct {
	Idle      string
	Interrupt string
	Nice      string
	Softirq   string
	Steal     string
	System    string
	User      string
	Wait      string
}{
	"idle",
	"interrupt",
	"nice",
	"softirq",
	"steal",
	"system",
	"user",
	"wait",
}

// LabelDiskDirection are the possible values that the label "disk.direction" can have.
var LabelDiskDirection = struct {
	Read  string
	Write string
}{
	"read",
	"write",
}

// LabelFilesystemState are the possible values that the label "filesystem.state" can have.
var LabelFilesystemState = struct {
	Free     string
	Reserved string
	Used     string
}{
	"free",
	"reserved",
	"used",
}

// LabelMemState are the possible values that the label "mem.state" can have.
var LabelMemState = struct {
	Buffered          string
	Cached            string
	Inactive          string
	Free              string
	SlabReclaimable   string
	SlabUnreclaimable string
	Used              string
}{
	"buffered",
	"cached",
	"inactive",
	"free",
	"slab_reclaimable",
	"slab_unreclaimable",
	"used",
}

// LabelNetworkDirection are the possible values that the label "network.direction" can have.
var LabelNetworkDirection = struct {
	Receive  string
	Transmit string
}{
	"receive",
	"transmit",
}

// LabelNetworkProtocol are the possible values that the label "network.protocol" can have.
var LabelNetworkProtocol = struct {
	Tcp string
}{
	"tcp",
}

// LabelPagingDirection are the possible values that the label "paging.direction" can have.
var LabelPagingDirection = struct {
	PageIn  string
	PageOut string
}{
	"page_in",
	"page_out",
}

// LabelPagingState are the possible values that the label "paging.state" can have.
var LabelPagingState = struct {
	Cached string
	Free   string
	Used   string
}{
	"cached",
	"free",
	"used",
}

// LabelPagingType are the possible values that the label "paging.type" can have.
var LabelPagingType = struct {
	Major string
	Minor string
}{
	"major",
	"minor",
}

// LabelProcessDirection are the possible values that the label "process.direction" can have.
var LabelProcessDirection = struct {
	Read  string
	Write string
}{
	"read",
	"write",
}

// LabelProcessState are the possible values that the label "process.state" can have.
var LabelProcessState = struct {
	System string
	User   string
	Wait   string
}{
	"system",
	"user",
	"wait",
}

// LabelProcessesStatus are the possible values that the label "processes.status" can have.
var LabelProcessesStatus = struct {
	Blocked string
	Running string
}{
	"blocked",
	"running",
}
