package etl

const (
	CPU_LIMIT           string = "cpu exceeds the threshold"
	CPU_RATE_LIMIT      string = "cpu increase rate exceeds the threshold"
	MEM_LIMIT           string = "memory exceeds the threshold"
	MEM_RATE_LIMIT      string = "memory increase rate exceeds the threshold"
	DISK_LIMIT          string = "disk exceeds the threshold"
	DISK_RATE_LIMIT     string = "disk increase rate exceeds the threshold"
	TCP_CONN_LIMIT      string = "tcp conn counts exceeds the threshold"
	TCP_CONN_RATE_LIMIT string = "tcp conn counts increase rate exceeds the threshold"
)
