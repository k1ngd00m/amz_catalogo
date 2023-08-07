package eventoinfra

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/k1ngd00m/amz_catalogo/dominio/entidad"
	"github.com/k1ngd00m/amz_catalogo/dominio/puerto/evento"
	amqp "github.com/rabbitmq/amqp091-go"
)

const exchangeName = "amz-exchange"
const key = "write_catalogo"

type RabbitEventoRegistrarProducto struct {
	conn *amqp.Connection
}

type RegistrarProductoRequest struct {
	Id          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	IdCategoria string `json:"idCategoria"`
	Stock       int    `json:"stock"`
	Estado      int    `json:"estado"`
}

func NewRabbitEventoRegistrarProducto(conn *amqp.Connection) evento.EventoRegistrarProducto {
	eventoRegistrarProducto := &RabbitEventoRegistrarProducto{
		conn: conn,
	}

	return eventoRegistrarProducto
}

func (m *RabbitEventoRegistrarProducto) Ejecutar(producto *entidad.Producto) error {

	channel, err := m.conn.Channel()

	if err != nil {
		return err
	}

	defer channel.Close()

	request := &RegistrarProductoRequest{
		Id:          producto.ID,
		Nombre:      producto.Nombre,
		Descripcion: producto.Descripcion,
		IdCategoria: producto.IdCategoria,
		Stock:       producto.Stock,
		Estado:      producto.Estado,
	}

	body, err := json.Marshal(request)

	if err != nil {
		return err
	}

	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)

	if err = channel.Publish(
		exchangeName, // publish to an exchange
		key,          // routing to 0 or more queues
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers: amqp.Table{
				"accion": "registrar_producto",
			},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Persistent, // 1=non-persistent, 2=persistent
			Priority:        0,               // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}
