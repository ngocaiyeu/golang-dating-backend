package router

import (
	"github.com/labstack/echo/v4"
	"lienquanMess/handler"
	"lienquanMess/middleware"
)

type API struct {
	Echo            *echo.Echo
	UserHandler     handler.UserHandler
	NewsFeedHandler handler.NewsFeedHandler
}

func (api *API) SetupRouter() {
	// User
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
	// Profile
	user := api.Echo.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.Profile)
	//user.PUT("/profile/update", api.UserHandler.UpdateProfile)
	user.PUT("/profile/update/line", api.UserHandler.UpdateLineProfile)
	user.GET("/profile/all", api.UserHandler.AllUser)

	// NewsFeed
	newsFeed := api.Echo.Group("/newsfeed", middleware.JWTMiddleware())
	newsFeed.GET("/all", api.NewsFeedHandler.HandlerSelectAllPosts)
	newsFeed.POST("/addpost", api.NewsFeedHandler.HandlerAddPost)
	newsFeed.GET("/userpost", api.NewsFeedHandler.SelectUser)
}
