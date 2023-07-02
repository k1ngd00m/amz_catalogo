package entidad

import validadores "github.com/k1ngd00m/amz_comun/validador"

const idententificadoInvalidoMsg = "se debe ingresar un identificador valido"
const nombreInvalidoMsg = "Se debe ingresar un nombre de producto valido"
const descripcionInvalidaMsg = "Se debe ingresar una descripcion valida"

type Producto struct {
	ID          string
	Nombre      string
	Descripcion string
	IdCategoria string
	Stock       int
	Estado      int
}

func NewEntidadProducto(id string, nombre string, descripcion string, idCategoria string, stock int, estado int) (*Producto, error) {

	error := validadores.StringObligatorio(id, idententificadoInvalidoMsg)

	if error != nil {
		return nil, error
	}

	error = validadores.StringObligatorio(nombre, nombreInvalidoMsg)

	if error != nil {
		return nil, error
	}

	error = validadores.StringObligatorio(descripcion, descripcionInvalidaMsg)

	if error != nil {
		return nil, error
	}

	producto := new(Producto)

	producto.ID = id
	producto.Nombre = nombre
	producto.Descripcion = descripcion
	producto.Stock = stock
	producto.Estado = estado
	producto.IdCategoria = idCategoria

	return producto, nil

}
