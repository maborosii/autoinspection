package main

import (
	"fmt"
	"node_metrics_go/cmd"
	"node_metrics_go/infra"
	"node_metrics_go/infra/starters"
)

func main() {
	// 获取配置文件
	var err error
	err = cmd.Execute()
	if err != nil {
		panic(fmt.Sprintf("cmd.Execute err: %v", err))
	}
	// conf, err := setting.NewSetting(global.ConfigPath)
	// if err != nil {
	// 	panic(fmt.Sprintf("get config from %s, occurred err; %s", global.ConfigPath, err))
	// }
	// confStruct := &setting.Config{}
	// if err = conf.ReadConfig(confStruct); err != nil {
	// 	panic(err)
	// }
	// app := infra.NewBootApplication(confStruct)
	// app.Run()
	// internal.WorkFlow(global.MetricsType)
}

func init() {
	infra.Register(&starters.LogStarter{}, &starters.PromClientsStarter{}, &starters.MonitorStarter{}, &starters.RulesStarter{}, &starters.MailStarter{})
}
