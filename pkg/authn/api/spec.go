package authn

import "time"

type ApiSpec struct {
	DisplayName string    `json:"displayName"`
	Key         string    `json:"key"`
	ExpDate     time.Time `json:"expDate"`
}

func (spec ApiSpec) ToApiKey() *ApiKey {
	return &ApiKey{
		DisplayName: spec.DisplayName,
		ApiKey:      spec.Key,
		ExpDate:     spec.ExpDate,
	}
}
