package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/joaoN1x/minilru/src/debugger"
)

var (
	cacheConnection *redis.Client
	config          = GetConfig()
)

func initRedis() *redis.Client {

	connectionString := fmt.Sprintf("%s:%s", config[0].RedisList[0].Server, config[0].RedisList[0].Port)

	connectionDb, connectionDbErr := strconv.Atoi(config[0].RedisList[0].Databases[0].Name)
	if connectionDbErr != nil {
		debugger.Log("error", "Error", connectionDbErr)
	}
	redisConnection := redis.NewClient(&redis.Options{
		Addr:     connectionString,
		Password: string(config[0].RedisList[0].Databases[0].Password),
		DB:       connectionDb,
	})
	return redisConnection
}

func GetUrl(key string) string {

	if cacheConnection == nil {
		cacheConnection = initRedis()
	}

	data, err := cacheConnection.Get(key).Result()
	if err == redis.Nil {
		debugger.Log("warning", string(key)+" does not exist. ", nil)
	} else if err != nil {
		debugger.Log("warning", "Error ", err)
	} else {
		return data
	}

	return ""

}

func SetUrl(key string, value string) {

	if cacheConnection == nil {
		cacheConnection = initRedis()
	}

	err := cacheConnection.Set(key, value, 0).Err()
	if err != nil {
		debugger.Log("error", "Error ", err)
	}

}
