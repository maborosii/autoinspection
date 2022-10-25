package setting

type MetricType string
type Config struct {
	// Address      string                   `toml:"address"`
	Endpoints    map[string]string        `toml:"endpoints" mapstructure:"endpoints"`
	TimeOut      int                      `toml:"timeout"`
	MonitorItems MonitorItems             `toml:"monitoritems"`
	Rules        map[string][]interface{} `toml:"rules" mapstructure:"rules"`
	LogConfig    *LogConf                 `toml:"logconfig"`
	MailConfig   *MailConf                `toml:"mailconfig"`
}
type MailConf struct {
	Host     string   `toml:"host"`
	UserName string   `toml:"username"`
	Password string   `toml:"password"`
	Port     int      `toml:"port"`
	To       []string `toml:"to"`
}
type LogConf struct {
	Level      string `toml:"level"`
	LogFile    string `toml:"logfile"`
	MaxSize    int    `toml:"maxsize"`
	MaxAge     int    `toml:"maxage"`
	MaxBackups int    `toml:"maxbackups"`
}

func (conf *Config) GetTimeOut() int {
	return conf.TimeOut
}

// func (conf *Config) GetMonitorItems() map[string]string {
// 	return conf.MonitorItems.ConvertToMap()
// }
func (conf *Config) GetLogConfig() *LogConf {
	return conf.LogConfig
}

func NewConfig() *Config {
	return &Config{}
}

type MonitorItems []*MonitorItem

type MonitorItem struct {
	Metrics string `toml:"metrics"`
	PromQL  string `toml:"promql"`
	// 从哪个数据源查询数据
	Endpoint []string `toml:"endpoint"`
}

// func (i MonitorItems) ConvertToMap() map[string]string {
// 	promQLs := make(map[string]string)
// 	for _, item := range i {
// 		promQLs[item.Metrics] = item.PromQL
// 	}
// 	return promQLs
// }

// 将配置文件数据映射到结构体中
func (s *Setting) ReadConfig(value interface{}) error {
	err := s.vp.Unmarshal(value)
	if err != nil {
		return err
	}
	return nil
}
