package dto

type DtoCategoria struct {
	ID     string `json:"id" bson:"_id"`
	Nombre string `json:"nombre" bson:"nombre"`
}

type DtoProducto struct {
	ID          string       `json:"id" bson:"_id"`
	Nombre      string       `json:"nombre" bson:"nombre"`
	Descripcion string       `json:"descripcion" bson:"descripcion"`
	Stock       int          `json:"stock" bson:"stock"`
	Estado      int          `json:"estado" bson:"estado"`
	Categoria   DtoCategoria `json:"categoria" bson:"categoria"`
}
