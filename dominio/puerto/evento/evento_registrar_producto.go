package evento

import "github.com/k1ngd00m/amz_catalogo/dominio/entidad"

type EventoRegistrarProducto interface {
	Ejecutar(producto *entidad.Producto) error
}
