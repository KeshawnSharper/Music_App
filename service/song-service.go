package service 
import "gitlab.com/pragmaticreviews/golang-gin-poc/entity"
type SongService interface {
	Save(entity.Song) entity.Song
	findAll() []entity.Song
}
type SongService struct {
	songs []entity.Song
}
func New() SongService {
	return &songService{}
}
func (service *songService) Save(entity.Song) entity.Song{
	service.songs = append(service.songs,song)
	return song
}
func (service *songService) findAll() []entity.Song{
	return service.songs
}