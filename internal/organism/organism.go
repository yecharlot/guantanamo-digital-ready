package organism

import (
	"time"
)

// Status represents the lifecycle state of a Digital Organism.
type Status string

const (
	StatusCreated   Status = "created"
	StatusReady     Status = "ready"
	StatusRunning   Status = "running"
	StatusStopped   Status = "stopped"
	StatusCompleted Status = "completed"
)

// Organism is the base Digital Organism used by GDR.
// Inspired by PrismaTec-Core patterns (RootCID, identity, memory, policy).
type Organism struct {
	ID        string            `json:"id"`
	RootCID   string            `json:"root_cid"`
	Name      string            `json:"name"`
	Kind      string            `json:"kind"` // diagnostico | inventario | expediente | estrategia
	Version   string            `json:"version"`
	Status    Status            `json:"status"`
	EntityID  string            `json:"entity_id"` // entidad pública (municipio, provincia...)
	Data      map[string]any    `json:"data"`
	Metadata  map[string]string `json:"metadata"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func New(kind, name, entityID string) *Organism {
	now := time.Now().UTC()
	id := kind + "-" + now.Format("20060102-150405")
	return &Organism{
		ID:        id,
		RootCID:   "rootcid:" + id,
		Name:      name,
		Kind:      kind,
		Version:   "1.0.0",
		Status:    StatusCreated,
		EntityID:  entityID,
		Data:      make(map[string]any),
		Metadata:  make(map[string]string),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (o *Organism) Touch() {
	o.UpdatedAt = time.Now().UTC()
}
