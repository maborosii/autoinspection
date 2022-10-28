package utils

// merge multi maps
func MergeMap(mObj ...map[string]string) map[string]string {
	newObj := map[string]string{}
	for _, m := range mObj {
		for k, v := range m {
			newObj[k] = v
		}
	}
	return newObj
}

// compute the increase rate from a to b
func IncreaseRate(a, b float32) float32 {
	if a <= 0 || b < 0 {
		return 0
	}
	return (b - a) / a * 100
}
