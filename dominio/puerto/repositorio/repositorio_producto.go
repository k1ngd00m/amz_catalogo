package repositorio

import "github.com/k1ngd00m/amz_catalogo/dominio/entidad"

type RepositorioProducto interface {
	Registrar(producto *entidad.Producto) error
	ContarPorId(idProducto string) (int, error)
	Actualizar(producto *entidad.Producto) error
}
