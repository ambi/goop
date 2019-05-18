package config

import (
	"github.com/ambi/goop/domain/model"
)

// SingleOP is the single OP. TODO: initialize.
var SingleOP *model.OP

func init() {
	jwks := model.NewJSONWebKeySetFromPEMs([]string{Config.OIDC.PrivateKey})
	SingleOP = model.NewSimpleOP(Config.Server.URL, jwks)
}
