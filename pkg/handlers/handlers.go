package handlers

import (
	"gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/handlers/files"
	"gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/models"
	"gitlab.mcsolutions.ru/lib/common/config"
	"gitlab.mcsolutions.ru/lib/common/header"
	"gitlab.mcsolutions.ru/lib/common/hh"
	"gitlab.mcsolutions.ru/lib/common/logger"
	"net/http"
)

var (
	TOKEN = config.GetEnv("TOKEN", "f9403fc5f537b4ab332d")
)

type Routes struct {
	Logger *logger.Logger
	Files  *files.Files
}

func (ro *Routes) UploadRoute() hh.Route {
	return hh.Route{
		Pattern:     "/upload",
		Methods:     http.MethodPost,
		HandlerFunc: hh.GetCodeObjHandlerFunc(ro.uploadFunc, ro.Logger),
	}
}

func (ro *Routes) uploadFunc(r *http.Request) (int, interface{}, error) {
	token, err := header.GetHeaderRequired(r, "token")
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if token != TOKEN {
		return http.StatusUnauthorized, nil, err
	}
	srcFile, info, err := r.FormFile("file")
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer srcFile.Close()
	path, err := ro.Files.Upload(srcFile, info)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, &models.PathOut{Path: path}, nil
}

//func (ro *Routes) GetDownloadRoute() hh.Route {
//	return hh.Route{
//		Pattern:     "/download/{file}",
//		Methods:     http.MethodGet,
//		HandlerFunc: hh.GetCodeObjFileHandlerFunc(ro.getDownloadFunc, ro.Logger),
//	}
//}
//
//func (ro *Routes) getDownloadFunc(r *http.Request) (int, interface{}, string, error) {
//	token, err := header.GetHeaderRequired(r, "token")
//	if err != nil {
//		return http.StatusBadRequest, nil, "", err
//	}
//	if token != TOKEN {
//		return http.StatusUnauthorized, nil, "", err
//	}
//	file, err := header.GetVarRequired(r, "file")
//	if err != nil {
//		return http.StatusBadRequest, nil, "", err
//	}
//	return http.StatusOK, nil, file, nil
//}
