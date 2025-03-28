package main

import (
	"github.com/wz-rjuarez/criptoagg/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCryptos(c *gin.Context) {
	c.JSON(http.StatusOK, configs.Layout)
}

func main() {
	r := gin.Default()

	r.GET("/cryptos", getCryptos)

	r.Run(":8080")
}
