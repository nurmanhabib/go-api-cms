package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func Get(key string, defaults ...string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return ""
}

func GetInt(key string, defaults ...int) int {
	if val, ok := os.LookupEnv(key); ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return 0
}

func GetFloat64(key string, defaults ...float64) float64 {
	if val, ok := os.LookupEnv(key); ok {
		if floatVal, err := strconv.ParseFloat(val, 64); err == nil {
			return floatVal
		}
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return 0
}

func GetBool(key string, defaults ...bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		if boolVal, err := strconv.ParseBool(val); err == nil {
			return boolVal
		}
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return false
}

func GetDuration(key string, defaults ...time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok {
		if durationVal, err := time.ParseDuration(val); err == nil {
			return durationVal
		}
	}

	if len(defaults) > 0 {
		return defaults[0]
	}

	return 0
}

func GetSlice(key string, sep string, defaults ...string) []string {
	if val, ok := os.LookupEnv(key); ok {
		return strings.Split(val, sep)
	}

	if len(defaults) > 0 {
		return defaults
	}

	return []string{}
}

func GetSliceInt(key string, sep string, defaults ...[]int) []int {
	val, ok := os.LookupEnv(key)
	if !ok {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return []int{}
	}

	var results []int

	intSlices := strings.Split(val, sep)

	for _, intSlice := range intSlices {
		if intVal, err := strconv.Atoi(strings.TrimSpace(intSlice)); err == nil {
			results = append(results, intVal)
		} else {
			if len(defaults) > 0 {
				return defaults[0]
			}
			return []int{}
		}
	}

	return results
}
