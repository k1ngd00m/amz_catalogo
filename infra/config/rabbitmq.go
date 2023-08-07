package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func GetRabbitMqConn() *amqp.Connection {
	// Nombre de usuario asignado por RabbitMQ
	var user string = os.Getenv("RABBITMQ_USER")
	// Contraseña del usuario de RabbitMQ
	var pwd string = os.Getenv("RABBITMQ_PASSWORD")
	// La dirección IP de RabbitMQ Broker
	var host string = os.Getenv("RABBITMQ_SERVER")
	// Puerto monitoreado por RabbitMQ Broker
	var port string = os.Getenv("RABBITMQ_PORT")
	url := "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/" + user
	// Crea una nueva conexión
	conn, err := amqp.Dial(url)
	// devuelve conexión y error

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
