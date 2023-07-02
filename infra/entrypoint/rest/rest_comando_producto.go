package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/comando"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/manejador"
)

type RestComandoProducto struct {
	handlerExcepcion          *HandlerExcepcion
	manejadorRegistarProducto *manejador.ManejadorRegistrarProducto
}

func NewRestComandoOrden(handlerExcepcion *HandlerExcepcion, manejadorRegistrarProducto *manejador.ManejadorRegistrarProducto) RestComandoProducto {

	restComando := RestComandoProducto{
		handlerExcepcion:          handlerExcepcion,
		manejadorRegistarProducto: manejadorRegistrarProducto,
	}

	return restComando

}

func (m *RestComandoProducto) Routes() chi.Router {

	router := chi.NewRouter()
	router.Post("/", m.registrarProducto)

	return router

}

func (m *RestComandoProducto) registrarProducto(w http.ResponseWriter, r *http.Request) {
	producto := &comando.ComandoRegistrarProducto{}

	if err := render.Bind(r, producto); err != nil {
		m.handlerExcepcion.Send(w, r, err)
		return
	}

	if err := m.manejadorRegistarProducto.Ejecutar(producto); err != nil {
		m.handlerExcepcion.Send(w, r, err)
		return
	}
}
