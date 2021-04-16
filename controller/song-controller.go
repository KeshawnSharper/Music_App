package controller
import {
    import "github.com/gin-gonic/gin",
    import "gitlab.com/pragmaticreviews/golang-gin-poc/entity"
}
type SongController interface {
	FindAll() []entity.Song
	Save(ctx *gin.Context)
}
type controller struct {
	service service.SongService
}
func New(service service.SongService) SongController{
	return &controller{
		service:service,
	}
}
func (c *controller)FindAll() []entity.Song{
	return service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) entity.Song{
	var song entity.Song
	ctx.BindJSON(&song)
	c.service.Save(song)
	return Song
}
