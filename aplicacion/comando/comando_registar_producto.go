package comando

import "net/http"

type ComandoRegistrarProducto struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	IdCategoria string `json:"id_categoria"`
	Stock       int    `json:"stock"`
}

func (m *ComandoRegistrarProducto) Bind(r *http.Request) error {
	return nil
}
