package cache

type CacheConfig struct {
	Type string `json:"type"`
	RedisList []RedisList `json:"redis_list"`
}

type RedisList struct {
	Server string `json:"server"`
	Port string `json:"port"`
	Databases []RedisDatabases `json:"databases"`
}

type RedisDatabases struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
