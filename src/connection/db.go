package connection

import (
	"app/src/config"
	appLogger "app/src/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var db *gorm.DB
var dsn string

// Connect to the database.
func Connect(config *config.Config) {
	// set db log level
	gormConfig := &gorm.Config{}
	if config.App.LogLevel == "debug" {
		newGormLogger := gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			gormLogger.Config{
				SlowThreshold:             time.Second,     // Slow SQL threshold
				LogLevel:                  gormLogger.Info, // Log level
				IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,            // Enable color
			},
		)
		gormConfig = &gorm.Config{Logger: newGormLogger}
	}

	// process the configuration file
	var dialector gorm.Dialector
	dbConfig := config.Database
	if dbConfig.Provider == "postgres" {
		dsn = "host=" + dbConfig.Host + " user=" + dbConfig.User + " password=" + dbConfig.Password + " dbname=" + dbConfig.Name + " port=" + dbConfig.Port + " sslmode=" + dbConfig.SslMode + " TimeZone=" + dbConfig.Timezone
		dialector = postgres.Open(dsn)
	} else if dbConfig.Provider == "mysql" {
		dsn = dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
		dialector = mysql.Open(dsn)
	} else {
		appLogger.Fatal("Database provider not supported.", nil)
	}

	// connect to the database
	_db, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		appLogger.Fatal("Error connecting to database.", err)
	}

	// set the maximum number of connections
	sqlDb, err := _db.DB()
	if err != nil {
		appLogger.Fatal("Error getting database connection.", err)
	}
	sqlDb.SetMaxIdleConns(dbConfig.MaxConnections)
	sqlDb.SetMaxOpenConns(dbConfig.MaxConnections)

	db = _db
	appLogger.Info("Connected to database.")
}

// migrate the tables.
func Migrate(db *gorm.DB, models []interface{}) {
	err := db.AutoMigrate(models...)
	if err != nil {
		appLogger.Fatal("Error migrating tables.", err)
	}

}

// GetDB returns the db instance.
func GetDB() *gorm.DB {
	if db == nil {
		appLogger.Fatal("Database connection not initialized.", nil)
	}
	return db
}

// DB ping.
func Ping() error {
	sqlDB, err := db.DB()
	if err != nil {
		appLogger.Fatal("Error getting database connection.", err)
	}
	err = sqlDB.Ping()
	return err
}
