package metrics

const (
	// node alert message
	CPU_LIMIT                 string = "cpu exceeds the threshold"
	CPU_RATE_LIMIT_1DAY       string = "cpu increase rate exceeds the threshold in one day"
	CPU_RATE_LIMIT_1WEEK      string = "cpu increase rate exceeds the threshold in one week"
	MEM_LIMIT                 string = "memory exceeds the threshold"
	MEM_RATE_LIMIT_1DAY       string = "memory increase rate exceeds the threshold in one day"
	MEM_RATE_LIMIT_1WEEK      string = "memory increase rate exceeds the threshold in one week"
	DISK_LIMIT                string = "disk exceeds the threshold"
	DISK_RATE_LIMIT_1DAY      string = "disk increase rate exceeds the threshold in one day"
	DISK_RATE_LIMIT_1WEEK     string = "disk increase rate exceeds the threshold in one week"
	TCP_CONN_LIMIT            string = "tcp conn counts exceeds the threshold"
	TCP_CONN_RATE_LIMIT_1DAY  string = "tcp conn counts increase rate exceeds the threshold in one day"
	TCP_CONN_RATE_LIMIT_1WEEK string = "tcp conn counts increase rate exceeds the threshold in one week"

	// redis alert message
	REDIS_CONN_LIMIT            string = "redis conn counts exceeds the threshold"
	REDIS_CONN_RATE_LIMIT_1DAY  string = "redis conn counts increase rate exceeds the threshold in one day"
	REDIS_CONN_RATE_LIMIT_1WEEK string = "redis conn counts increase rate exceeds the threshold in one week"
)
const (
	// NODE_METRICS          global.MetricType  = "node"
	// REDIS_METRICS         global.MetricType  = "redis"
	// KAFKA_METRICS         global.MetricType  = "kafka"
	// RABBITMQ_METRICS      global.MetricType  = "rabbitmq"
	// ELASTICSEARCH_METRICS global.MMetricType = "node"
	NODE_METRICS          string = "node"
	REDIS_METRICS         string = "redis"
	KAFKA_METRICS         string = "kafka"
	RABBITMQ_METRICS      string = "rabbitmq"
	ELASTICSEARCH_METRICS string = "es"
)
