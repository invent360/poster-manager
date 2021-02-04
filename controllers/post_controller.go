package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/post-manager/models"
	"net/http"
)

var posts []models.Post

func GetMessage(c echo.Context) error {
	return c.String(http.StatusOK, "Bingo!!!")
}

func FetchPosts(c echo.Context) error {
	return c.String(http.StatusOK, "Success")
}

func FetchPost(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
