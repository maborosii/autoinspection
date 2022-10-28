package etl

import "regexp"

// 用于 Node 初始化获取 instance，job，nodename 之间的映射关系
func (q *QueryResult) NodeInitInstanceMap() (map[string]string, map[string]string) {
	instanceToNodeName := make(map[string]string, 400)
	instanceToJob := make(map[string]string, 400)
	var re = regexp.MustCompile(mapPattenForNode)
	matched := re.FindAllStringSubmatch(q.value.String(), -1)
	for _, match := range matched {
		instanceToJob[match[1]] = match[2]
		instanceToNodeName[match[1]] = match[3]
	}
	return instanceToJob, instanceToNodeName
}

// 用于 Redis 初始化获取 instance，job 之间的映射关系
func (q *QueryResult) RedisInitInstanceMap() map[string]string {
	instanceToJob := make(map[string]string, 100)
	var re = regexp.MustCompile(mapPattenForRedis)
	matched := re.FindAllStringSubmatch(q.value.String(), -1)
	for _, match := range matched {
		instanceToJob[match[2]] = match[1]
	}
	return instanceToJob
}

// 用于 Kafka 初始化获取 instance，job 之间的映射关系
func (q *QueryResult) KafkaInitInstanceMap() map[string]string {
	instanceToJob := make(map[string]string, 100)
	var re = regexp.MustCompile(mapPattenForKafka)
	matched := re.FindAllStringSubmatch(q.value.String(), -1)
	for _, match := range matched {
		instanceToJob[match[1]] = match[2]
	}
	return instanceToJob
}
