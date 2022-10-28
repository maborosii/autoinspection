package cmd

import (
	"node_metrics_go/global"
	"strings"

	"github.com/spf13/cobra"
)

var kafkaConfig string

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "获取 kafka 指标",
	Long:  kafkaDesc,
	Run: func(cmd *cobra.Command, args []string) {
		global.ConfigPath = kafkaConfig
		global.MetricsType = "kafka"
	},
}
var kafkaDesc = strings.Join([]string{
	"该子命令支持获取 kafka 指标"}, "\n")

func init() {
	// 参数分别表示，绑定的变量，参数长名(--str)，参数短名(-s)，默认内容，帮助信息
	kafkaCmd.Flags().StringVarP(&kafkaConfig, "config", "c", "configs", "请选择配置文件")
}
