package services

import "moment/pkg/db"

type SessionCreateRequest struct {
	UserID PK `json:"userID"`

	IP     string `json:"ip"`
	OS     string `json:"os"`
	Mobile bool   `json:"mobile"`
}

func (r *SessionCreateRequest) toFullSession() *db.SessionFull {
	return &db.SessionFull{
		Session: db.Session{
			UserID: r.UserID,

			IP:     r.IP,
			OS:     r.OS,
			Mobile: r.Mobile,
		},
	}
}
