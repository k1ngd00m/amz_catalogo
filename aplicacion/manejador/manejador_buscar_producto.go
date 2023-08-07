package manejador

import (
	"github.com/k1ngd00m/amz_catalogo/aplicacion/dto"
	"github.com/k1ngd00m/amz_catalogo/dominio/servicio"
)

type ManejadorBuscarProducto struct {
	servicioBuscarProducto servicio.ServicioBuscarProducto
}

func NewManejadorBuscarProducto(servicioBuscarProducto *servicio.ServicioBuscarProducto) ManejadorBuscarProducto {
	return ManejadorBuscarProducto{
		servicioBuscarProducto: *servicioBuscarProducto,
	}
}

func (m *ManejadorBuscarProducto) Ejecutar(pagina int, cantidad int) (*[]dto.DtoProducto, error) {
	return m.servicioBuscarProducto.Ejecutar(pagina, cantidad)
}
