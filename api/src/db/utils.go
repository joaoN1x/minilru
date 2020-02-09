package db

import (
	"encoding/json"
	"os"
)

/* GetConfig bring configurations related to databases to be in use */
func GetConfig() []DbConfig {
	var dbList, _ = os.LookupEnv("SERVICE_DB_CONFIG")
	bytes := []byte(dbList)
	var dbConf []DbConfig
	json.Unmarshal(bytes, &dbConf)
	return dbConf
}
