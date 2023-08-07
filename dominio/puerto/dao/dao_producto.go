package dao

import "github.com/k1ngd00m/amz_catalogo/aplicacion/dto"

type DaoProducto interface {
	BuscarPorPaginado(pagina int, cantidad int) (*[]dto.DtoProducto, error)
}
