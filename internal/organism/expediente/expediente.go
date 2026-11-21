package expediente

import (
	"time"

	"github.com/yecharlot/guantanamo-digital-ready/internal/organism"
)

// ExpedienteDigital es el documento central exigido por la Resolución 24/2026.
type ExpedienteDigital struct {
	Version          string         `json:"version"`
	EntityID         string         `json:"entity_id"`
	EntityName       string         `json:"entity_name"`
	FechaGeneracion  string         `json:"fecha_generacion"`
	DiagnosticoRef   string         `json:"diagnostico_ref"`
	InventarioRef    string         `json:"inventario_ref"`
	PuntajeMadurez   float64        `json:"puntaje_madurez"`
	NivelMadurez     string         `json:"nivel_madurez"`
	ResumenEjecutivo string         `json:"resumen_ejecutivo"`
	HallazgosClave   []string       `json:"hallazgos_clave"`
	BrechasCriticas  []string       `json:"brechas_criticas"`
	ProximosPasos    []string       `json:"proximos_pasos"`
	Metadata         map[string]any `json:"metadata"`
}

// Generar crea un expediente a partir de los resultados de diagnóstico e inventario.
func Generar(entityID, entityName, diagRef, invRef string, puntaje float64, nivel string, hallazgos, brechas, pasos []string) *ExpedienteDigital {
	return &ExpedienteDigital{
		Version:          "1.0",
		EntityID:         entityID,
		EntityName:       entityName,
		FechaGeneracion:  time.Now().UTC().Format(time.RFC3339),
		DiagnosticoRef:   diagRef,
		InventarioRef:    invRef,
		PuntajeMadurez:   puntaje,
		NivelMadurez:     nivel,
		ResumenEjecutivo: "Expediente Digital de Arranque generado conforme a la Resolución 24/2026 del Ministerio de Comunicaciones.",
		HallazgosClave:   hallazgos,
		BrechasCriticas:  brechas,
		ProximosPasos:    pasos,
		Metadata:         map[string]any{"generado_por": "GDR-MVP-1.0"},
	}
}

func CrearOrganismo(entityID, entityName string) *organism.Organism {
	o := organism.New("expediente", "Expediente Digital de Arranque – "+entityName, entityID)
	o.Status = organism.StatusReady
	o.Data["entity_name"] = entityName
	return o
}
