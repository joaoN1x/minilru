package cache

import (
	"encoding/json"
	"os"

	"github.com/joaoN1x/minilru/src/debugger"
)

/* GetConfig bring configurations related to databases to be in use */
func GetConfig() []CacheConfig {
	var cacheList, ok = os.LookupEnv("SERVICE_CACHE_CONFIG")
	if !ok {
		debugger.Log("error", "Get Cache service ENV", nil)
	}
	bytes := []byte(cacheList)
	var cacheConf []CacheConfig
	json.Unmarshal(bytes, &cacheConf)
	return cacheConf
}
