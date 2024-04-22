package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer"
)

func (h *Post) create(c *gin.Context) {

	f := datatransfer.PostCreate{}
	err := c.ShouldBindJSON(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := database.Post{
		Title:   f.Title,
		Content: f.Content,
	}
	err = post.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
