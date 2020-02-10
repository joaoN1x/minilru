package db

type DbConfig struct {
	Type           string           `json:"type"`
	PostgresqlList []PostgresqlList `json:"postgresql_list"`
}

type PostgresqlList struct {
	Server    string                `json:"server"`
	Port      string                `json:"port"`
	Databases []PostgresqlDatabases `json:"databases"`
}

type PostgresqlDatabases struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Url struct {
	Id    int64  `json:"id"`
	Long  string `json:"long"`
	Short string `json:"short,omitempty"`
}

type UrlStats struct {
	UrlShort int64  `json:"id"`
	Today    string `json:"long"`
	Count    string `json:"today,omitempty"`
}
