package servicio

import (
	"github.com/k1ngd00m/amz_catalogo/dominio/entidad"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/repositorio"
)

type ServicioRegistrarProducto struct {
	repositorioProducto repositorio.RepositorioProducto
}

func NewServicioRegistarProducto(repositorioProducto repositorio.RepositorioProducto) ServicioRegistrarProducto {

	return *&ServicioRegistrarProducto{
		repositorioProducto: repositorioProducto,
	}
}

func (m *ServicioRegistrarProducto) Ejecutar(producto *entidad.Producto) error {
	error := m.repositorioProducto.Registrar(producto)

	if error != nil {
		return error
	}
	return nil
}
