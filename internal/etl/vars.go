package etl

// 用于终止任务的通知通道
// var notifyChan = make(chan struct{})

// 并发控制通道
var concurrencyChan = make(chan struct{}, 10)

// 获取 job 和 intance, nodename 的映射关系
var mapPattenForNode = `(?m)instance="(.*?)".*\sjob="(.*?)".*\snodename="(.*?)".*\s=>\s.*$`
var mapPattenForRedis = `(?m)group="(.*?)".*\sinstance="(.*?)".*\s=>\s.*$`
var mapPattenForKafka = `(?m)instance="(.*?)".*\sjob="(.*?)".*\s=>\s.*$`
var mapPattenForRabbitMQ = `(?m)group="(.*?)".*\sinstance="(.*?)".*\s=>\s.*$`
var mapPattenForElasticSearch = `(?m)cluster="(.*?)".*\sinstance="(.*?)".*\s=>\s.*$`

// 正则表达式匹配模式 --> 筛选出 intance 和 value
// 获取指标值的正则
var valuePattern = `(?m)instance="(.*?)".*\s=>\s*(-?\d*\.?\d{0,2}).*$`
