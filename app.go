package resk

import (
	_ "github.com/gzl-tommy/account/apis/web"
	_ "github.com/gzl-tommy/account/core/accounts"
	"github.com/gzl-tommy/infra"
	"github.com/gzl-tommy/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
