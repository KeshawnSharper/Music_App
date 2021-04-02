package main
import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)
var(
	songService    service.VideoService       = service.New()
	songController controller.VideoController = controller.New(songService)

)
func main(){
	server := gin.New()
	server.Use(gin.Recovery(),middlewares.Logger())

	server.GET("/songs",func(ctx *gin.Context){
		ctx.JSON(200,songController.FindAll())
	})
	server.POST("/songs",func(ctx *gin.Context){
		ctx.JSON(200,songController.Save(ctx))
	})
server.Run()
}