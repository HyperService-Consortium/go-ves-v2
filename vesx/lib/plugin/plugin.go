package plugin

import (
	"context"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/go-ves/vesx/config"
	"github.com/Myriad-Dreamin/go-ves/vesx/model"
	"github.com/Myriad-Dreamin/go-ves/vesx/service"
	"github.com/Myriad-Dreamin/go-ves/vesx/types"
)

type Logger = types.Logger
type ConfigLoader = types.ConfigLoader
type ServiceProvider = service.Provider
type DatabaseProvider = model.Provider
type ServerConfig = config.ServerConfig
type Module = module.Module

type Plugin interface {
	Configuration(logger Logger, loader ConfigLoader, cfg *ServerConfig) (plg Plugin)
	Inject(services *ServiceProvider, dbs *DatabaseProvider, module Module) (plg Plugin)
	Work(ctx context.Context)
}
