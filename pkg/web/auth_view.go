package web

import (
	"moment/pkg"
	phoneapi "moment/pkg/phoneAPI"
	"moment/pkg/services"
	"moment/pkg/utilty"
	"net/http"

	"gitlab.com/knopkalab/go/logger"
)

func AuthView(user_s services.User, ucl phoneapi.Ucaller, log logger.Logger) View {
	return func(r *Request) error {
		var resp services.UserUpdateRequest
		if err := r.ParseJSON(&resp); err != nil {
			return r.Response.StatusBadRequest()
		}
		defer r.Body.Close()
		if resp.Phone == "" {
			return r.Response.StatusBadRequest()
		}
		user, err := user_s.GetUserByPhone(resp.Phone, r.Context)
		if err == nil && user != nil {
			code := ucl.GenerateCode()
			go ucl.Send(code, resp.Phone, log)

			resp.Code = utilty.NewPassword(code)

			_, err = user_s.UpdateUser(r.Context, &resp, user)
			if err != nil {
				return r.Response.StatusBadRequest()
			}
			return r.Response.OKJSON(map[string]string{
				"code": code,
			})
		}

		code := ucl.GenerateCode()
		go ucl.Send(code, resp.Phone, log)

		resp.Code = utilty.NewPassword(code)
		user_s.CreateUser(&resp)
		return r.Response.OKJSON(map[string]string{
			"code": code,
		})

	}
}

func LoginView(user_s services.User, sessions services.Sessions, log logger.Logger, conf *pkg.Config) View {
	return func(r *Request) error {
		var resp services.UserUpdateRequest

		if err := r.ParseJSON(&resp); err != nil {
			return r.Response.StatusBadRequest()
		}
		user, err := user_s.GetUserByPhone(resp.Phone, r.Context)
		if err != nil || user == nil {
			return r.Response.StatusBadRequest()
		}

		if !utilty.Match(user.Code, resp.Code) {
			return r.Response.StatusBadRequest()
		}
		ua := r.UserAgent()
		resp_session := services.SessionCreateRequest{
			UserID: user.ID,
			IP:     r.RealIP(),
			OS:     ua.OS(),
			Mobile: ua.Mobile(),
		}
		session, err := sessions.Create(&resp_session)
		if err != nil {
			return r.Response.StatusBadRequest()
		}
		r.Response.SetCookie(&http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Name:     conf.Sessions.CookieName,
			MaxAge:   conf.Sessions.CookieMaxAge.Seconds(),
			Value:    session.Token,
		})
		return r.Response.StatusOK()

	}

}

func Tesst(r *Request) error {
	return r.Response.OKJSON(map[string]string{
		"authTest": "ok",
	})
}
