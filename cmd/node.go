package cmd

import (
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/internal"
	"strings"

	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "获取 node 指标",
	Long:  nodeDesc,
	Run: func(cmd *cobra.Command, args []string) {
		mType := "node"
		app := infra.NewBootApplication(confStruct)
		app.Run()
		global.MetricsType = mType
		internal.WorkFlow(mType)
	},
}
var nodeDesc = strings.Join([]string{
	"该子命令支持获取 node 指标"}, "\n")

func init() {
}
