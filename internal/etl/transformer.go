package etl

import (
	"node_metrics_go/global"
	"node_metrics_go/internal/metrics"
	"strconv"

	"go.uber.org/zap"
)

// 将返回的结果进行转换
func ShuffleResult(mChan <-chan *QueryResult, storeResults *metrics.MetricsMap) {
RES_CHANNEL_LOOP:
	for queryResult := range mChan {
		if queryResult == nil {
			continue RES_CHANNEL_LOOP
		}
		results := queryResult.CleanValue(valuePattern)
		for _, result := range results {
			value, err := strconv.ParseFloat(result[1], 32)
			if err != nil {
				global.Logger.Error("Failed to convert value from string to float", zap.Error(err))
				value = 0
			}
			newValue := float32(value)
			switch queryResult.GetLabel() {
			// node metrics
			case "cpu_usage_percents":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithCPUUsage(newValue))
			case "cpu_usage_percents_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithBefore1DayCPUUsage(newValue))
			case "cpu_usage_percents_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithBefore1WeekCPUUsage(newValue))
			case "mem_usage_percents":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithMemUsage(newValue))
			case "mem_usage_percents_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithBefore1DayMemUsage(newValue))
			case "mem_usage_percents_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithBefore1WeekMemUsage(newValue))
			case "disk_usage_percents":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithDiskUsage(newValue))
			case "disk_usage_percents_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithBefore1DayDiskUsage(newValue))
			case "disk_usage_percents_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithBefore1WeekDiskUsage(newValue))
			case "tcp_conn_counts":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithTCPConnUsage(newValue))
			case "tcp_conn_counts_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithBefore1DayTCPConnUsage(newValue))
			case "tcp_conn_counts_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewNodeMetrics(result[0]), metrics.WithBefore1WeekTCPConnUsage(newValue))

			// redis metrics
			case "redis_conn_counts":
				storeResults.CreateOrModify(result[0], metrics.NewRedisMetrics(result[0]), metrics.WithRedisConnsUsage(newValue))
			case "redis_conn_counts_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewRedisMetrics(result[0]), metrics.WithBefore1DayRedisConnsUsage(newValue))
			case "redis_conn_counts_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewRedisMetrics(result[0]), metrics.WithBefore1WeekRedisConnsUsage(newValue))
			case "redis_used_mem":
				storeResults.CreateOrModify(result[0], metrics.NewRedisMetrics(result[0]), metrics.WithRedisMemUsage(newValue))
			case "redis_used_mem_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewRedisMetrics(result[0]), metrics.WithBefore1DayRedisMemUsage(newValue))
			case "redis_used_mem_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewRedisMetrics(result[0]), metrics.WithBefore1WeekRedisMemUsage(newValue))

			// kafka metrics
			case "kafka_lag_sum":
				storeResults.CreateOrModify(result[0], metrics.NewKafkaMetrics(result[0]), metrics.WithKafkaLagSumUsage(newValue))
			case "kafka_lag_sum_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewKafkaMetrics(result[0]), metrics.WithBefore1DayKafkaLagSumUsage(newValue))
			case "kafka_lag_sum_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewKafkaMetrics(result[0]), metrics.WithBefore1WeekKafkaLagSumUsage(newValue))

			// kafka metrics
			case "rabbitmq_running_nodes":
				storeResults.CreateOrModify(result[0], metrics.NewRabbitMQMetrics(result[0]), metrics.WithRabbitMQRunningNodesUsage(int8(newValue)))
			case "rabbitmq_running_nodes_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewRabbitMQMetrics(result[0]), metrics.WithBefore1DayRabbitMQRunningNodesUsage(int8(newValue)))
			case "rabbitmq_running_nodes_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewRabbitMQMetrics(result[0]), metrics.WithBefore1WeekRabbitMQRunningNodesUsage(int8(newValue)))
			case "rabbitmq_lag_sum":
				storeResults.CreateOrModify(result[0], metrics.NewRabbitMQMetrics(result[0]), metrics.WithRabbitMQLagSumUsage(newValue))
			case "rabbitmq_lag_sum_before_1day":
				storeResults.CreateOrModify(result[0], metrics.NewRabbitMQMetrics(result[0]), metrics.WithBefore1DayRabbitMQLagSumUsage(newValue))
			case "rabbitmq_lag_sum_before_1week":
				storeResults.CreateOrModify(result[0], metrics.NewRabbitMQMetrics(result[0]), metrics.WithBefore1WeekRabbitMQLagSumUsage(newValue))

			default:
				global.Logger.Info("NOT FOUND IN USE METRICS LABEL")
			}
		}
	}
}
