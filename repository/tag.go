package repository

import (
	"post-api/config"
	"post-api/model"
)

type TagRepositoryImpl struct{}

func (p *TagRepositoryImpl) GetAllTag() []model.Tag {
	var tags []model.Tag

	config.DB.Preload("Posts").Find(&tags)

	return tags
}

func (p *TagRepositoryImpl) GetTagByID(id string) model.Tag {
	var tags model.Tag

	config.DB.Preload("Posts").First(&tags, "id = ?", id)

	return tags
}

func (p *TagRepositoryImpl) CreateTag(input model.TagInput) model.Tag {
	var newTag model.Tag = model.Tag{
		Label: input.Label,
	}

	var addedTag model.Tag = model.Tag{}

	result := config.DB.Create(&newTag)

	result.Last(&addedTag)

	return addedTag
}

func (p *TagRepositoryImpl) UpdateTag(id string, input model.TagInput) model.Tag {
	var tag model.Tag = p.GetTagByID(id)

	tag.Label = input.Label

	config.DB.Save(&tag)

	return tag
}

func (p *TagRepositoryImpl) DeleteTag(id string) bool {
	var tag model.Tag = p.GetTagByID(id)

	result := config.DB.Delete(&tag)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
