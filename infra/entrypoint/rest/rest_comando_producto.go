package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/comando"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/manejador"
)

type RestComandoProducto struct {
	handlerExcepcion            *HandlerExcepcion
	manejadorRegistarProducto   *manejador.ManejadorRegistrarProducto
	manejadorActualizarProducto *manejador.ManejadorActualizarProducto
}

func NewRestComandoOrden(handlerExcepcion *HandlerExcepcion,
	manejadorRegistrarProducto *manejador.ManejadorRegistrarProducto,
	manejadorActualizarProducto *manejador.ManejadorActualizarProducto) RestComandoProducto {

	restComando := RestComandoProducto{
		handlerExcepcion:            handlerExcepcion,
		manejadorRegistarProducto:   manejadorRegistrarProducto,
		manejadorActualizarProducto: manejadorActualizarProducto,
	}

	return restComando

}

func (m *RestComandoProducto) Routes() chi.Router {

	router := chi.NewRouter()
	router.Post("/producto", m.registrarProducto)
	router.Put("/producto", m.actualizarProducto)

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

func (m *RestComandoProducto) actualizarProducto(w http.ResponseWriter, r *http.Request) {
	producto := &comando.ComandoActualizarProducto{}

	if err := render.Bind(r, producto); err != nil {
		m.handlerExcepcion.Send(w, r, err)
		return
	}

	if err := m.manejadorActualizarProducto.Ejecutar(producto); err != nil {
		m.handlerExcepcion.Send(w, r, err)
		return
	}
}
