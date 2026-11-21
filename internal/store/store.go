package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/yecharlot/guantanamo-digital-ready/internal/organism"
)

// Store is a simple offline-first JSON file store.
type Store struct {
	dir string
	mu  sync.RWMutex
}

func New(dir string) (*Store, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	return &Store{dir: dir}, nil
}

func (s *Store) path(id string) string {
	return filepath.Join(s.dir, id+".json")
}

func (s *Store) Save(o *organism.Organism) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path(o.ID), data, 0o644)
}

func (s *Store) Load(id string) (*organism.Organism, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	data, err := os.ReadFile(s.path(id))
	if err != nil {
		return nil, err
	}
	var o organism.Organism
	if err := json.Unmarshal(data, &o); err != nil {
		return nil, err
	}
	return &o, nil
}

func (s *Store) List() ([]*organism.Organism, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	entries, err := os.ReadDir(s.dir)
	if err != nil {
		return nil, err
	}
	var list []*organism.Organism
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		data, err := os.ReadFile(filepath.Join(s.dir, e.Name()))
		if err != nil {
			continue
		}
		var o organism.Organism
		if json.Unmarshal(data, &o) == nil {
			list = append(list, &o)
		}
	}
	return list, nil
}
