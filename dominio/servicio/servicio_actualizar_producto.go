package servicio

import (
	"github.com/k1ngd00m/amz_catalogo/dominio/entidad"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/evento"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/repositorio"
	"github.com/k1ngd00m/amz_comun/excepciones"
)

type ServicioActualizarProducto struct {
	repositorioProducto      repositorio.RepositorioProducto
	eventoActualizarProducto evento.EventoActualizarProducto
}

const noExisteElProductoMsg = "no existe el producto"

func NewServicioActualizarProducto(repositorioProducto *repositorio.RepositorioProducto,
	eventoActualizarProducto *evento.EventoActualizarProducto) ServicioActualizarProducto {

	return ServicioActualizarProducto{
		repositorioProducto:      *repositorioProducto,
		eventoActualizarProducto: *eventoActualizarProducto,
	}
}

func (m *ServicioActualizarProducto) Ejecutar(entidad *entidad.Producto) error {

	cantidadRegistros, err := m.repositorioProducto.ContarPorId(entidad.ID)

	if err != nil {
		return err
	}

	if cantidadRegistros == 0 {
		return excepciones.DataInvalida(noExisteElProductoMsg)
	}

	err = m.repositorioProducto.Actualizar(entidad)

	if err != nil {
		return err
	}

	err = m.eventoActualizarProducto.Ejecutar(entidad)

	if err != nil {
		return err
	}

	return nil

}
