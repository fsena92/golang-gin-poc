package service

import (
	"github.com/fsena92/golang-gin-poc/entity"
	"github.com/fsena92/golang-gin-poc/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	//videos []entity.Video
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	//service.videos = append(service.videos, video)
	service.videoRepository.Save(video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	//return service.videos
	return service.videoRepository.FindAll()
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}
