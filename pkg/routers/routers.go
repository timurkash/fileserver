package routers

import (
	_ "gitlab.mcsolutions.ru/find-psy/back/users/images/docs"
	"gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/handlers"
	"gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/handlers/files"
	"gitlab.mcsolutions.ru/lib/common/hh"
	"gitlab.mcsolutions.ru/lib/common/logger"
)

func GetRoutes(logger *logger.Logger) *[]hh.Route {
	routes := handlers.Routes{
		Logger: logger,
		Files:  &files.Files{},
	}
	return &[]hh.Route{
		routes.UploadRoute(), // images
	}
}
