package etl

import (
	"node_metrics_go/global"
	"node_metrics_go/internal/metrics"
	"strconv"

	"go.uber.org/zap"
)

// 将返回的结果进行转换
func ShuffleResult(mChan <-chan *QueryResult, storeResults *metrics.MetricsMap) {
	initAllLabelHandleMap()

RES_CHANNEL_LOOP:
	for queryResult := range mChan {
		if queryResult == nil {
			continue RES_CHANNEL_LOOP
		}
		results := queryResult.CleanValue(valuePattern)
		label := queryResult.GetLabel()
		for _, result := range results {
			value, err := strconv.ParseFloat(result[1], 32)
			if err != nil {
				global.Logger.Error("Failed to convert value from string to float", zap.Error(err))
				value = 0
			}
			newValue := float32(value)
			if _, exists := metricsLabelHandlerMap[label]; !exists {
				global.Logger.Warn("NOT FOUND IN USE METRICS LABEL")
				continue RES_CHANNEL_LOOP
			}
			metricsLabelHandlerMap[label](result[0], newValue, storeResults)
		}
	}
}

type labelHandler func(string, float32, *metrics.MetricsMap)
type labelHandlerMap map[string]labelHandler

var metricsLabelHandlerMap = labelHandlerMap{}

func (m labelHandlerMap) registerHandler(label string, handler labelHandler) {
	if _, exists := m[label]; exists {
		return
	}
	m[label] = handler
}

func initAllLabelHandleMap() {
	initNodeLabelHandleMap()
	initRedisLabelHandler()
	initKafkaLabelHandler()
	initRabbitMQLabelHandler()
	initElasticSearchLabelHandler()
}

func initNodeLabelHandleMap() {
	// register node metrics handler
	metricsLabelHandlerMap.registerHandler("cpu_usage_percents", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithCPUUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("cpu_usage_percents_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithBefore1DayCPUUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("cpu_usage_percents_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithBefore1WeekCPUUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("mem_usage_percents", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithMemUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("mem_usage_percents_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithBefore1DayMemUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("mem_usage_percents_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithBefore1WeekMemUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("disk_usage_percents", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithDiskUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("disk_usage_percents_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithBefore1DayDiskUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("disk_usage_percents_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithBefore1WeekDiskUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("tcp_conn_counts", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithTCPConnUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("tcp_conn_counts_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithBefore1DayTCPConnUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("tcp_conn_counts_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewNodeMetrics(instance), metrics.WithBefore1WeekTCPConnUsage(value))
	})
}

func initRedisLabelHandler() {
	// register redis metrics handler
	metricsLabelHandlerMap.registerHandler("redis_conn_counts", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRedisMetrics(instance), metrics.WithRedisConnsUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("redis_conn_counts_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRedisMetrics(instance), metrics.WithBefore1DayRedisConnsUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("redis_conn_counts_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRedisMetrics(instance), metrics.WithBefore1WeekRedisConnsUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("redis_used_mem", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRedisMetrics(instance), metrics.WithRedisMemUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("redis_used_mem_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRedisMetrics(instance), metrics.WithBefore1DayRedisMemUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("redis_used_mem_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRedisMetrics(instance), metrics.WithBefore1WeekRedisMemUsage(value))
	})
}

func initKafkaLabelHandler() {
	// register kafka metrics handler
	metricsLabelHandlerMap.registerHandler("kafka_lag_sum", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewKafkaMetrics(instance), metrics.WithKafkaLagSumUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("kafka_lag_sum_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewKafkaMetrics(instance), metrics.WithBefore1DayKafkaLagSumUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("kafka_lag_sum_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewKafkaMetrics(instance), metrics.WithBefore1WeekKafkaLagSumUsage(value))
	})
}

func initRabbitMQLabelHandler() {
	// register rabbimq metrics handler
	metricsLabelHandlerMap.registerHandler("rabbitmq_running_nodes", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRabbitMQMetrics(instance), metrics.WithRabbitMQRunningNodesUsage(int8(value)))
	})
	metricsLabelHandlerMap.registerHandler("rabbitmq_running_nodes_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRabbitMQMetrics(instance), metrics.WithBefore1DayRabbitMQRunningNodesUsage(int8(value)))
	})
	metricsLabelHandlerMap.registerHandler("rabbitmq_running_nodes_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRabbitMQMetrics(instance), metrics.WithBefore1WeekRabbitMQRunningNodesUsage(int8(value)))
	})
	metricsLabelHandlerMap.registerHandler("rabbitmq_lag_sum", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRabbitMQMetrics(instance), metrics.WithRabbitMQLagSumUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("rabbitmq_lag_sum_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRabbitMQMetrics(instance), metrics.WithBefore1DayRabbitMQLagSumUsage(value))
	})
	metricsLabelHandlerMap.registerHandler("rabbitmq_lag_sum_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewRabbitMQMetrics(instance), metrics.WithBefore1WeekRabbitMQLagSumUsage(value))
	})
}

func initElasticSearchLabelHandler() {
	// register elasticsearch metrics handler
	metricsLabelHandlerMap.registerHandler("es_health_status", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewElasticSearchMetrics(instance), metrics.WithElasticSearchHeathStatus(int8(value)))
	})
	metricsLabelHandlerMap.registerHandler("es_health_status_before_1day", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewElasticSearchMetrics(instance), metrics.WithBefore1DayElasticSearchHealthStatus(int8(value)))
	})
	metricsLabelHandlerMap.registerHandler("es_health_status_before_1week", func(instance string, value float32, mm *metrics.MetricsMap) {
		mm.CreateOrModify(instance, metrics.NewElasticSearchMetrics(instance), metrics.WithBefore1WeekElasticSearchHealthStatus(int8(value)))
	})
}
