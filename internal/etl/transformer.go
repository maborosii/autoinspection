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
			case "cpu_usage_percents_before":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBeforeCpuUsage(newValue))
			case "mem_usage_percents":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithMemUsage(newValue))
			case "mem_usage_percents_before":
				storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithBeforeMemUsage(newValue))
			default:
				global.Logger.Info("Default")
			}
			// 添加jobname 和 nodename
			// TODO: 重复添加
			// storeResults.CreateOrModify(result[0], NewNodeMetrics(result[0]), WithNodeName(instanceToJob[result[0]]), WithNodeName(instanceToNodeName[result[0]]))
		}
	}
	close(notifyChan)
}
