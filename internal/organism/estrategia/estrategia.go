package estrategia

import (
	"github.com/yecharlot/guantanamo-digital-ready/internal/organism"
)

// Recomendacion es una acción priorizada del roadmap.
type Recomendacion struct {
	ID          string `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Prioridad   string `json:"prioridad"` // alta | media | baja
	Horizonte   string `json:"horizonte"` // 2026-2027 | 2028 | 2029-2030
	Dimension   string `json:"dimension"`
}

// Roadmap contiene las recomendaciones generadas.
type Roadmap struct {
	EntityID        string          `json:"entity_id"`
	EntityName      string          `json:"entity_name"`
	Recomendaciones []Recomendacion `json:"recomendaciones"`
	Resumen         string          `json:"resumen"`
}

// GenerarBasico produce un conjunto inicial de recomendaciones según el nivel de madurez.
func GenerarBasico(entityID, entityName, nivel string) *Roadmap {
	recs := []Recomendacion{
		{
			ID:          "R-01",
			Titulo:      "Conformar el Comité / Oficina de Transformación Digital",
			Descripcion: "Designar responsables y definir roles según la Resolución 24/2026.",
			Prioridad:   "alta",
			Horizonte:   "2026-2027",
			Dimension:   "Gobernanza y liderazgo digital",
		},
		{
			ID:          "R-02",
			Titulo:      "Completar el inventario de datos institucionales",
			Descripcion: "Catalogar los principales conjuntos de datos, su sensibilidad y calidad. Principio: Primero los datos.",
			Prioridad:   "alta",
			Horizonte:   "2026-2027",
			Dimension:   "Disponibilidad y calidad de datos",
		},
		{
			ID:          "R-03",
			Titulo:      "Evaluar infraestructura y conectividad crítica",
			Descripcion: "Identificar puntos únicos de fallo y necesidades mínimas de continuidad.",
			Prioridad:   "alta",
			Horizonte:   "2026-2027",
			Dimension:   "Infraestructura tecnológica",
		},
		{
			ID:          "R-04",
			Titulo:      "Capacitación básica en transformación digital",
			Descripcion: "Programa de sensibilización y formación para el personal clave.",
			Prioridad:   "media",
			Horizonte:   "2026-2027",
			Dimension:   "Capacidades del personal",
		},
		{
			ID:          "R-05",
			Titulo:      "Definir política institucional de datos",
			Descripcion: "Establecer reglas de acceso, calidad, retención y protección de datos personales.",
			Prioridad:   "alta",
			Horizonte:   "2026-2027",
			Dimension:   "Marco normativo y cumplimiento",
		},
	}

	return &Roadmap{
		EntityID:        entityID,
		EntityName:      entityName,
		Recomendaciones: recs,
		Resumen:         "Roadmap inicial orientado al Horizonte 1 (2026-2027) de la Resolución 24/2026. Enfocado en capacidades habilitadoras.",
	}
}

func CrearOrganismo(entityID, entityName string) *organism.Organism {
	o := organism.New("estrategia", "Estrategia TD e IA – "+entityName, entityID)
	o.Status = organism.StatusReady
	o.Data["entity_name"] = entityName
	return o
}
