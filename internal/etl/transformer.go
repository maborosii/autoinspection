package etl

import (
	"node_metrics_go/global"
	"strconv"

	"go.uber.org/zap"
)

// 将返回的结果进行转换
func ShuffleResult(series int, storeResults *MetricsMap) {
	defer WgReceiver.Done()
	for i := 0; i < series; i++ {
		queryResult := <-metricsChan
		results := queryResult.CleanValue(valuePattern)
		for _, result := range results {
			value, err := strconv.ParseFloat(result[1], 32)
			if err != nil {
				global.Logger.Error("Failed to convert value from string to float, err: ", zap.Error(err))
				value = 0
			}
			newValue := float32(value)
			switch queryResult.GetLabel() {
			case "cpu_usage_percents":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithCpuUsage(newValue))
			case "cpu_usage_percents_before_1day":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBefore1DayCpuUsage(newValue))
			case "cpu_usage_percents_before_1week":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBefore1WeekCpuUsage(newValue))
			case "mem_usage_percents":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithMemUsage(newValue))
			case "mem_usage_percents_before_1day":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBefore1DayMemUsage(newValue))
			case "mem_usage_percents_before_1week":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBefore1WeekMemUsage(newValue))
			case "disk_usage_percents":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithDiskUsage(newValue))
			case "disk_usage_percents_before_1day":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBefore1DayDiskUsage(newValue))
			case "disk_usage_percents_before_1week":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBefore1WeekDiskUsage(newValue))
			case "tcp_conn_counts":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithTcpConnUsage(newValue))
			case "tcp_conn_counts_before_1day":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBefore1DayTcpConnUsage(newValue))
			case "tcp_conn_counts_before_1week":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBefore1WeekTcpConnUsage(newValue))
			default:
				global.Logger.Info("NOT FOUND IN USE METRICS LABEL")
			}
		}
	}
	// close(notifyChan)
}
