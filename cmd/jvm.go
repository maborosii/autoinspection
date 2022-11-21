package cmd

import (
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/internal"
	"strings"

	"github.com/spf13/cobra"
)

var jvmCmd = &cobra.Command{
	Use:   "jvm",
	Short: "获取 jvm 指标",
	Long:  jvmDesc,
	Run: func(cmd *cobra.Command, args []string) {
		mType := "jvm"
		app := infra.NewBootApplication(confStruct)
		app.Run()
		global.MetricsType = mType
		internal.WorkFlow(mType)
	},
}
var jvmDesc = strings.Join([]string{
	"该子命令支持获取 jvm 指标"}, "\n")

func init() {
}
