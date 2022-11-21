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
	// REDIS_MEM_LIMIT            string = "redis memory exceeds the threshold"
	// REDIS_MEM_RATE_LIMIT_1DAY  string = "redis memory increase rate exceeds the threshold in one day"
	// REDIS_MEM_RATE_LIMIT_1WEEK string = "redis memory increase rate exceeds the threshold in one week"
	// cn
	REDIS_CONN_LIMIT            string = "[redis][连接数] 瞬时值超过阈值 |数值|"
	REDIS_CONN_RATE_LIMIT_1DAY  string = "[redis][连接数] 一天增长率超过阈值 |百分比|"
	REDIS_CONN_RATE_LIMIT_1WEEK string = "[redis][连接数] 一周增长率超过阈值 |百分比|"
	REDIS_MEM_LIMIT             string = "[redis][内存] 瞬时值超过阈值 |百分比|"
	REDIS_MEM_RATE_LIMIT_1DAY   string = "[redis][内存] 一天增长率超过阈值 |百分比|"
	REDIS_MEM_RATE_LIMIT_1WEEK  string = "[redis][内存] 一周增长率超过阈值 |百分比|"

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

	/*rabbitmq alert message
	 */
	// en
	// RABBITMQ_RUNNING_NODES               string = "rabbitmq running nodes not equals the threshold"
	// RABBITMQ_RUNNING_NODES_CHANGE_1DAY   string = "rabbitmq running nodes changed in one day"
	// RABBITMQ_RUNNING_NODES_CHANGE_1WEEK  string = "rabbitmq running nodes changed in one week"
	// RABBITMQ_LAG_SUM_LIMIT               string = "rabbitmq lag sum exceeds the threshold"
	// RABBITMQ_LAG_SUM_RATE_LIMIT_1DAY     string = "rabbitmq lag sum increase rate exceeds the threshold in one day"
	// RABBITMQ_LAG_SUM_RATE_LIMIT_1WEEK    string = "rabbitmq lag sum increase rate exceeds the threshold in one week"
	// cn
	RABBITMQ_RUNNING_NODES               string = "[rabbitmq][运行节点数] 瞬时值异常 |数值|"
	RABBITMQ_RUNNING_NODES_CHANGED_1DAY  string = "[rabbitmq][运行节点数] 一天变化数目异常 |数值|"
	RABBITMQ_RUNNING_NODES_CHANGED_1WEEK string = "[rabbitmq][运行节点数] 一周变化数目异常 |数值|"
	RABBITMQ_LAG_SUM_LIMIT               string = "[rabbitmq][总堆积量] 瞬时值超过阈值 |数值|"
	RABBITMQ_LAG_SUM_RATE_LIMIT_1DAY     string = "[rabbitmq][总堆积量] 一天增长率超过阈值 |百分比|"
	RABBITMQ_LAG_SUM_RATE_LIMIT_1WEEK    string = "[rabbitmq][总堆积量] 一周增长率超过阈值 |百分比|"

	/*elasticsearch alert message
	 */
	// en
	// ELASTICSEARCH_HEALTH_STATUS               string = "elasticsearch is not healthy"
	// ELASTICSEARCH_HEALTH_STATUS_CHANGED_1DAY  string = "elasticsearch health status changed in one day"
	// ELASTICSEARCH_HEALTH_STATUS_CHANGED_1WEEK string = "elasticsearch health status changed in one week"
	// cn
	ELASTICSEARCH_HEALTH_STATUS               string = "[elasticsearch][健康值] 瞬时值异常 |数值|"
	ELASTICSEARCH_HEALTH_STATUS_CHANGED_1DAY  string = "[elasticsearch][健康值] 一天变化异常 |数值|"
	ELASTICSEARCH_HEALTH_STATUS_CHANGED_1WEEK string = "[elasticsearch][健康值] 一周变化异常 |数值|"
	/*jvm alert message
	 */
	// en
	// JVM_BLOCKED_THREAD_COUNT  string = "jvm blocked thread count exceed the threshold"
	// JVM_GARBAGE_COLLECT_TIME  string = "jvm gc time exceed the threshold"
	// JVM_GARBAGE_COLLECT_COUNT string = "jvm gc count exceed the threshold"
	// cn
	JVM_BLOCKED_THREAD_COUNT  string = "[jvm][阻塞线程数] 瞬时值异常 |数值|"
	JVM_GARBAGE_COLLECT_TIME  string = "[jvm][gc 时间] 瞬时值异常 |数值|"
	JVM_GARBAGE_COLLECT_COUNT string = "[jvm][gc 次数] 瞬时值异常 |数值|"
)
const (
	NODE_METRICS          string = "node"
	REDIS_METRICS         string = "redis"
	KAFKA_METRICS         string = "kafka"
	RABBITMQ_METRICS      string = "rabbitmq"
	ELASTICSEARCH_METRICS string = "es"
	JVM_METRICS           string = "jvm"
)
