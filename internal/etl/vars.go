package etl

import (
	"sync"
)

// 初始化时获取映射关系
var instanceToNodeName = make(map[string]string, 400)
var instanceToJob = make(map[string]string, 400)
var promQLForMap = "node_uname_info - 0"

// 数据传输通道
var metricsChan = make(chan *QueryResult)

// 用于终止任务的通知通道
var notifyChan = make(chan struct{})

// 并发控制通道
var concurrencyChan = make(chan struct{}, 10)

var WgReceiver sync.WaitGroup

// 获取 job 和 intance, nodename 的映射关系
var mapPattenForNode = `(?m)instance="(.*?)".*\sjob="(.*?)".*\snodename="(.*?)".*\s=>\s.*$`

// 正则表达式匹配模式 --> 筛选出 intance 和 value
// 获取指标值的正则
var valuePattern = `(?m)instance="(.*?)".*\s=>\s*(\d*\.?\d{0,2}).*$`
