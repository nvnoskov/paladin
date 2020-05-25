package main

import (
	"flag"
	_ "net/http/pprof"
	"runtime"

	"github.com/savsgio/atreugo/v11"
	"github.com/sirupsen/logrus"
	"kurs.kz/paladin/cache"
	"kurs.kz/paladin/controllers"
	"kurs.kz/paladin/db"
)

var databasePath string
var listenAddr string
var verboseLogging bool
var profile bool

func main() {
	flag.StringVar(&databasePath, "database", "./database",
		"Path to the database")
	flag.StringVar(&listenAddr, "listen", ":8000",
		"Listen address")
	flag.BoolVar(&verboseLogging, "verbose", true,
		"Enable verbose logging")
	flag.BoolVar(&profile, "profile", false,
		"Enable blocks and mutexes profiling")
	flag.Parse()

	if profile {
		runtime.SetBlockProfileRate(20)
		runtime.SetMutexProfileFraction(20)
	}

	if verboseLogging {
		logrus.SetLevel(logrus.DebugLevel)
	}

	err := db.OpenDatabase(databasePath)
	if err != nil {
		logrus.Fatalf("openDatabase: %v\n", err)
	}

	defer func() {
		err := db.DB.Close()
		if err != nil {
			logrus.Fatalf("DB.Close(): %v\n", err)
		}
	}()

	logrus.Infof("Starting badgerCleanupProc...")
	go db.BadgerCleanupProc()

	logrus.Infof("Database ready! Starting HTTP server at %d...",
		listenAddr)

	config := atreugo.Config{
		Addr: listenAddr,		
	}
	server := atreugo.New(config)

	punkts:=server.NewGroupPath("/punkts")
	punkts.POST("/", controllers.SyncPunkts)
	punkts.PUT("/{id}", controllers.UpdatePunkt)
	punkts.POST("/{id}", controllers.UpdatePunktData)	
	punkts.GET("/", controllers.GetPunkts)	
	punkts.GET("/{id}", controllers.GetPunkt)
	server.Path("GET", "/stats", func(ctx *atreugo.RequestCtx) error {
		stats := cache.PaladinCache.Stats()
		return ctx.JSONResponse(stats)
	})

	err = server.ListenAndServe()
	if err != nil {
		logrus.Fatalf("http.ListenAndServe: %v\n", err)
	}

	logrus.Info("HTTP server terminated\n")
}
