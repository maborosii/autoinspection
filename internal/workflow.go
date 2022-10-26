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
	var nodeStoreResults = make(metrics.MetricsMap)
	var metricsChan = make(chan *etl.QueryResult, 10)

	// 初始化映射关系
	nodeInToJob, nodeInToNodename := etl.QueryFromProm("init", global.PromQLForNodeInfo, global.PromClients[metrics.NODE_METRICS]).NodeInitInstanceMap()
	redisInToJob := etl.QueryFromProm("init", global.PromQLForRedisInfo, global.PromClients[metrics.REDIS_METRICS]).RedisInitInstanceMap()

	mergeMap := func(mObj ...map[string]string) map[string]string {
		newObj := map[string]string{}
		for _, m := range mObj {
			for k, v := range m {
				newObj[k] = v
			}
		}
		return newObj
	}
	allInToJob := mergeMap(nodeInToJob, redisInToJob)

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

	nodeStoreResults.MapToJobAndNodeName(allInToJob, nodeInToNodename)
	nodeStoreResults.MapToRules()
	nodeStoreResults.Notify()
}
