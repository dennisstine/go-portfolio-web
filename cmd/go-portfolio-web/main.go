package main

import (
	"github.com/dennisstine/go-portfolio-web/pkg/messages"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var router *gin.Engine

func setupRouter() *gin.Engine {

	router = gin.Default()
	//router.Static("/static", "web/static")
	router.Use(static.ServeRoot("/", "web/static"))
	router.LoadHTMLGlob("web/templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"facebook":   viper.GetString("links.facebook"),
			"twitter":    viper.GetString("links.twitter"),
			"linkedin":   viper.GetString("links.linkedIn"),
			"github":     viper.GetString("links.gitHub"),
			"hackerrank": viper.GetString("links.hackerRank"),
		})
	})

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	router.POST("/messages", messages.HandleMessage)

	return router
}

func init() {

	initLogger()
	initViper()
}

func initLogger() {

	format := new(log.TextFormatter)
	format.ForceColors = true
	format.FullTimestamp = true
	format.TimestampFormat = "2006-01-02 15:04:05-700"

	log.SetFormatter(format)
	log.SetLevel(log.Level(log.DebugLevel))

	log.Info("Initialization complete")
}

func initViper() {

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}

	log.Infof("Viper initialized")
	log.Debugf("%d keys found in %s", len(viper.AllKeys()), viper.ConfigFileUsed())

	for _, key := range viper.AllKeys() {
		log.Debug(key)
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := setupRouter()

	port := os.Getenv("PORT")

	if port == "" {
		log.Warnf("$PORT not set. Defaulting to :8080")
		port = ":8080"
	} else {
		port = ":" + port
	}

	err := router.Run(port)

	if err != nil {
		log.Fatal("Could not start")
	}
}
