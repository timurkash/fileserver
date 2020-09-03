package handlers

import (
	"strconv"
	"testing"

	"gitlab.mcsolutions.ru/lib/common/config"
)

const (
	EXPECTED = 123
)

var (
	LOG_LEVEL       = config.GetEnv("LOG_LEVEL", "info")
	SENTRY_DSN      = config.GetEnv("SENTRY_DSN", "")
	EXPECTED_STRING = strconv.Itoa(EXPECTED)
)

func TestGetExampleRoute(t *testing.T) {
	//req, err := http.NewRequest(http.MethodGet, configmap.IMAGES_URL + "/crud?value=" + EXPECTED_STRING, nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//rr := httptest.NewRecorder()
	//routes := Routes{Logger: logger.NewLogger(LOG_LEVEL, SENTRY_DSN)}
	//handler := http.HandlerFunc(hh.GetObjHandlerFunc(routes.getExampleFunc, routes.Logger))
	//req.Header.Set(consts.CONTENT_TYPE, consts.CONTENT_TYPE_JSON)
	//handler.ServeHTTP(rr, req)
	//if rr.Code != http.StatusOK {
	//	t.Errorf("Expected %d but got %d", http.StatusOK, rr.Code)
	//}
	//examplesOut := &models.ExamplesOut{}
	//if err := json.Decode(rr.Body, examplesOut); err != nil {
	//	t.Fatal(err)
	//}
}

func TestPostExampleRoute(t *testing.T) {
	//example := &models.Example{Example: EXPECTED}
	//bytes_, err := json.EncodeBytes(example)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//req, err := http.NewRequest(http.MethodPost, configmap.IMAGES_URL + "/crud", bytes.NewBuffer(bytes_))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//rr := httptest.NewRecorder()
	//routes := Routes{Logger: logger.NewLogger(LOG_LEVEL, SENTRY_DSN)}
	//handler := http.HandlerFunc(hh.GetObjHandlerFunc(routes.postExampleFunc, routes.Logger))
	//req.Header.Set(consts.CONTENT_TYPE, consts.CONTENT_TYPE_JSON)
	//handler.ServeHTTP(rr, req)
	//if rr.Code != http.StatusOK {
	//	t.Errorf("Expected %d but got %d", http.StatusOK, rr.Code)
	//}
	//idOut := &status.IdOut{}
	//if err := json.Decode(rr.Body, idOut); err != nil {
	//	t.Fatal(err)
	//}
}

func TestPutExampleRoute(t *testing.T) {
	//example := &models.Example{Example: EXPECTED}
	//bytes_, err := json.EncodeBytes(example)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//req, err := http.NewRequest(http.MethodGet, configmap.IMAGES_URL + "/crud", bytes.NewBuffer(bytes_))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//req = mux.SetURLVars(req, map[string]string{"id": EXPECTED_STRING})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//rr := httptest.NewRecorder()
	//routes := Routes{Logger: logger.NewLogger(LOG_LEVEL, SENTRY_DSN)}
	//handler := http.HandlerFunc(hh.GetHandlerFunc(routes.putExampleFunc, routes.Logger))
	//req.Header.Set(consts.CONTENT_TYPE, consts.CONTENT_TYPE_JSON)
	//handler.ServeHTTP(rr, req)
	//if rr.Code != http.StatusOK {
	//	t.Errorf("Expected %d but got %d", http.StatusOK, rr.Code)
	//}
	//statusOut := &status.StatusOut{}
	//if err := json.Decode(rr.Body, statusOut); err != nil {
	//	t.Fatal(err)
	//}
	//if statusOut.Status == nil {
	//	t.Fatal(fmt.Errorf("statusOut.Status is nil"))
	//}
	//if !statusOut.Status.Ok {
	//	t.Fatal(fmt.Errorf("statusOut.Status.Ok is false"))
	//}
}

func TestDeleteExampleRoute(t *testing.T) {
	//req, err := http.NewRequest(http.MethodGet, configmap.IMAGES_URL + "/crud", nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//req = mux.SetURLVars(req, map[string]string{"id": EXPECTED_STRING})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//rr := httptest.NewRecorder()
	//routes := Routes{Logger: logger.NewLogger(LOG_LEVEL, SENTRY_DSN)}
	//handler := http.HandlerFunc(hh.GetHandlerFunc(routes.deleteExampleFunc, routes.Logger))
	//req.Header.Set(consts.CONTENT_TYPE, consts.CONTENT_TYPE_JSON)
	//handler.ServeHTTP(rr, req)
	//if rr.Code != http.StatusOK {
	//	t.Errorf("Expected %d but got %d", http.StatusOK, rr.Code)
	//}
	//statusOut := &status.StatusOut{}
	//if err := json.Decode(rr.Body, statusOut); err != nil {
	//	t.Fatal(err)
	//}
	//if statusOut.Status == nil {
	//	t.Fatal(fmt.Errorf("statusOut.Status is nil"))
	//}
	//if !statusOut.Status.Ok {
	//	t.Fatal(fmt.Errorf("statusOut.Status.Ok is false"))
	//}
}
