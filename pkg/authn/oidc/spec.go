package authn

type OauthProvider struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	ImgURL       string   `json:"imgUrl"`
	UserIdClaim  string   `json:"userIdClaim"`
	Scopes       []string `json:"scopes"`
	ClientID     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	CodeURL      string   `json:"codeUrl"`
	TokenURL     string   `json:"tokenUrl"`
}
