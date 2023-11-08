package authn

import (
	"time"
)

type Model interface {
	GetID() int64
	GetSessionId() string
	GetUserId() string
	GetLastActivity() time.Time
	GetIdleTimeout() time.Duration
	GetExpTime() time.Time
	GetUserAgent() string
	GetClientIp() string
	IsExpired() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time
	ToSessionSpec() *SessionSpec
}

type Session struct {
	ID           int64      `mysql:"id" postgres:"id" sqlite:"id"`
	SessionId    string     `mysql:"sessionId" postgres:"session_id" sqlite:"sessionId"`
	UserId       string     `mysql:"userId" postgres:"user_id" sqlite:"userId"`
	LastActivity time.Time  `mysql:"lastActivity" postgres:"las_activity" sqlite:"lastActivity"`
	IdleTimeout  int64      `mysql:"idleTimeout" postgres:"idle_timeout" sqlite:"idleTimeout"`
	ExpTime      time.Time  `mysql:"expTime" postgres:"exp_time" sqlite:"expTime"`
	UserAgent    string     `mysql:"userAgent" postgres:"user_agent" sqlite:"userAgent"`
	ClientIp     string     `mysql:"clientIp" postgres:"client_ip" sqlite:"clientIp"`
	CreatedAt    time.Time  `mysql:"createdAt" postgres:"created_at" sqlite:"createdAt"`
	UpdatedAt    time.Time  `mysql:"updatedAt" postgres:"updated_at" sqlite:"updatedAt"`
	DeletedAt    *time.Time `mysql:"deletedAt" postgres:"deleted_at" sqlite:"deletedAt"`
}

func (session Session) GetID() int64 {
	return session.ID
}

func (session Session) GetSessionId() string {
	return session.SessionId
}

func (session Session) GetUserId() string {
	return session.UserId
}

func (session Session) GetLastActivity() time.Time {
	return session.LastActivity
}

func (session Session) GetIdleTimeout() time.Duration {
	return time.Duration(session.IdleTimeout)
}

func (session Session) GetExpTime() time.Time {
	return session.ExpTime
}

func (session Session) GetUserAgent() string {
	return session.UserAgent
}

func (session Session) GetClientIp() string {
	return session.ClientIp
}

func (session Session) IsExpired() bool {
	now := time.Now().UTC()
	idleExpired := now.After(session.LastActivity.Add(session.GetIdleTimeout()))
	absoluteExpired := now.After(session.ExpTime)

	// The session is expired if it's either idle expired or absolute expired
	return idleExpired || absoluteExpired
}

func (session Session) GetCreatedAt() time.Time {
	return session.CreatedAt
}

func (session Session) GetUpdatedAt() time.Time {
	return session.UpdatedAt
}

func (session Session) GetDeletedAt() *time.Time {
	return session.DeletedAt
}

func (session Session) ToSessionSpec() *SessionSpec {
	return &SessionSpec{
		SessionId: session.SessionId,
	}
}
