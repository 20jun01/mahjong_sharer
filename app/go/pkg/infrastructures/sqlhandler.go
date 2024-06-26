package infrastructures

import (
	"fmt"
	"log"
	"os"
	"time"

	sqldriver "github.com/go-sql-driver/mysql"
	"github.com/mahjong_sharer/pkg/infrastructures/migration"
	"github.com/mahjong_sharer/pkg/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	MAX_IDLE_CONNECTIONS  = 2
	MAX_OPEN_CONNECTIONS  = 16
	THRESHOLD_MILLISECOND = 200
)

func NewGormDB() (*gorm.DB, error) {
	engine, err := gorm.Open(
		mysql.New(mysql.Config{DSNConfig: newDsnConfig()}),
		newGormConfig(),
	)
	if err != nil {
		return nil, err
	}

	db, err := engine.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
	db.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
	if err := initDB(engine); err != nil {
		return nil, err
	}

	return engine, nil
}

func initDB(db *gorm.DB) error {
	_, err := migration.Migrate(db, migration.AllTables())
	if err != nil {
		return err
	}
	return nil
}

func newGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             THRESHOLD_MILLISECOND * time.Millisecond,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		NowFunc: func() time.Time {
			return time.Now().Truncate(time.Microsecond)
		},
	}
}

func newDsnConfig() *sqldriver.Config {
	if os.Getenv("NEOSHOWCASE") == "true" {
		return &sqldriver.Config{
			User:   util.ReadEnvs("NS_MARIADB_USER"),
			Passwd: util.ReadEnvs("NS_MARIADB_PASSWORD"),
			Net:    "tcp",
			Addr: fmt.Sprintf(
				"%s:%s",
				util.ReadEnvs("NS_MARIADB_HOSTNAME"),
				util.ReadEnvs("NS_MARIADB_PORT"),
			),
			DBName:               util.ReadEnvs("NS_MARIADB_DATABASE"),
			Collation:            "utf8mb4_general_ci",
			ParseTime:            true,
			AllowNativePasswords: true,
			Params: map[string]string{
				"charset": "utf8mb4",
			},
		}
	}
	return &sqldriver.Config{
		User:                 util.ReadEnvs("DB_USERNAME"),
		Passwd:               util.ReadEnvs("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", util.ReadEnvs("DB_HOSTNAME"), util.ReadEnvs("DB_PORT")),
		DBName:               util.ReadEnvs("DB_DATABASE"),
		Collation:            "utf8mb4_general_ci",
		ParseTime:            true,
		AllowNativePasswords: true,
		Params: map[string]string{
			"charset": "utf8mb4",
		},
	}
}
