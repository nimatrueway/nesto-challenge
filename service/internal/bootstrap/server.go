package bootstrap

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"readcommend/internal/api"
	"readcommend/internal/controller"
	"readcommend/internal/repository"
	"readcommend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/lmittmann/tint"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Run() {
	// load yaml configuration
	config, err := LoadConfig()
	if err != nil {
		log.Panicf("Unable to load configuration: %#v", err)
	}

	// configure logger
	var logLevel slog.Level
	if err := logLevel.UnmarshalText([]byte(config.Log.Level)); err != nil {
		log.Panicf("Unable to parse log level: %#v", err)
	}
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{Level: logLevel, AddSource: true}))
	slog.SetDefault(logger)

	// configure gorm
	gormLoggerOpts := []slogGorm.Option{
		slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelDebug),
		slogGorm.WithSlowThreshold(config.Database.SlowQueryThreshold),
	}
	if logLevel <= slog.LevelDebug {
		gormLoggerOpts = append(gormLoggerOpts, slogGorm.WithTraceAll())
	}
	gormLogger := slogGorm.New(gormLoggerOpts...)
	gormConf := gorm.Config{
		Logger:         gormLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}
	var db *gorm.DB
	if db, err = gorm.Open(postgres.Open(config.Database.URL), &gormConf); err != nil {
		log.Panicf("failed to connect to database: %#v", err)
	}

	// configure database connection pool
	sqlDB, err := db.DB()
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("failed to close sql.DB: %#v", err)
		}
	}()
	if err != nil {
		log.Panicf("failed to get sql.DB: %#v", err)
	}
	sqlDB.SetMaxOpenConns(config.Database.MaxConns)
	sqlDB.SetConnMaxIdleTime(config.Database.MaxConnIdleTime)

	// instantiate router
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	server := controller.NewController(bookService)
	router := api.NewRouter(server, []gin.HandlerFunc{api.CORSMiddleware(config.Server.CorsAllowedOrigins)})

	// start server
	httpServer := http.Server{
		Addr:              config.Server.Host,
		Handler:           router,
		ReadHeaderTimeout: config.Server.RequestReadHeaderTimeout,
	}
	slog.Info("Starting server", slog.String("address", config.Server.Host))

	if err := httpServer.ListenAndServe(); err != nil {
		log.Panicf("failed to start server: %#v", err)
	}
}
