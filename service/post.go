package service

import (
	"post-api/model"
	"post-api/repository"
)

type PostService struct {
	Repository repository.PostRepository
}

func NewPostService() PostService {
	return PostService{
		Repository: &repository.PostRepositoryImpl{},
	}
}

func (n *PostService) GetAllPost() []model.Post {
	return n.Repository.GetAllPost()
}

func (n *PostService) GetPostByID(id string) model.Post {
	return n.Repository.GetPostByID(id)
}

func (n *PostService) CreatePost(input model.PostInput) model.Post {
	return n.Repository.CreatePost(input)
}

func (n *PostService) UpdatePost(id string, input model.PostInput) model.Post {
	return n.Repository.UpdatePost(id, input)
}

func (n *PostService) DeletePost(id string) bool {
	return n.Repository.DeletePost(id)
}
