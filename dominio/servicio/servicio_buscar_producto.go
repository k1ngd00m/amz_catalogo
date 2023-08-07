package servicio

import (
	"github.com/k1ngd00m/amz_catalogo/aplicacion/dto"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/dao"
)

type ServicioBuscarProducto struct {
	daoProducto dao.DaoProducto
}

func NewServicioBuscarProducto(daoProducto dao.DaoProducto) ServicioBuscarProducto {
	return ServicioBuscarProducto{
		daoProducto: daoProducto,
	}
}

func (m *ServicioBuscarProducto) Ejecutar(pagina int, cantidad int) (*[]dto.DtoProducto, error) {
	productos, err := m.daoProducto.BuscarPorPaginado(pagina, cantidad)
	if err != nil {
		return nil, err
	}

	return productos, nil
}
