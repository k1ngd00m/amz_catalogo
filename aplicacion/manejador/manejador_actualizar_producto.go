package manejador

import (
	"github.com/k1ngd00m/amz_catalogo/aplicacion/comando"
	"github.com/k1ngd00m/amz_catalogo/dominio/entidad"
	"github.com/k1ngd00m/amz_catalogo/dominio/servicio"
)

type ManejadorActualizarProducto struct {
	servicioActualizarProducto servicio.ServicioActualizarProducto
}

func NewManejadorActualizarProducto(servicioActualizarProducto *servicio.ServicioActualizarProducto) ManejadorActualizarProducto {
	return ManejadorActualizarProducto{
		servicioActualizarProducto: *servicioActualizarProducto,
	}
}

func (m *ManejadorActualizarProducto) Ejecutar(comando *comando.ComandoActualizarProducto) error {
	producto, err := entidad.NewEntidadProducto(comando.ID,
		comando.Nombre,
		comando.Descripcion,
		comando.IdCategoria,
		comando.Stock,
		comando.Estado)

	if err != nil {
		return err
	}

	err = m.servicioActualizarProducto.Ejecutar(producto)

	if err != nil {
		return err
	}

	return nil
}
