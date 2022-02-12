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
	"lienquanMess/security"
	"net/http"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	req := req.ReqSignIp{}
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
	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Sai mật khẩu",
			Data:       nil,
		})
	}
	// gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}
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

	hash := security.HashAndSalt([]byte(req.Password))
	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

	}

	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Phone:    req.Phone,
		Password: hash,
	}

	userInfo := model.AllUser{
		UserId:   userId.String(),
		FullName: req.FullName,
	}

	userProfile := model.UserProfile{
		UserId:   userId.String(),
		FullName: req.FullName,
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user, userInfo, userProfile)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	// gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Đăng ký thành công",
		Data:       user,
	})
}

func (u *UserHandler) Profile(c echo.Context) error {
	//tokenData := c.Get("user").(*jwt.Token)
	//claims := tokenData.Claims.(*model.JwtCustomClaims)
	userId := c.QueryParam("userId")
	user, err := u.UserRepo.SelectUserProfileById(c.Request().Context(), userId)
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
		Message:    "Xử lý thành công",
		Data:       user,
	})
}

//func (u UserHandler) UpdateProfile(c echo.Context) error {
//	req := req.ReqUpdateUser{}
//	if err := c.Bind(&req); err != nil {
//		return err
//	}
//
//	// validate thông tin gửi lên
//	err := c.Validate(req)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.Response{
//			StatusCode: http.StatusBadRequest,
//			Message:    err.Error(),
//		})
//	}
//
//	token := c.Get("user").(*jwt.Token)
//	claims := token.Claims.(*model.JwtCustomClaims)
//	user := model.User{
//		UserId:    claims.UserId,
//		FullName:  req.FullName,
//		Height:    req.Height,
//		Age:       req.Age,
//		Sex:       req.Sex,
//		Job:       req.Job,
//		Income:    req.Income,
//		Marriage:  req.Marriage,
//		Children:  req.Children,
//		Home:      req.Home,
//		Zodiac:    req.Zodiac,
//		Status:    req.Status,
//		Formality: req.Formality,
//		LinkFb:    req.LinkFb,
//		LinkIs:    req.LinkIs,
//		Zalo:      req.Zalo,
//		Address:   req.Address,
//		Target:    req.Target,
//		About:     req.About,
//	}
//
//	user, err = u.UserRepo.UpdateUser(c.Request().Context(), user)
//	if err != nil {
//		return c.JSON(http.StatusUnprocessableEntity, model.Response{
//			StatusCode: http.StatusUnprocessableEntity,
//			Message:    err.Error(),
//		})
//	}
//
//	return c.JSON(http.StatusCreated, model.Response{
//		StatusCode: http.StatusCreated,
//		Message:    "Lưu thành công",
//		Data:       user,
//	})
//}
//
func (u UserHandler) UpdateLineProfile(c echo.Context) error {
	req := req.ReqUpdateUserLine{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate thông tin gửi lên
	err := c.Validate(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	user := model.UserProfile{
		UserId:         claims.UserId,
		Avatar:         req.Avatar,
		FullName:       req.FullName,
		Age:            req.Age,
		Sex:            req.Sex,
		Height:         req.Height,
		Job:            req.Job,
		Income:         req.Income,
		Marriage:       req.Marriage,
		Children:       req.Children,
		Home:           req.Home,
		Zodiac:         req.Zodiac,
		Status:         req.Status,
		Formality:      req.Formality,
		LinkFb:         req.LinkFb,
		LinkIs:         req.LinkIs,
		ZlPhone:        req.ZlPhone,
		Address:        req.Address,
		Target:         req.Target,
		About:          req.About,
		CountFollower:  req.CountFollower,
		CountFollowing: req.CountFollowing,
		CountLike:      req.CountLike,
	}

	user, err = u.UserRepo.UpdateLineUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, model.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, model.Response{
		StatusCode: http.StatusCreated,
		Message:    "Lưu thành công",
		Data:       user,
	})
}

func (u *UserHandler) AllUser(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	start := c.QueryParam("start")
	user, err := u.UserRepo.SelectAllUserById(c.Request().Context(), start, claims.UserId)
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
