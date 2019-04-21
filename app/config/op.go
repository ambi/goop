package config

import (
	"github.com/ambi/goop/domain/model"
)

// SingleOP は唯一の OP。TODO: initialize.
var SingleOP *model.OP

func init() {
	jwks := model.NewJSONWebKeySetFromPEMs([]string{Config.OIDC.PrivateKey})
	SingleOP = model.NewSimpleOP(Config.Server.URL, jwks)
}
