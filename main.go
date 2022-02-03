package main

import (
	"github.com/gin-gonic/gin"
)

/* main es una función que inicializa el servidor predeterminado
 */
func main() {
	servidor := gin.Default()
	servidor.SetTrustedProxies([]string{"192.168.0.5"})

	// gin.Context incluye la solicitud HTTP
	servidor.GET("/test", func(contexto *gin.Context) {
		contexto.JSON(200, gin.H{
			"messaje": "OK!!",
		})
	})

	servidor.Run(":9098")
}
