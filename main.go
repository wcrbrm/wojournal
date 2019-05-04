package main

import (
	"fmt"
	"os"

	"github.com/Depado/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
)

var app = cli.App("wojournal", "Write/Only Journal")

func main() {
	app.Before = prepareApp
	app.Action = runApp
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("[ERR]", err)
	}
}

func prepareApp() {
	log.SetLevel(log.Level(toNatural(*logLevel, 4)))
	if log.GetLevel() > log.InfoLevel {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func runApp() {
	dbClient := NewDatabaseClient()
	_ = dbClient

	log.WithFields(log.Fields{
		"address": "http://" + *webListenAddr,
	}).Println("Starting HTTP Server")

	r := gin.Default()

	// adding PROMETHEUS metrics client endpoint
	if *enableMetrics {
		p := ginprom.New(
			ginprom.Engine(r),
			ginprom.Subsystem("gin"),
			ginprom.Path("/metrics"),
		)
		r.Use(p.Instrument())
	}
	// adding CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// adding new records
	r.POST("/", func(c *gin.Context) {
		body, errBody := c.GetRawData()
		if errBody != nil {
			log.WithField("error", errBody.Error()).Error("c.GetRawData")
			c.Data(400, "text/plain; charset=utf-8", []byte("Invalid Body\n"))
			return
		}
		_ = body
		// TODO: add new record
		c.Data(200, "text/plain; charset=utf-8", []byte("OK\n"))
	})

	r.GET("/:tag", func(c *gin.Context) {
		c.Data(200, "text/plain; charset=utf-8", []byte("TAG\n"))
	})
	r.GET("/day/:day", func(c *gin.Context) {
		c.Data(200, "text/plain; charset=utf-8", []byte("STATS\n"))
	})
	r.GET("/stats", func(c *gin.Context) {
		c.Data(200, "text/plain; charset=utf-8", []byte("STATS\n"))
	})

	fmt.Println("Server started. Almost")
	r.Run(*webListenAddr)
}
