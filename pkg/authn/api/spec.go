package authn

import "time"

type ApiSpec struct {
	DisplayName string `json:"displayName"`
	Key         string `json:"key"`
	ExpDate     string `json:"expDate"`
}

func (spec ApiSpec) ToNonce() *ApiKey {
	return &ApiKey{
		DisplayName: spec.DisplayName,
		Key:         spec.Key,
		ExpDate:     time.Now().Add(time.Hour * 24 * 90),
	}
}
