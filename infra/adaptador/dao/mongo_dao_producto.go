package daoInfra

import (
	"context"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/dto"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/dao"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDaoProducto struct {
	cliente *mongo.Client
}

func NewMongoDaoProducto(cliente *mongo.Client) dao.DaoProducto {
	daoProducto := &MongoDaoProducto{
		cliente: cliente,
	}

	return daoProducto
}

func (m *MongoDaoProducto) BuscarPorPaginado(pagina int, cantidad int) (*[]dto.DtoProducto, error) {

	coleccion := m.cliente.Database("catalogo").Collection("productos")

	filter := bson.D{}
	//opts := options.Find().SetLimit(int64(cantidad)).SetSkip(int64(pagina))

	cursor, err := coleccion.Find(context.TODO(), filter)
	var data []dto.DtoProducto

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}

	return &data, nil

}
