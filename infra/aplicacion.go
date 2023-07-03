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

	// inicializar conexion a rdms
	db := config.GetDBConexion()
	log.Info("conexion a write catalogo db existosa")

	// iniciando adaptadores
	dependencias := adaptador.NewAdaptador(db)

	// servicios
	servicioRegistrarProducto := servicio.NewServicioRegistarProducto(*dependencias.RepositorioProducto)

	// manejadores
	manejadorRegistrarProducto := manejador.NewManejadorRegistarProducto(&servicioRegistrarProducto)

	app := chi.NewRouter()

	// Default middleware config
	app.Use(middleware.RequestID)

	handlerExcepcion := rest.NewHandlerExcepcion(log)

	restComandoProducto := rest.NewRestComandoOrden(&handlerExcepcion, &manejadorRegistrarProducto)

	app.Mount("/catalogo/comand", restComandoProducto.Routes())

	http.ListenAndServe(":3200", app)

}
