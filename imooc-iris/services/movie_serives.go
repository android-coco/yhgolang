package services

import (
	"fmt"
	"yhgolang/imooc-iris/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManger struct {
	repo repositories.MovieRepository
}

func (m *MovieServiceManger) ShowMovieName() string {
	fmt.Println("我们获取到的视频名称为：" + m.repo.GetMovieName())
	return "我们获取到的视频名称为：" + m.repo.GetMovieName()
}

func NewMovieServiceManger(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManger{repo: repo}
}
