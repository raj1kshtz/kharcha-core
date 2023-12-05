package db

const (
	Host            = "DB_HOST"
	Port            = "DB_PORT"
	Database        = "DB_DATABASE"
	MaxPoolSize     = "DB_MAX_POOL_SIZE"
	MinPoolSize     = "DB_MIN_POOL_SIZE"
	MaxConnIdleTime = "DB_MAX_CONN_IDLE_TIME"
	MaxConnLifeTime = "DB_MAX_CONN_LIFE_TIME"
)

type ConnectionProperties struct {
	Host            string
	Port            string
	Database        string
	MaxPoolSize     int
	MinPoolSize     int
	MaxConnIdleTime int
	MaxConnLifeTime int
}
