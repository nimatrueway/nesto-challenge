package server

import (
	"github.com/lmittmann/tint"
	"github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"log/slog"
	"net/http"
	"os"
	"readcommend/internal/api"
	"readcommend/internal/repository"
)

func Run() {
	// load yaml configuration
	config, err := LoadConfig()
	if err != nil {
		log.Panicf("Unable to load configuration: %#v", err)
	}

	// configure logger
	var logLevel slog.Level
	err = logLevel.UnmarshalText([]byte(config.Log.Level))
	if err != nil {
		log.Panicf("Unable to parse log level: %#v", err)
	}
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{Level: logLevel, AddSource: true}))
	slog.SetDefault(logger)

	// configure gorm
	gormLoggerOpts := []slogGorm.Option{slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelDebug)}
	if logLevel <= slog.LevelDebug {
		gormLoggerOpts = append(gormLoggerOpts, slogGorm.WithTraceAll())
	}
	gormLogger := slogGorm.New(gormLoggerOpts...)
	gormConf := gorm.Config{
		Logger:         gormLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}
	db, err := gorm.Open(postgres.Open(config.Database.Dns), &gormConf)
	if err != nil {
		log.Panicf("failed to connect to database: %#v", err)
	}

	// configure database connection pool
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if err != nil {
		log.Panicf("failed to get sql.DB: %#v", err)
	}
	sqlDB.SetMaxOpenConns(config.Database.MaxConns)
	sqlDB.SetConnMaxIdleTime(config.Database.MaxConnIdleTime)

	// instantiate router
	repo := repository.NewPgRepository(db)
	server := api.NewServer(repo)
	router := api.NewRouter(server)

	// start server
	slog.Info("Starting server", slog.String("address", config.Server.Host))
	if err := http.ListenAndServe(config.Server.Host, router); err != nil {
		log.Fatalf("failed to start server: %#v", err)
	}
}
