package cmd

import (
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/internal"
	"strings"

	"github.com/spf13/cobra"
)

// var redisConfig string

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "获取 redis 指标",
	Long:  redisDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// global.ConfigPath = redisConfig
		// metricsType = "redis"
		// internal.WorkFlow("redis")
		mType := "redis"
		app := infra.NewBootApplication(confStruct)
		app.Run()
		global.MetricsType = mType
		internal.WorkFlow(mType)
	},
}
var redisDesc = strings.Join([]string{
	"该子命令支持获取 redis 指标"}, "\n")

func init() {
	// 参数分别表示，绑定的变量，参数长名(--str)，参数短名(-s)，默认内容，帮助信息
	// redisCmd.Flags().StringVarP(&redisConfig, "config", "c", "configs", "请选择配置文件")
}
