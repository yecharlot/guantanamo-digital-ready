package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/yecharlot/guantanamo-digital-ready/internal/api"
	"github.com/yecharlot/guantanamo-digital-ready/internal/audit"
	"github.com/yecharlot/guantanamo-digital-ready/internal/policy"
	"github.com/yecharlot/guantanamo-digital-ready/internal/store"
)

func main() {
	dataDir := os.Getenv("GDR_DATA_DIR")
	if dataDir == "" {
		dataDir = filepath.Join("data", "runtime")
	}

	st, err := store.New(dataDir)
	if err != nil {
		log.Fatal(err)
	}
	au, err := audit.New(filepath.Join(dataDir, "audit"))
	if err != nil {
		log.Fatal(err)
	}
	pol := policy.New()

	srv := api.New(st, au, pol)

	addr := ":8090"
	if v := os.Getenv("GDR_ADDR"); v != "" {
		addr = v
	}

	fmt.Println("╔══════════════════════════════════════════════════════╗")
	fmt.Println("║     Guantánamo Digital Ready (GDR) – MVP 1.0         ║")
	fmt.Println("║     Resolución 24/2026 · Soberanía Digital           ║")
	fmt.Println("╚══════════════════════════════════════════════════════╝")
	fmt.Printf("→ Servidor escuchando en http://localhost%s\n", addr)
	fmt.Println("→ Almacenamiento local:", dataDir)
	fmt.Println("→ Modo: offline-first")
	fmt.Println()

	log.Fatal(http.ListenAndServe(addr, srv.Mux))
}
