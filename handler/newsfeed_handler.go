package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"lienquanMess/banana"
	"lienquanMess/model"
	"lienquanMess/model/req"
	"lienquanMess/repository"
	"net/http"
)

type NewsFeedHandler struct {
	NewsFeedRepo repository.NewsFeedRepo
}

func (n *NewsFeedHandler) HandlerSelectAllPosts(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	start := c.QueryParam("start")

	posts, err := n.NewsFeedRepo.SelectAllNewsFeed(c.Request().Context(), start, claims.UserId)
	if err != nil {
		if err == banana.UserNotFound {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Ok",
		Data:       posts,
	})
}

func (n *NewsFeedHandler) HandlerAddPost(c echo.Context) error {
	req := req.ReqAddPost{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	postId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

	}
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	post := model.Posts{
		Id:             postId.String(),
		UserId:         claims.UserId,
		AccessModifier: req.AccessModifier,
		Content:        req.Content,
		ImageUrl:       req.ImageUrl,
	}

	likeId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

	}
	like := model.Likes{
		Id:     likeId.String(),
		PostID: postId.String(),
	}

	post, err = n.NewsFeedRepo.AddPosts(c.Request().Context(), post, like)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng bài thành công",
		Data:       post,
	})
}

func (n *NewsFeedHandler) SelectUser(c echo.Context) error {
	userId := c.QueryParam("userId")
	user, err := n.NewsFeedRepo.SelectUserById(c.Request().Context(), userId)
	if err != nil {
		if err == banana.UserNotFound {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})
		}

		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Ok",
		Data:       user,
	})
}
