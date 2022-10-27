package cmd

import (
	"node_metrics_go/global"
	"strings"

	"github.com/spf13/cobra"
)

var nodeConfig string

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "获取主机指标",
	Long:  nodeDesc,
	Run: func(cmd *cobra.Command, args []string) {
		global.ConfigPath = nodeConfig
		global.MetricsType = "node"
	},
}
var nodeDesc = strings.Join([]string{
	"该子命令支持获取主机指标"}, "\n")

func init() {
	// 参数分别表示，绑定的变量，参数长名(--str)，参数短名(-s)，默认内容，帮助信息
	nodeCmd.Flags().StringVarP(&nodeConfig, "config", "c", "configs", "请选择配置文件")
}
