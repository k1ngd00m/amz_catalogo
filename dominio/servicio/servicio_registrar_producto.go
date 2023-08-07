package servicio

import (
	"github.com/k1ngd00m/amz_catalogo/dominio/entidad"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/evento"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/repositorio"
)

type ServicioRegistrarProducto struct {
	repositorioProducto     repositorio.RepositorioProducto
	eventoRegistrarProducto evento.EventoRegistrarProducto
}

func NewServicioRegistarProducto(repositorioProducto repositorio.RepositorioProducto,
	eventoRegistrarProducto evento.EventoRegistrarProducto) ServicioRegistrarProducto {

	return ServicioRegistrarProducto{
		repositorioProducto:     repositorioProducto,
		eventoRegistrarProducto: eventoRegistrarProducto,
	}
}

func (m *ServicioRegistrarProducto) Ejecutar(producto *entidad.Producto) error {
	err := m.repositorioProducto.Registrar(producto)

	if err != nil {
		return err
	}

	err = m.eventoRegistrarProducto.Ejecutar(producto)

	if err != nil {
		return err
	}

	return nil
}
