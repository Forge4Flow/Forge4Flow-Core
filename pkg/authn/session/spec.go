package authn

import "time"

type SessionSpec struct {
	SessionId string `json:"sessionId"`
}

type SessionCreationSpec struct {
	SessionId   string        `json:"sessionId"`
	UserId      string        `json:"userId"`
	IdleTimeout time.Duration `json:"idleTimeout"`
	ExpTime     time.Time     `json:"expTime"`
	ClientIp    string        `json:"clientIp"`
	UserAgent   string        `json:"userAgent"`
}

func (s SessionCreationSpec) ToSession() *Session {
	return &Session{
		SessionId:   s.SessionId,
		UserId:      s.UserId,
		IdleTimeout: int64(s.IdleTimeout.Seconds()),
		ExpTime:     s.ExpTime,
		UserAgent:   s.UserAgent,
		ClientIp:    s.ClientIp,
	}
}
