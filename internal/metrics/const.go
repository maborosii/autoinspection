package metrics

const (
	/* node alert message
	 */
	// en
	// CPU_LIMIT                 string = "cpu exceeds the threshold"
	// CPU_RATE_LIMIT_1DAY       string = "cpu increase rate exceeds the threshold in one day"
	// CPU_RATE_LIMIT_1WEEK      string = "cpu increase rate exceeds the threshold in one week"
	// MEM_LIMIT                 string = "memory exceeds the threshold"
	// MEM_RATE_LIMIT_1DAY       string = "memory increase rate exceeds the threshold in one day"
	// MEM_RATE_LIMIT_1WEEK      string = "memory increase rate exceeds the threshold in one week"
	// DISK_LIMIT                string = "disk exceeds the threshold"
	// DISK_RATE_LIMIT_1DAY      string = "disk increase rate exceeds the threshold in one day"
	// DISK_RATE_LIMIT_1WEEK     string = "disk increase rate exceeds the threshold in one week"
	// TCP_CONN_LIMIT            string = "tcp conn counts exceeds the threshold"
	// TCP_CONN_RATE_LIMIT_1DAY  string = "tcp conn counts increase rate exceeds the threshold in one day"
	// TCP_CONN_RATE_LIMIT_1WEEK string = "tcp conn counts increase rate exceeds the threshold in one week"
	// cn
	CPU_LIMIT                 string = "[node][cpu] 瞬时使用率超过阈值 |百分比|"
	CPU_RATE_LIMIT_1DAY       string = "[node][cpu] 一天使用增长率超过阈值 |百分比|"
	CPU_RATE_LIMIT_1WEEK      string = "[node][cpu] 一周使用增长率超过阈值 |百分比|"
	MEM_LIMIT                 string = "[node][内存] 瞬时使用率超过阈值 |百分比|"
	MEM_RATE_LIMIT_1DAY       string = "[node][内存] 一天使用增长率超过阈值 |百分比|"
	MEM_RATE_LIMIT_1WEEK      string = "[node][内存] 一周使用增长率超过阈值 |百分比|"
	DISK_LIMIT                string = "[node][磁盘] 瞬时使用率超过阈值 |百分比|"
	DISK_RATE_LIMIT_1DAY      string = "[node][磁盘] 一天使用增长率超过阈值 |百分比|"
	DISK_RATE_LIMIT_1WEEK     string = "[node][磁盘] 一周使用增长率超过阈值 |百分比|"
	TCP_CONN_LIMIT            string = "[node][ tcp 连接数] 瞬时值超过阈值 |数值|"
	TCP_CONN_RATE_LIMIT_1DAY  string = "[node][ tcp 连接数] 一天增长率超过阈值 |百分比|"
	TCP_CONN_RATE_LIMIT_1WEEK string = "[node][ tcp 连接数] 一周增长率超过阈值 |百分比|"

	/*redis alert message
	 */
	// en
	// REDIS_CONN_LIMIT            string = "redis conn counts exceeds the threshold"
	// REDIS_CONN_RATE_LIMIT_1DAY  string = "redis conn counts increase rate exceeds the threshold in one day"
	// REDIS_CONN_RATE_LIMIT_1WEEK string = "redis conn counts increase rate exceeds the threshold in one week"
	// cn
	REDIS_CONN_LIMIT            string = "[redis][连接数] 瞬时值超过阈值 |数值|"
	REDIS_CONN_RATE_LIMIT_1DAY  string = "[redis][连接数] 一天增长率超过阈值 |百分比|"
	REDIS_CONN_RATE_LIMIT_1WEEK string = "[redis][连接数] 一周增长率超过阈值 |百分比|"

	/*kafka alert message
	 */
	// en
	// KAFKA_LAG_SUM_LIMIT            string = "kafka lag sum exceeds the threshold"
	// KAFKA_LAG_SUM_RATE_LIMIT_1DAY  string = "kafka lag sum increase rate exceeds the threshold in one day"
	// KAFKA_LAG_SUM_RATE_LIMIT_1WEEK string = "kafka lag sum increase rate exceeds the threshold in one week"
	// cn
	KAFKA_LAG_SUM_LIMIT            string = "[kafka][总堆积量] 瞬时值超过阈值 |数值|"
	KAFKA_LAG_SUM_RATE_LIMIT_1DAY  string = "[kafka][总堆积量] 一天增长率超过阈值 |百分比|"
	KAFKA_LAG_SUM_RATE_LIMIT_1WEEK string = "[kafka][总堆积量] 一周增长率超过阈值 |百分比|"
)
const (
	NODE_METRICS          string = "node"
	REDIS_METRICS         string = "redis"
	KAFKA_METRICS         string = "kafka"
	RABBITMQ_METRICS      string = "rabbitmq"
	ELASTICSEARCH_METRICS string = "es"
)
