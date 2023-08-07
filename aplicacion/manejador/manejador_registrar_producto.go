package manejador

import (
	"github.com/google/uuid"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/comando"
	"github.com/k1ngd00m/amz_catalogo/dominio/entidad"
	"github.com/k1ngd00m/amz_catalogo/dominio/servicio"
)

type ManejadorRegistrarProducto struct {
	servicioRegistarProducto servicio.ServicioRegistrarProducto
}

func NewManejadorRegistarProducto(servicio *servicio.ServicioRegistrarProducto) ManejadorRegistrarProducto {
	return ManejadorRegistrarProducto{
		servicioRegistarProducto: *servicio,
	}
}

func (m *ManejadorRegistrarProducto) Ejecutar(comando *comando.ComandoRegistrarProducto) error {

	id := uuid.New()

	entidad, err := entidad.NewEntidadProducto(
		id.String(),
		comando.Nombre,
		comando.Descripcion,
		comando.IdCategoria,
		comando.Stock,
		1)

	if err != nil {
		return err
	}

	err = m.servicioRegistarProducto.Ejecutar(entidad)

	if err != nil {
		return err
	}

	return nil

}
