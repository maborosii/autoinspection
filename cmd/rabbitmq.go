package cmd

import (
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/internal"
	"strings"

	"github.com/spf13/cobra"
)

var rabbitMQCmd = &cobra.Command{
	Use:   "rabbitmq",
	Short: "获取 rabbitmq 指标",
	Long:  rabbitMQDesc,
	Run: func(cmd *cobra.Command, args []string) {
		mType := "rabbitmq"
		app := infra.NewBootApplication(confStruct)
		app.Run()
		global.MetricsType = mType
		internal.WorkFlow(mType)
	},
}
var rabbitMQDesc = strings.Join([]string{
	"该子命令支持获取 rabbitmq 指标"}, "\n")

func init() {
}
