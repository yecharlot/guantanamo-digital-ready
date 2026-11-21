package diagnostico

import (
	"github.com/yecharlot/guantanamo-digital-ready/internal/organism"
)

// Dimensiones de madurez digital según la Resolución 24/2026
// (adaptadas a las 10 áreas de evaluación mencionadas en la metodología).
var Dimensiones = []string{
	"Gobernanza y liderazgo digital",
	"Infraestructura tecnológica",
	"Disponibilidad y calidad de datos",
	"Seguridad informática y ciberseguridad",
	"Procesos de trabajo digitalizados",
	"Servicios al ciudadano",
	"Capacidades del personal",
	"Marco normativo y cumplimiento",
	"Interoperabilidad",
	"Preparación para IA (Pirámide Funcional)",
}

// Nivel representa el nivel de madurez de 0 a 5.
type Nivel int

const (
	NivelInexistente Nivel = 0
	NivelInicial     Nivel = 1
	NivelBasico      Nivel = 2
	NivelIntermedio  Nivel = 3
	NivelAvanzado    Nivel = 4
	NivelOptimo      Nivel = 5
)

// Respuesta almacena la evaluación de una dimensión.
type Respuesta struct {
	Dimension string `json:"dimension"`
	Nivel     Nivel  `json:"nivel"`
	Evidencia string `json:"evidencia"`
	Notas     string `json:"notas"`
}

// Resultado contiene el diagnóstico completo.
type Resultado struct {
	OrganismID   string      `json:"organism_id"`
	EntityID     string      `json:"entity_id"`
	EntityName   string      `json:"entity_name"`
	Respuestas   []Respuesta `json:"respuestas"`
	PuntajeTotal float64     `json:"puntaje_total"` // 0-100
	NivelGlobal  string      `json:"nivel_global"`
	Fecha        string      `json:"fecha"`
}

// Evaluar calcula el puntaje y nivel global.
func Evaluar(respuestas []Respuesta) (float64, string) {
	if len(respuestas) == 0 {
		return 0, "Sin datos"
	}
	suma := 0
	for _, r := range respuestas {
		suma += int(r.Nivel)
	}
	max := len(respuestas) * 5
	puntaje := (float64(suma) / float64(max)) * 100

	var nivel string
	switch {
	case puntaje < 20:
		nivel = "Inicial / Crítico"
	case puntaje < 40:
		nivel = "Básico"
	case puntaje < 60:
		nivel = "Intermedio"
	case puntaje < 80:
		nivel = "Avanzado"
	default:
		nivel = "Óptimo"
	}
	return puntaje, nivel
}

// CrearOrganismo inicializa un organismo de diagnóstico.
func CrearOrganismo(entityID, entityName string) *organism.Organism {
	o := organism.New("diagnostico", "Diagnóstico de Madurez Digital – "+entityName, entityID)
	o.Status = organism.StatusReady
	o.Data["entity_name"] = entityName
	o.Data["dimensiones"] = Dimensiones
	o.Data["respuestas"] = []Respuesta{}
	return o
}
