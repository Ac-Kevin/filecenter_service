package user

import (
	"errors"
	"strings"

	"github.com/Ac-Kevin/filecenter_service/interfaces"
)

// Service User Handler Def.
type Service struct {
	src interfaces.UserHandler
}

// NewHandler Def .
func NewService(h *interfaces.UserHandler) *Service {
	return &Service{*h}
}

// Auth login
func (h *Service) Auth(username, password string) (*User, error) {
	u, e := h.Get(username)
	if e != nil {
		return nil, e
	}
	if u == nil || !strings.EqualFold(u.GetUserPassword(), password) {
		return nil, errors.New("user or password invalid")
	}
	return u, nil
}

func (h *Service) Get(username string) (*User, error) {
	u, e := h.src.Get(username)
	if e != nil {
		return nil, e
	}
	return &User{
		_UserName:         u.GetUserName(),
		_Password:         u.GetUserPassword(),
		_UploadRateByte:   u.GetFileUploadRateByte(),
		_DownloadRateByte: u.GetFileDownloadRateByte(),
	}, nil
}

func (h *Service) Save(user *User) error {
	return h.src.Save(user)
}
