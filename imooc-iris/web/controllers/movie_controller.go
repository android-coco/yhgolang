package controllers

import (
	"github.com/kataras/iris/mvc"
	"yhgolang/imooc-iris/repositories"
	"yhgolang/imooc-iris/services"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManger(movieRepository)

	moiveResult := movieService.ShowMovieName()
	return mvc.View{
		Name:   "movie/index.html",
		Data:   moiveResult,
	}
}
