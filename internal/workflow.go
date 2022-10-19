package internal

import (
	"fmt"
	"node_metrics_go/global"
	"node_metrics_go/internal/etl"
	"node_metrics_go/internal/metrics"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func WorkFlow() {
	// 存储最终指标
	var nodeStoreResults = make(metrics.MetricsMap)
	nodeInToJob, nodeInToNodename := etl.QueryFromProm("init", global.PromQLForMap, global.PromClients["node"]).InitInstanceMap()

	// 查询具体指标
	for label, sql := range global.MonitorSetting.GetMonitorItems() {
		fmt.Println(label, sql)
		go func(label, sql string, queryApi v1.API) {
			etl.SendQueryResultToChan(label, sql, queryApi)
		}(label, sql, global.PromClients["node"])
	}

	etl.WgReceiver.Add(1)
	// 转换数据
	go etl.ShuffleResult(len(global.MonitorSetting.GetMonitorItems()), &nodeStoreResults)
	etl.WgReceiver.Wait()

	nodeStoreResults.MapToJobAndNodeName(nodeInToJob, nodeInToNodename)
	nodeStoreResults.MapToRules()
	nodeStoreResults.Notify()
}
