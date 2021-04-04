package main
import (
	"io"
	"os"
	
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/middlewares
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)
var(
	songService    service.VideoService       = service.New()
	songController controller.VideoController = controller.New(songService)

)
func setUpLogOutput(){
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main(){
	setUpLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(),middlewares.Logger(),
	middlewares.BasicAuth(),
)

	server.GET("/songs",func(ctx *gin.Context){
		ctx.JSON(200,songController.FindAll())
	})
	server.POST("/songs",func(ctx *gin.Context){
		ctx.JSON(200,songController.Save(ctx))
	})
server.Run()
}
