package main

import (
	"fmt"
	"log"
	"node_metrics_go/cmd"
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/infra/starters"
	"node_metrics_go/internal/etl"
	"node_metrics_go/pkg/setting"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func main() {
	// 获取配置文件
	var err error
	err = cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
	conf, err := setting.NewSetting(global.ConfigPath)
	if err != nil {
		panic(fmt.Sprintf("get config from %s, occurred err; %s", global.ConfigPath, err))
	}
	confStruct := &setting.Config{}
	if err = conf.ReadConfig(confStruct); err != nil {
		panic(err)
	}
	app := infra.NewBootApplication(confStruct)
	app.Run()

	// 存储最终指标
	var nodeStoreResults = make(etl.MetricsMap)
	// var nodeStoreResults = etl.NewNodeMetricsSlice()

	// prom客户端api
	queryApi := etl.ClientForProm(global.MonitorSetting.GetAddress())

	// 初始化映射关系
	etl.QueryFromProm("init", global.PromQLForMap, queryApi).InitInstanceMap()

	// 查询具体指标
	for label, sql := range global.MonitorSetting.GetMonitorItems() {
		fmt.Println(label, sql)
		go func(label, sql string, queryApi v1.API) {
			etl.SendQueryResultToChan(label, sql, queryApi)
		}(label, sql, queryApi)
	}

	etl.WgReceiver.Add(1)
	// 转换数据
	go etl.ShuffleResult(len(global.MonitorSetting.GetMonitorItems()), &nodeStoreResults)
	etl.WgReceiver.Wait()

	nodeStoreResults.MapToJobAndNodeName()
	nodeStoreResults.MapToRules()
	// for _, sr := range nodeStoreResults {
	// 	global.Logger.Info("get node of all metrics", zap.String("metrics", sr.Print()))
	// }
	nodeStoreResults.Notify()

}

func init() {
	infra.Register(&starters.LogStarter{}, &starters.MonitorStarter{}, &starters.RulesStarter{})
}
