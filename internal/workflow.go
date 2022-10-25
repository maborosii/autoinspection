package internal

import (
	"fmt"
	"node_metrics_go/global"
	"node_metrics_go/internal/etl"
	"node_metrics_go/internal/metrics"
	"sync"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// 意义不大 并发接收者并阻塞
var wgReceiver sync.WaitGroup

// 发送者组
var wgSender sync.WaitGroup

// 数据传输通道

func WorkFlow() {
	// 存储最终指标
	var nodeStoreResults = make(metrics.MetricsMap)
	var metricsChan = make(chan *etl.QueryResult)
	nodeInToJob, nodeInToNodename := etl.QueryFromProm("init", global.PromQLForMap, global.PromClients["node"]).InitInstanceMap()

	// 查询具体指标
	// for label, sql := range global.MonitorSetting.GetMonitorItems() {
	for _, monitorItem := range global.MonitorSetting.MonitorItems {
		label := monitorItem.Metrics
		sql := monitorItem.PromQL
		eps := monitorItem.Endpoint
		fmt.Println(label, sql, eps)
		for _, ep := range eps {
			wgSender.Add(1)
			go func(label, sql string, queryApi v1.API) {
				defer wgSender.Done()
				etl.SendQueryResultToChan(label, sql, queryApi, metricsChan)
			}(label, sql, global.PromClients[ep])
		}
	}
	// 关闭 消息体通道
	go func() {
		wgSender.Wait()
		close(metricsChan)
	}()

	wgReceiver.Add(1)
	// 转换数据
	go func() {
		defer wgReceiver.Done()
		etl.ShuffleResult(metricsChan, &nodeStoreResults)
	}()
	wgReceiver.Wait()

	nodeStoreResults.MapToJobAndNodeName(nodeInToJob, nodeInToNodename)
	nodeStoreResults.MapToRules()
	nodeStoreResults.Notify()
}
