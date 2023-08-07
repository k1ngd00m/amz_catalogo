package comando

import "net/http"

type ComandoActualizarProducto struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	IdCategoria string `json:"id_categoria"`
	Stock       int    `json:"stock"`
	Estado      int    `json:"estado"`
}

func (m *ComandoActualizarProducto) Bind(r *http.Request) error {
	return nil
}
