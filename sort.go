package main

type Processes []*ProcessStatus

// 获取此 slice 的长度
func (p Processes) Len() int {
	return len(p)
}

// 根据元素的最大打开文件数
func (p Processes) Less(i, j int) bool {
	return p[i].OpenFiles > p[j].OpenFiles
}

// 交换数据
func (p Processes) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
