package cmd

import (
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/internal"
	"strings"

	"github.com/spf13/cobra"
)

var elasticSearchCmd = &cobra.Command{
	Use:   "es",
	Short: "获取 elasticsearch 指标",
	Long:  elasticSearchDesc,
	Run: func(cmd *cobra.Command, args []string) {
		mType := "es"
		app := infra.NewBootApplication(confStruct)
		app.Run()
		global.MetricsType = mType
		internal.WorkFlow(mType)
	},
}
var elasticSearchDesc = strings.Join([]string{
	"该子命令支持获取 elasticsearch 指标"}, "\n")

func init() {
}
