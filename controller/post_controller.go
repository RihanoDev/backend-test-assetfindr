package controller

import (
	"net/http"
	"post-api/model"
	"post-api/service"

	"github.com/gin-gonic/gin"
)

var PostService service.PostService = service.NewPostService()

func GetAllPost(c *gin.Context) {
	var posts []model.Post = PostService.GetAllPost()

	c.JSON(http.StatusOK, posts)
}

func GetPostByID(c *gin.Context) {
	var id = c.Param("id")

	post := PostService.GetPostByID(id)

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, map[string]string{
			"message": "Post not Found",
		})
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	var input *model.PostInput = new(model.PostInput)

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

	post := PostService.CreatePost(*input)

	c.JSON(http.StatusCreated, post)
}

func UpdatePost(c *gin.Context) {
	var input *model.PostInput = new(model.PostInput)

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

	var postId string = c.Param("id")

	post := PostService.UpdatePost(postId, *input)

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	var postId string = c.Param("id")

	isSuccess := PostService.DeletePost(postId)

	if !isSuccess {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete a data",
		})
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Data deleted",
	})
}
