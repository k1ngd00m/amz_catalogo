package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/dto"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/manejador"
)

type RestQueryProducto struct {
	manejadorBuscarProducto *manejador.ManejadorBuscarProducto
	handlerExcepcion        *HandlerExcepcion
}

func NewRestQueryProducto(handlerExcepcion *HandlerExcepcion, manejadorBuscarProducto *manejador.ManejadorBuscarProducto) RestQueryProducto {

	return *&RestQueryProducto{
		manejadorBuscarProducto: manejadorBuscarProducto,
		handlerExcepcion:        handlerExcepcion,
	}
}

func (m *RestQueryProducto) Routes() chi.Router {

	router := chi.NewRouter()
	router.Get("/producto", m.buscarProductos)

	return router

}

func (m *RestQueryProducto) buscarProductos(w http.ResponseWriter, r *http.Request) {

	var numeroRegistros int
	var productos *[]dto.DtoProducto

	pagina := r.URL.Query().Get("pagina")
	cantidad := r.URL.Query().Get("cantidad")

	numeroPagina, err := strconv.Atoi(pagina)

	if err != nil {
		m.handlerExcepcion.Send(w, r, err)
		return
	}

	numeroRegistros, err = strconv.Atoi(cantidad)

	if err != nil {
		m.handlerExcepcion.Send(w, r, err)
		return
	}

	productos, err = m.manejadorBuscarProducto.Ejecutar(numeroPagina, numeroRegistros)

	render.JSON(w, r, productos)
}
