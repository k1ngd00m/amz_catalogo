package repositorioInfra

import (
	"database/sql"

	"github.com/k1ngd00m/amz_catalogo/dominio/entidad"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/repositorio"
)

type MysqlRepositorioProducto struct {
	db *sql.DB
}

func NewMysqlRepositorioProducto(db *sql.DB) repositorio.RepositorioProducto {

	repo := &MysqlRepositorioProducto{
		db: db,
	}

	return repo
}

func (m *MysqlRepositorioProducto) Registrar(producto *entidad.Producto) error {

	stmt, err := m.db.Prepare("INSERT INTO producto(id_producto, nombre, descripcion, id_categoria, stock, estado) VALUES (?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(producto.ID, producto.Nombre, producto.Descripcion, producto.IdCategoria, producto.Stock, producto.Estado)

	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil

}
