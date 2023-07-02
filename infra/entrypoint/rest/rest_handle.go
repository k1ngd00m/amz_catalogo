package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/k1ngd00m/amz_comun/excepciones"
)

type HandlerExcepcion struct {
	logger *zap.Logger
}

type ErrorResponse struct {
	Mensaje   string
	RequestId string
}

func NewHandlerExcepcion(logger *zap.Logger) HandlerExcepcion {
	handle := HandlerExcepcion{
		logger: logger,
	}

	return handle
}

func (m *HandlerExcepcion) Send(w http.ResponseWriter, r *http.Request, err error) {

	idRequest := middleware.GetReqID(r.Context())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	response := ErrorResponse{
		Mensaje:   err.Error(),
		RequestId: idRequest,
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(response); err != nil {
		m.logger.Error(err.Error(), zap.String("id-request", idRequest), zap.String("category", "general"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	re, ok := err.(*excepciones.ErrorCustom)

	if ok {
		m.logger.Error(err.Error(), zap.String("id-request", idRequest), zap.String("category", re.Categoria()))
		w.WriteHeader(re.Codigo())

	} else {

		m.logger.Error(err.Error(), zap.String("id-request", idRequest), zap.String("category", "general"))
		w.WriteHeader(500)

	}

	w.WriteHeader(400)
	w.Write(buf.Bytes())

}
