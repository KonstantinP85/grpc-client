package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type result struct {
	Result string `json:"result"`
}

type response struct {
	Message string `json:"message"`
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, response{message})
}

func (h *Handler) getNews(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	news, err := h.services.News.GetNews(id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, news)
}

func (h *Handler) uploadNews(c *gin.Context) {
	count, err := h.services.News.UploadNews()
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, count)
}
