package audit

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Entry struct {
	ID        string            `json:"id"`
	Timestamp time.Time         `json:"timestamp"`
	Actor     string            `json:"actor"`
	Action    string            `json:"action"`
	Resource  string            `json:"resource"`
	Details   map[string]string `json:"details"`
}

type Log struct {
	dir string
	mu  sync.Mutex
}

func New(dir string) (*Log, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	return &Log{dir: dir}, nil
}

func (l *Log) Record(actor, action, resource string, details map[string]string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	e := Entry{
		ID:        time.Now().UTC().Format("20060102T150405.000"),
		Timestamp: time.Now().UTC(),
		Actor:     actor,
		Action:    action,
		Resource:  resource,
		Details:   details,
	}
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(l.dir, "audit.jsonl"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(append(data, '\n'))
	return err
}
