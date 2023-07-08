package authn

import "time"

type NonceSpec struct {
	Nonce         string `json:"nonce"`
	AppIdentifier string `json:"appIdentifier"`
}

func (spec NonceSpec) ToNonce() *Nonce {
	return &Nonce{
		Nonce:   spec.Nonce,
		ExpDate: time.Now().Add(time.Second * 30),
	}
}
