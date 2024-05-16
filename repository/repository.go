package repository

import "post-api/model"

type PostRepository interface {
	GetAllPost() []model.Post
	GetPostByID(id string) model.Post
	CreatePost(input model.PostInput) model.Post
	UpdatePost(id string, input model.PostInput) model.Post
	DeletePost(id string) bool
}

type TagRepository interface {
	GetAllTag() []model.Tag
	GetTagByID(id string) model.Tag
	CreateTag(input model.TagInput) model.Tag
	UpdateTag(id string, input model.TagInput) model.Tag
	DeleteTag(id string) bool
}
