package infra

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/k1ngd00m/amz_catalogo/aplicacion/manejador"
	"github.com/k1ngd00m/amz_catalogo/dominio/servicio"
	"github.com/k1ngd00m/amz_catalogo/infra/adaptador"
	"github.com/k1ngd00m/amz_catalogo/infra/config"
	"github.com/k1ngd00m/amz_catalogo/infra/entrypoint/rest"
)

func Start() {

	// inicializar variables de entorno
	config.LoadEnvironment()

	log := config.GetLogger()

	// inicializar conexion a rdms write
	dbWrite := config.GetWriteDBConexion()
	log.Info("conexion a write catalogo db existosa")

	// inicializar conexion a rdms read
	dbRead := config.GetReadDBConexion()
	log.Info("conexion a read catalogo db existosa")

	// inicializar rabbitMQ
	broker := config.GetRabbitMqConn()

	// iniciando adaptadores
	dependencias := adaptador.NewAdaptador(dbWrite, dbRead, broker)

	// servicios
	servicioRegistrarProducto := servicio.NewServicioRegistarProducto(*dependencias.RepositorioProducto, *dependencias.EventoRegistrarProducto)
	servicioActualizarProducto := servicio.NewServicioActualizarProducto(dependencias.RepositorioProducto, dependencias.EventoActualizarProducto)
	servicioBuscarProducto := servicio.NewServicioBuscarProducto(*dependencias.DaoProducto)

	// manejadores
	manejadorRegistrarProducto := manejador.NewManejadorRegistarProducto(&servicioRegistrarProducto)
	manejadorActualizarProducto := manejador.NewManejadorActualizarProducto(&servicioActualizarProducto)
	manejadorBuscarProducto := manejador.NewManejadorBuscarProducto(&servicioBuscarProducto)

	app := chi.NewRouter()

	// Default middleware config
	app.Use(middleware.RequestID)

	handlerExcepcion := rest.NewHandlerExcepcion(log)

	restComandoProducto := rest.NewRestComandoOrden(&handlerExcepcion, &manejadorRegistrarProducto, &manejadorActualizarProducto)
	restQueryProducto := rest.NewRestQueryProducto(&handlerExcepcion, &manejadorBuscarProducto)

	app.Mount("/catalogo/comand", restComandoProducto.Routes())
	app.Mount("/catalogo/query", restQueryProducto.Routes())

	http.ListenAndServe(":3200", app)

}
