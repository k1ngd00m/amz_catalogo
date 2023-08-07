package adaptador

import (
	"database/sql"

	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/dao"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/evento"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/repositorio"
	daoInfra "github.com/k1ngd00m/amz_catalogo/infra/adaptador/dao"
	eventoinfra "github.com/k1ngd00m/amz_catalogo/infra/adaptador/evento"
	repositorioInfra "github.com/k1ngd00m/amz_catalogo/infra/adaptador/repositorio"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Adaptador struct {
	RepositorioProducto      *repositorio.RepositorioProducto
	DaoProducto              *dao.DaoProducto
	EventoRegistrarProducto  *evento.EventoRegistrarProducto
	EventoActualizarProducto *evento.EventoActualizarProducto
}

func NewAdaptador(dbWrite *sql.DB, dbReader *mongo.Client, conn *amqp.Connection) *Adaptador {

	repositorioProducto := repositorioInfra.NewMysqlRepositorioProducto(dbWrite)

	daoProducto := daoInfra.NewMongoDaoProducto(dbReader)

	eventoRegistrarProducto := eventoinfra.NewRabbitEventoRegistrarProducto(conn)
	eventoActualizarProducto := eventoinfra.NewRabbitEventoActualizarProducto(conn)

	return &Adaptador{
		RepositorioProducto:      &repositorioProducto,
		DaoProducto:              &daoProducto,
		EventoRegistrarProducto:  &eventoRegistrarProducto,
		EventoActualizarProducto: &eventoActualizarProducto,
	}

}
