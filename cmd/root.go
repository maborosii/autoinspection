package cmd

import (
	"fmt"
	"node_metrics_go/pkg/setting"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var confStruct = &setting.Config{}

var rootCmd = &cobra.Command{
	Use:   "ai",
	Short: "auto inspection in leyaoyao",
	Long: `ai 可不是人工智能，它是指 Auto Inspection （自动化巡检）
    用于快速生成超出用户期望的异常资源指标，应用于问题资源的快速定位`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				panic(err)
			}
			return
		}
	},
}

// 供主程序调用
func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	var vp = viper.New()
	if cfgFile != "" {
		// 使用 flag 标志中传递的配置文件
		vp.SetConfigFile(cfgFile)
	} else {
		localPath, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// 在 当前目录的 configs 目录下面查找名为 "config.toml" 的配置文件
		vp.AddConfigPath(path.Join(localPath, "configs"))
		vp.SetConfigName("config")
		vp.SetConfigType("toml")
	}
	// 读取匹配的环境变量
	// viper.AutomaticEnv()
	// 如果有配置文件，则读取它
	if err := vp.ReadInConfig(); err != nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
		panic(err)
	}
	conf := setting.NewSettingV(vp)
	if err := conf.ReadConfig(confStruct); err != nil {
		panic(err)
	}
}

// 初始化，将子命令嵌入到根命令中
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./configs/config.toml)")
	rootCmd.AddCommand(nodeCmd)
	rootCmd.AddCommand(redisCmd)
	rootCmd.AddCommand(kafkaCmd)
	rootCmd.AddCommand(rabbitMQCmd)
	rootCmd.AddCommand(elasticSearchCmd)
	rootCmd.AddCommand(jvmCmd)
	rootCmd.AddCommand(allCmd)
	rootCmd.AddCommand(versionCmd)
}
