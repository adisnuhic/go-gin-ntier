package initialize

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

// returns instance of GIN framework
func Gin() *gin.Engine {
	g := gin.Default()

	// use request id
	g.Use(requestid.New())

	// init logrus
	lgr := initLogger()
	g.Use(ginlogrus.Logger(lgr), gin.Recovery())

	// set CORS default allow *
	g.Use(cors.Default())

	return g
}

// InitLogger configuration for logger
func initLogger() *logrus.Logger {
	lgr := logrus.New()
	lgr.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	return lgr
}
