package main

type Processes []*ProcessStatus

// 获取此 slice 的长度
func (p Processes) Len() int {
	return len(p)
}

// 根据元素的最大打开文件数排序
func (p Processes) Less(i, j int) bool {
	return p[i].OpenFiles > p[j].OpenFiles
}

// 交换数据
func (p Processes) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ProcessSortByCPU struct {
	Processes
}

// 按照CPU使用率排序
func (p ProcessSortByCPU) Less(i, j int) bool {
	return p.Processes[i].CPU > p.Processes[j].CPU
}

type ProcessSortByMem struct {
	Processes
}

// 按照内存使用率排序
func (p ProcessSortByMem) Less(i, j int) bool {
	return p.Processes[i].Mem > p.Processes[j].Mem
}

type ProcessSortByConnections struct {
	Processes
}

// 按照连接数使用率排序
func (p ProcessSortByConnections) Less(i, j int) bool {
	return p.Processes[i].Connections > p.Processes[j].Connections
}
