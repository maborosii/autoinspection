package main

import (
	"fmt"
	"log"
	"node_metrics_go/cmd"
	"node_metrics_go/global"
	"node_metrics_go/internal/etl"
	"node_metrics_go/pkg/logger"
	"node_metrics_go/pkg/setting"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func init() {
	// 获取配置文件
	var err error
	err = cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
	fmt.Println(global.ConfigPath)
	// 初始化配置
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// 初始化日志
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	// 手动将缓冲区日志内容刷入
	defer global.Logger.Sync()

	// 存储最终指标
	var nodeStoreResults = etl.NewNodeMetricsSlice()

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

	// writeResults := [][]string{}
	// for _, sr := range nodeStoreResults {
	// 	global.Logger.Info("get node of all metrics", zap.String("metrics", sr.Print()))
	// 	writeResults = append(writeResults, sr.ConvertToSlice())
	// }
	// fmt.Printf("%+v", writeResults)
}

// 根据配置文件位置读取配置文件
func setupSetting() error {
	setting, err := setting.NewSetting(global.ConfigPath)
	if err != nil {
		return err
	}
	err = setting.ReadConfig(&global.MonitorSetting)
	if err != nil {
		return err
	}

	return nil
}

// 初始化日志配置
func setupLogger() error {
	logConfig := global.MonitorSetting.GetLogConfig()
	err := logger.InitLogger(logConfig)
	if err != nil {
		return err
	}
	return nil
}
