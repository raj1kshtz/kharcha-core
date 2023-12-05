package conf

import (
	"github.com/raj1kshtz/kharcha-core/db"
	"os"
	"strconv"
)

func LoadConfiguration(opts ...string) {
	loadFromEnv()
}

func loadFromEnv() {
	props := db.ConnectionProperties{}
	props.Host = os.Getenv(db.Host)
	props.Port = os.Getenv(db.Port)
	props.Database = os.Getenv(db.Database)
	props.MaxPoolSize, _ = strconv.Atoi(os.Getenv(db.MaxPoolSize))
	props.MinPoolSize, _ = strconv.Atoi(os.Getenv(db.MinPoolSize))
	props.MaxConnIdleTime, _ = strconv.Atoi(os.Getenv(db.MaxConnIdleTime))
	props.MaxConnLifeTime, _ = strconv.Atoi(os.Getenv(db.MaxConnLifeTime))

	db.ConnProps = &props
}
