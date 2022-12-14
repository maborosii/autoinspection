package internal

import (
	"fmt"
	"node_metrics_go/global"
	"node_metrics_go/internal/etl"
	"node_metrics_go/internal/metrics"
	"node_metrics_go/internal/utils"
	"sync"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"go.uber.org/zap"
)

// 意义不大 并发接收者并阻塞
var wgReceiver sync.WaitGroup

// 发送者组
var wgSender sync.WaitGroup
var metricsChan = make(chan *etl.QueryResult, 10)

// promql for get mapping of
var promQLForNodeInfo = "node_uname_info - 0"
var promQLForRedisInfo = "redis_instance_info - 0"
var promQLForKafkaInfo = "kafka_exporter_build_info - 0"
var promQLForRabbitMQInfo = "rabbitmq_exporter_build_info-0"
var promQLForElasticSearchInfo = "elasticsearch_clusterinfo_version_info-0"
var promQLForJVMInfo = "process_uptime_seconds-0"

// mian flow
func WorkFlow(mType string) {
	// store structure date for all king of metrics
	var storeResults = make(metrics.MetricsMap)

	//  get all maps of instance to job and nodeName
	allInToJob, nodeInToNodeName := initAllMap(mType)
	// get all metrics from proms
	extractMetrics(mType)

	wgReceiver.Add(1)
	// 转换数据
	go func() {
		defer wgReceiver.Done()
		// shuffle metrics from proms
		etl.ShuffleResult(metricsChan, &storeResults)
	}()
	wgReceiver.Wait()

	// bind job and nodeName for structure metrics data
	storeResults.MapToJobAndNodeName(allInToJob, nodeInToNodeName)
	// bind alert rules for structure metrics data by job
	storeResults.MapToRules()

	// execute metrics's alert rules
	storeResults.Notify()
}

// Get filter metrics type from external prometheus
func extractMetrics(metricType string) {
	filterType := metricType
	if metricType == "all" {
		filterType = ""
	}

	for _, monitorItem := range global.MonitorSetting.MonitorItems.Filter(filterType) {
		label := monitorItem.Metrics
		sql := monitorItem.PromQL
		eps := monitorItem.Endpoint
		global.Logger.Debug("query metrics from prom", zap.String("label", label), zap.String("promql", sql), zap.Strings("endpoints", eps))

		for _, ep := range eps {
			wgSender.Add(1)
			go func(label, sql string, queryApi v1.API) {
				defer wgSender.Done()
				etl.SendQueryResultToChan(label, sql, queryApi, metricsChan)
			}(label, sql, global.PromClients[ep])
		}
	}
	// close channel of message
	go func() {
		wgSender.Wait()
		close(metricsChan)
	}()
}

// 聚合所有指标的映射关系
func initAllMap(metricType string) (map[string]string, map[string]string) {
	switch metricType {
	case "all":
		return initMetricMap("")
	default:
		return initMetricMap(metricType)
	}
}

// 生成不同类型指标的 instance -> job 映射关系
func initMetricMap(metricType string) (map[string]string, map[string]string) {
	var instanceToJob, instanceToNodeName map[string]string

	// 从 monitorItems 找到对应的 endpoints 集合 (map[string]struct{})
	for k := range global.MonitorSetting.MonitorItems.FindAdaptEndpoints(metricType) {
		// 判断 endpoint 的指标类型
		v := global.MonitorSetting.Endpoints[k].Type
		switch v {
		case "node":
			a, b := etl.QueryFromProm(fmt.Sprintf("init node, endpoint: %s", k), promQLForNodeInfo, global.PromClients[k]).NodeInitInstanceMap()
			instanceToJob = utils.MergeMap(instanceToJob, a)
			instanceToNodeName = utils.MergeMap(instanceToNodeName, b)
		case "redis":
			c := etl.QueryFromProm(fmt.Sprintf("init redis, endpoint: %s", k), promQLForRedisInfo, global.PromClients[k]).RedisInitInstanceMap()
			instanceToJob = utils.MergeMap(instanceToJob, c)
		case "kafka":
			d := etl.QueryFromProm(fmt.Sprintf("init kafka, endpoint: %s", k), promQLForKafkaInfo, global.PromClients[k]).KafkaInitInstanceMap()
			instanceToJob = utils.MergeMap(instanceToJob, d)
		case "rabbitmq":
			e := etl.QueryFromProm(fmt.Sprintf("init rabbitMQ, endpoint: %s", k), promQLForRabbitMQInfo, global.PromClients[k]).RabbitMQInitInstanceMap()
			instanceToJob = utils.MergeMap(instanceToJob, e)
		case "es":
			f := etl.QueryFromProm(fmt.Sprintf("init elasticsearch, endpoint: %s", k), promQLForElasticSearchInfo, global.PromClients[k]).ElasticSearchInitInstanceMap()
			instanceToJob = utils.MergeMap(instanceToJob, f)
		case "jvm":
			f := etl.QueryFromProm(fmt.Sprintf("init jvm, endpoint: %s", k), promQLForJVMInfo, global.PromClients[k]).JVMInitInstanceMap()
			// map app_name to instance
			instanceToJob = utils.MergeMap(instanceToJob, f)
		default:
			global.Logger.Error("endpoint's metric type is not supported", zap.String("type", v))
		}
	}
	return instanceToJob, instanceToNodeName
}
