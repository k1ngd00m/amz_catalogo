package adaptador

import (
	"database/sql"

	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/repositorio"
	repositorioInfra "github.com/k1ngd00m/amz_catalogo/infra/adaptador/repositorio"
)

type Adaptador struct {
	RepositorioProducto *repositorio.RepositorioProducto
}

func NewAdaptador(db *sql.DB) *Adaptador {

	repositorioProducto := repositorioInfra.NewMysqlRepositorioProducto(db)

	return &Adaptador{
		RepositorioProducto: &repositorioProducto,
	}

}
