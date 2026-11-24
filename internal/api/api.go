package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yecharlot/guantanamo-digital-ready/internal/audit"
	"github.com/yecharlot/guantanamo-digital-ready/internal/organism"
	"github.com/yecharlot/guantanamo-digital-ready/internal/organism/diagnostico"
	"github.com/yecharlot/guantanamo-digital-ready/internal/organism/estrategia"
	"github.com/yecharlot/guantanamo-digital-ready/internal/organism/expediente"
	"github.com/yecharlot/guantanamo-digital-ready/internal/organism/inventario"
	"github.com/yecharlot/guantanamo-digital-ready/internal/policy"
	"github.com/yecharlot/guantanamo-digital-ready/internal/store"
)

type Server struct {
	Store  *store.Store
	Audit  *audit.Log
	Policy *policy.Engine
	Mux    *http.ServeMux
}

func New(s *store.Store, a *audit.Log, p *policy.Engine) *Server {
	srv := &Server{Store: s, Audit: a, Policy: p, Mux: http.NewServeMux()}
	srv.routes()
	return srv
}

func (s *Server) routes() {
	s.Mux.HandleFunc("/api/health", s.handleHealth)
	s.Mux.HandleFunc("/api/diagnostico", s.handleDiagnostico)
	s.Mux.HandleFunc("/api/diagnostico/evaluar", s.handleEvaluar)
	s.Mux.HandleFunc("/api/diagnostico/criterios", s.handleCriterios)
	s.Mux.HandleFunc("/api/inventario", s.handleInventario)
	s.Mux.HandleFunc("/api/expediente", s.handleExpediente)
	s.Mux.HandleFunc("/api/estrategia", s.handleEstrategia)
	s.Mux.HandleFunc("/api/organismos", s.handleListOrganismos)
	s.Mux.HandleFunc("/", s.handleFrontend)
}


func (s *Server) handleCriterios(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"dimensiones": diagnostico.Dimensiones,
		"guias":       diagnostico.Guias,
		"leyenda_niveles": []map[string]any{
			{"nivel": 0, "etiqueta": "Inexistente"},
			{"nivel": 1, "etiqueta": "Inicial"},
			{"nivel": 2, "etiqueta": "Básico"},
			{"nivel": 3, "etiqueta": "Intermedio"},
			{"nivel": 4, "etiqueta": "Avanzado"},
			{"nivel": 5, "etiqueta": "Óptimo"},
		},
	})
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"status":    "ok",
		"service":   "Guantánamo Digital Ready",
		"version":   "1.0.0-mvp",
		"timestamp": time.Now().UTC(),
		"offline":   true,
	})
}

func (s *Server) handleDiagnostico(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body struct {
			EntityID   string `json:"entity_id"`
			EntityName string `json:"entity_name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		o := diagnostico.CrearOrganismo(body.EntityID, body.EntityName)
		if err := s.Store.Save(o); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		_ = s.Audit.Record("system", "create", o.ID, map[string]string{"kind": "diagnostico"})
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(o)
		return
	}
	http.Error(w, "method not allowed", 405)
}

func (s *Server) handleEvaluar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	var body struct {
		OrganismID string                  `json:"organism_id"`
		EntityID   string                  `json:"entity_id"`
		EntityName string                  `json:"entity_name"`
		Respuestas []diagnostico.Respuesta `json:"respuestas"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	puntaje, nivel := diagnostico.Evaluar(body.Respuestas)
	result := diagnostico.Resultado{
		OrganismID:   body.OrganismID,
		EntityID:     body.EntityID,
		EntityName:   body.EntityName,
		Respuestas:   body.Respuestas,
		PuntajeTotal: puntaje,
		NivelGlobal:  nivel,
		Fecha:        time.Now().UTC().Format(time.RFC3339),
	}

	// Update organism
	if o, err := s.Store.Load(body.OrganismID); err == nil {
		o.Data["resultado"] = result
		o.Data["respuestas"] = body.Respuestas
		o.Status = organism.StatusCompleted
		o.Touch()
		_ = s.Store.Save(o)
	}
	_ = s.Audit.Record("user", "evaluate", body.OrganismID, map[string]string{"puntaje": nivel})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (s *Server) handleInventario(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var body struct {
			EntityID   string                     `json:"entity_id"`
			EntityName string                     `json:"entity_name"`
			Conjuntos  []inventario.ConjuntoDatos `json:"conjuntos"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		o := inventario.CrearOrganismo(body.EntityID, body.EntityName)
		o.Data["conjuntos"] = body.Conjuntos
		o.Status = organism.StatusCompleted
		if err := s.Store.Save(o); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		_ = s.Audit.Record("user", "inventory", o.ID, nil)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(o)
		return
	}
	http.Error(w, "method not allowed", 405)
}

func (s *Server) handleExpediente(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	var body struct {
		EntityID   string   `json:"entity_id"`
		EntityName string   `json:"entity_name"`
		DiagRef    string   `json:"diagnostico_ref"`
		InvRef     string   `json:"inventario_ref"`
		Puntaje    float64  `json:"puntaje"`
		Nivel      string   `json:"nivel"`
		Hallazgos  []string `json:"hallazgos"`
		Brechas    []string `json:"brechas"`
		Pasos      []string `json:"pasos"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if !s.Policy.CanExportExpediente(body.DiagRef != "") {
		http.Error(w, "se requiere un diagnóstico previo", 403)
		return
	}
	exp := expediente.Generar(body.EntityID, body.EntityName, body.DiagRef, body.InvRef, body.Puntaje, body.Nivel, body.Hallazgos, body.Brechas, body.Pasos)
	o := expediente.CrearOrganismo(body.EntityID, body.EntityName)
	o.Data["expediente"] = exp
	o.Status = organism.StatusCompleted
	_ = s.Store.Save(o)
	_ = s.Audit.Record("user", "expediente", o.ID, nil)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exp)
}

func (s *Server) handleEstrategia(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}
	var body struct {
		EntityID   string `json:"entity_id"`
		EntityName string `json:"entity_name"`
		Nivel      string `json:"nivel"`
		HasDiag    bool   `json:"has_diagnostico"`
		HasInv     bool   `json:"has_inventario"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if !s.Policy.CanGenerateStrategy(body.HasDiag, body.HasInv) {
		http.Error(w, "se requiere diagnóstico e inventario de datos previos (principio: Primero los datos)", 403)
		return
	}
	rm := estrategia.GenerarBasico(body.EntityID, body.EntityName, body.Nivel)
	o := estrategia.CrearOrganismo(body.EntityID, body.EntityName)
	o.Data["roadmap"] = rm
	o.Status = organism.StatusCompleted
	_ = s.Store.Save(o)
	_ = s.Audit.Record("user", "estrategia", o.ID, nil)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rm)
}

func (s *Server) handleListOrganismos(w http.ResponseWriter, r *http.Request) {
	list, err := s.Store.List()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (s *Server) handleFrontend(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("web")).ServeHTTP(w, r)
}
