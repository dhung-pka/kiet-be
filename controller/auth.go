package controller

import (
	"encoding/json"
	"net/http"
	"service/config"
	"service/model"
	"service/service"
	"time"

	"github.com/go-chi/render"
)

type authController struct {
	authService service.AuthService
}

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request)
}

// @Summary      Login
// @Description  Login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param 			 req body model.LoginRequest true "request"
// @Success      200  {object}  model.Response
// @Router       /auth/login [post]
func (a *authController) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		badRequest(w, r, err)
		return
	}

	user, err := a.authService.Login(loginRequest)

	if err != nil {
		handlingError(w, r, err)
		return
	}

	_, tokenString, _ := config.GetJWT().Encode(map[string]interface{}{
		"profile_id": user.ProfileId,
		"user_name":  user.Username,
		"role":       user.Role,
	})

	accessCookie := http.Cookie{
		Name:    "access_token",
		Path:    "/",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 1 * 24),
	}

	refreshCookie := http.Cookie{
		Name:    "refresh_token",
		Path:    "/", // <--- add this line
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 3 * 24),
	}

	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &refreshCookie)

	res := model.Response{
		Data: map[string]interface{}{
			"access_token":  tokenString,
			"refresh_token": tokenString,
		},
		Page:    nil,
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)
}

func NewAuthController() AuthController {
	authService := service.NewAuthService()
	return &authController{
		authService: authService,
	}
}
