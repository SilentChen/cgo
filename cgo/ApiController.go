package cgo

import (
	"net/http"
	"cgo/entity"
	"cgo/constant"
)

type ApiController struct {
	Controller
}

func (p *ApiController) GetUserId(r *http.Request) uint {
	user := p.GetUser(r)
	if user == nil {
		return 0
	}
	return user.ID
}

func (p *ApiController) GetUser(r *http.Request) *entity.User {
	session := Get(r)
	key_user := session.Get(constant.KEY_USER)
	if user,ok := key_user.(*entity.User);ok{
		return user
	}
	return nil
}
