package controller

import (
	"net/http"
	"post-api/model"
	"post-api/service"

	"github.com/gin-gonic/gin"
)

var TagService service.TagService = service.NewTagService()

func GetAllTag(c *gin.Context) {
	var tags []model.Tag = TagService.GetAllTag()

	c.JSON(http.StatusOK, tags)
}

func GetTagByID(c *gin.Context) {
	var id = c.Param("id")

	tag := TagService.GetTagByID(id)

	if tag.ID == 0 {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "tag not Found",
		})
	}

	c.JSON(http.StatusOK, tag)
}

func CreateTag(c *gin.Context) {
	var input *model.TagInput = new(model.TagInput)

	if err := c.Bind(input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := input.Validate()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	tag := TagService.CreateTag(*input)

	c.JSON(http.StatusCreated, tag)
}

func UpdateTag(c *gin.Context) {
	var input *model.TagInput = new(model.TagInput)

	if err := c.Bind(input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := input.Validate()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	var tagId string = c.Param("id")

	tag := TagService.UpdateTag(tagId, *input)

	c.JSON(http.StatusOK, tag)
}

func DeleteTag(c *gin.Context) {
	var tagId string = c.Param("id")

	isSuccess := TagService.DeleteTag(tagId)

	if !isSuccess {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete a data",
		})
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Data deleted",
	})
}
