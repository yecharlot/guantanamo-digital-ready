package inventario

import (
	"github.com/yecharlot/guantanamo-digital-ready/internal/organism"
)

// Sensibilidad de un conjunto de datos.
type Sensibilidad string

const (
	SensibilidadPublica     Sensibilidad = "publica"
	SensibilidadInterna     Sensibilidad = "interna"
	SensibilidadRestringida Sensibilidad = "restringida"
	SensibilidadPersonal    Sensibilidad = "datos_personales"
)

// Calidad estimada del dato.
type Calidad string

const (
	CalidadDesconocida Calidad = "desconocida"
	CalidadBaja        Calidad = "baja"
	CalidadMedia       Calidad = "media"
	CalidadAlta        Calidad = "alta"
)

// ConjuntoDatos representa un activo de datos de la entidad.
type ConjuntoDatos struct {
	ID           string       `json:"id"`
	Nombre       string       `json:"nombre"`
	Descripcion  string       `json:"descripcion"`
	Propietario  string       `json:"propietario"`
	Sensibilidad Sensibilidad `json:"sensibilidad"`
	Calidad      Calidad      `json:"calidad"`
	Formato      string       `json:"formato"` // excel, pdf, sistema, papel...
	Ubicacion    string       `json:"ubicacion"`
	Actualizado  string       `json:"actualizado"`
}

// CrearOrganismo inicializa el organismo de inventario de datos.
func CrearOrganismo(entityID, entityName string) *organism.Organism {
	o := organism.New("inventario", "Inventario de Datos – "+entityName, entityID)
	o.Status = organism.StatusReady
	o.Data["entity_name"] = entityName
	o.Data["conjuntos"] = []ConjuntoDatos{}
	return o
}
