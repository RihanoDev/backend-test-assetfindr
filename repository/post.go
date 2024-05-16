package repository

import (
	"post-api/config"
	"post-api/model"
)

type PostRepositoryImpl struct{}

func (p *PostRepositoryImpl) GetAllPost() []model.Post {
	var posts []model.Post

	config.DB.Preload("Tags").Find(&posts)

	return posts
}

func (p *PostRepositoryImpl) GetPostByID(id string) model.Post {
	var posts model.Post

	config.DB.Preload("Tags").First(&posts, "id = ?", id)

	return posts
}

func (p *PostRepositoryImpl) CreatePost(input model.PostInput) model.Post {
	var tags []model.Tag

	for _, tagName := range input.Tags {
		tag := model.Tag{Label: tagName}
		config.DB.FirstOrCreate(&tag, model.Tag{Label: tagName})
		tags = append(tags, tag)
	}

	var newPost model.Post = model.Post{
		Title:   input.Title,
		Content: input.Content,
		Tags:    tags,
	}

	var addedPost model.Post = model.Post{}

	result := config.DB.Create(&newPost)

	result.Last(&addedPost)

	return addedPost
}

func (p *PostRepositoryImpl) UpdatePost(id string, input model.PostInput) model.Post {
	var tags []model.Tag

	for _, tagName := range input.Tags {
		tag := model.Tag{Label: tagName}
		config.DB.FirstOrCreate(&tag, model.Tag{Label: tagName})
		tags = append(tags, tag)
	}

	var post model.Post = p.GetPostByID(id)

	post.Title = input.Title
	post.Content = input.Content
	post.Tags = tags

	config.DB.Save(&post)

	return post
}

func (p *PostRepositoryImpl) DeletePost(id string) bool {
	var post model.Post = p.GetPostByID(id)

	result := config.DB.Delete(&post)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
