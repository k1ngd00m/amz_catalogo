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

func (m *MysqlRepositorioProducto) ContarPorId(idProducto string) (int, error) {
	stmt, err := m.db.Prepare("SELECT COUNT(*) FROM producto WHERE id_producto = ?")

	if err != nil {
		return 0, err
	}

	var cantidadRegistros = 0

	err = stmt.QueryRow(idProducto).Scan(&cantidadRegistros)

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	return cantidadRegistros, nil

}

func (m *MysqlRepositorioProducto) Actualizar(producto *entidad.Producto) error {
	stmt, err := m.db.Prepare("UPDATE producto SET nombre = ?, descripcion = ?, id_categoria = ?, stock = ?, estado = ? WHERE id_producto = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(producto.Nombre, producto.Descripcion, producto.IdCategoria, producto.Stock, producto.Estado, producto.ID)

	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil

}
