package service

import (
	"post-api/model"
	"post-api/repository"
)

type TagService struct {
	Repository repository.TagRepository
}

func NewTagService() TagService {
	return TagService{
		Repository: &repository.TagRepositoryImpl{},
	}
}

func (n *TagService) GetAllTag() []model.Tag {
	return n.Repository.GetAllTag()
}

func (n *TagService) GetTagByID(id string) model.Tag {
	return n.Repository.GetTagByID(id)
}

func (n *TagService) CreateTag(input model.TagInput) model.Tag {
	return n.Repository.CreateTag(input)
}

func (n *TagService) UpdateTag(id string, input model.TagInput) model.Tag {
	return n.Repository.UpdateTag(id, input)
}

func (n *TagService) DeleteTag(id string) bool {
	return n.Repository.DeleteTag(id)
}
