package evento

import "github.com/k1ngd00m/amz_catalogo/dominio/entidad"

type EventoActualizarProducto interface {
	Ejecutar(producto *entidad.Producto) error
}
