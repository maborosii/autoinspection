package cmd

import (
	"node_metrics_go/global"
	"strings"

	"github.com/spf13/cobra"
)

var allConfig string

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "获取全部指标",
	Long:  allDesc,
	Run: func(cmd *cobra.Command, args []string) {
		global.ConfigPath = allConfig
		global.MetricsType = "all"
	},
}
var allDesc = strings.Join([]string{
	"该子命令支持获取全部指标"}, "\n")

func init() {
	// 参数分别表示，绑定的变量，参数长名(--str)，参数短名(-s)，默认内容，帮助信息
	allCmd.Flags().StringVarP(&allConfig, "config", "c", "configs", "请选择配置文件")
}