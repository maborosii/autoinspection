package etl

import (
	"node_metrics_go/global"
)

// 将返回的结果进行转换
func ShuffleResult(series int, storeResults *NodeMetricsSlice) {
	defer WgReceiver.Done()
	for i := 0; i < series; i++ {
		queryResult := <-metricsChan
		results := queryResult.CleanValue(valuePattern)
		for _, result := range results {
			switch queryResult.GetLabel() {
			case "cpu_usage_percents":
				storeResults.CreateOrModifyStoreResults(result[0], WithCpuUsage(result[1]))
			case "cpu_usage_percents_before":
				storeResults.CreateOrModifyStoreResults(result[0], WithBeforeCpuUsage(result[1]))
			case "mem_usage_percents":
				storeResults.CreateOrModifyStoreResults(result[0], WithMemUsage(result[1]))
			case "mem_usage_percents_before":
				storeResults.CreateOrModifyStoreResults(result[0], WithBeforeMemUsage(result[1]))
			default:
				global.Logger.Info("Default")
			}
			// 添加jobname 和 nodename
			storeResults.CreateOrModifyStoreResults(result[0], WithNodeJob(instanceToJob[result[0]]), WithNodeName(instanceToNodeName[result[0]]))
		}
	}
	close(notifyChan)
}
