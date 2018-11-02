package http

import (
	"DiskCheck/config"
	"DiskCheck/diskCheck"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateHttpWeb() error {
	addr := config.GetServerConfig("host").(string)+":1212"

	router := CreateRouter()
	s := http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()

	return err
}

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/diskStatus", diskCheck.GetDiskStatus)

	return router
}
