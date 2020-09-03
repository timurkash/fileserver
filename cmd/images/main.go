package main

import (
	"gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/handlers/files"
	"gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/routers"
	"gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/version"
	"gitlab.mcsolutions.ru/find-psy/common/consts"
	"gitlab.mcsolutions.ru/find-psy/configmap"
	"gitlab.mcsolutions.ru/lib/common/config"
	"gitlab.mcsolutions.ru/lib/common/logger"
	"gitlab.mcsolutions.ru/lib/common/server"
	"time"
)

var (
	LOG_LEVEL  = config.GetEnv("LOG_LEVEL", "info")
	SENTRY_DSN = config.GetEnv("SENTRY_DSN", "")
)

func main() {
	logger := logger.NewLogger(LOG_LEVEL, SENTRY_DSN)
	server := server.HttpServer{
		Started:  time.Now(),
		Subgroup: configmap.IMAGES_GITLAB_SUBGROUP,
		Name:     configmap.IMAGES,
		Port:     configmap.IMAGES_PORT,
		BasePath: configmap.IMAGES_BASE_PATH,
		Duration: consts.GRSH_SECONDS * time.Second,
		Routes:   routers.GetRoutes(logger),
		Docs:     "./docs",
		Logger:   logger,
		Version:  version.VERSION,
		Revision: version.REVISION,
		Envs:     &version.Envs,
		Static: &server.Static{
			Route:   configmap.IMAGES_BASE_PATH + "/download",
			FileDir: files.FILES_DIR,
		},
	}
	server.RunGrSh()
}
