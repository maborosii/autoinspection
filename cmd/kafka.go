package cmd

import (
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/internal"
	"strings"

	"github.com/spf13/cobra"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "获取 kafka 指标",
	Long:  kafkaDesc,
	Run: func(cmd *cobra.Command, args []string) {
		mType := "kafka"
		app := infra.NewBootApplication(confStruct)
		app.Run()
		global.MetricsType = mType
		internal.WorkFlow(mType)
	},
}
var kafkaDesc = strings.Join([]string{
	"该子命令支持获取 kafka 指标"}, "\n")

func init() {
}
