# Guantánamo Digital Ready (GDR)

**Plataforma soberana de diagnóstico de madurez digital e inteligencia artificial**  
Alineada con la **Resolución 24/2026** del Ministerio de Comunicaciones de Cuba.

> “Primero los datos, la IA después.”

---

## Propósito

**Guantánamo Digital Ready** es una herramienta diseñada para ayudar a entidades municipales, provinciales y organismos de la Administración Pública (especialmente del territorio de Guantánamo) a cumplir las primeras etapas de la Resolución 24/2026:

- Diagnóstico de madurez digital (10 dimensiones)
- Inventario y gobernanza básica de datos
- Generación del **Expediente Digital de Arranque**
- Borrador de Estrategia de Transformación Digital e IA (Horizonte 2026-2027)

La plataforma prioriza:

- **Soberanía tecnológica**
- **Trazabilidad completa** (audit log + identidades)
- **Funcionamiento offline-first / sincronización diferida**
- **Bajo consumo de recursos**
- Arquitectura basada en **Digital Organisms** (inspirada en PrismaTec-Core)

---

## Arquitectura (visión)

```
┌─────────────────────────────────────────────────────────┐
│                 Interfaz Web (HTML+CSS+JS)              │
│        Colores claros · Animaciones ligeras · Charts    │
└──────────────────────────┬──────────────────────────────┘
                           │
┌──────────────────────────▼──────────────────────────────┐
│                    API HTTP (Go)                        │
│         Policy Engine · Audit Log · Offline store       │
└──────────────────────────┬──────────────────────────────┘
                           │
        ┌──────────────────┼──────────────────┐
        │                  │                  │
        ▼                  ▼                  ▼
┌───────────────┐  ┌───────────────┐  ┌───────────────┐
│ Org-Diagnóstico│  │Org-Inventario │  │ Org-Expediente│
│ Madurez digital│  │Datos + calidad│  │ + Estrategia  │
└───────────────┘  └───────────────┘  └───────────────┘
```

### Organismos del MVP

| Organismo               | Responsabilidad principal                                      |
|-------------------------|----------------------------------------------------------------|
| **Org-Diagnóstico**     | Cuestionario y puntuación de madurez digital (10 dimensiones)  |
| **Org-InventarioDatos** | Catálogo de conjuntos de datos, sensibilidad y calidad         |
| **Org-Expediente**      | Generación y versionado del Expediente Digital de Arranque     |
| **Org-Estrategia**      | Recomendaciones priorizadas y roadmap 2026-2027                |

---

## Características del MVP 1.0

- Runtime ligero en Go
- Cuatro organismos funcionales
- Interfaz web moderna, clara y con animaciones suaves
- Gráficos de madurez (radar / barras)
- Exportación a PDF del Expediente (versión inicial)
- Almacenamiento local (JSON) con posibilidad de sincronización posterior
- Policy engine + registro de auditoría desde el primer día
- Documentación completa en español

---

## Requisitos

- Go 1.22 o superior
- Navegador moderno

## Instalación rápida

```bash
git clone https://github.com/yecharlot/guantanamo-digital-ready.git
cd guantanamo-digital-ready

go mod tidy
go run ./cmd/gdr
```

Abre el navegador en: **http://localhost:8090**

---

## Estructura del repositorio

```
guantanamo-digital-ready/
├── cmd/gdr/                  # Punto de entrada
├── internal/
│   ├── organism/             # Digital Organisms
│   ├── policy/               # Motor de políticas
│   ├── audit/                # Registro de auditoría
│   ├── memory/               # Memoria content-addressed simple
│   ├── api/                  # Handlers HTTP
│   ├── pdf/                  # Generación de PDF
│   └── store/                # Persistencia local
├── web/                      # Frontend (HTML + CSS + JS)
├── docs/                     # Documentación ampliada
├── data/                     # Datos de ejemplo y almacenamiento
├── go.mod
├── LICENSE                   # Apache-2.0
└── README.md
```

---

## Alineación con la Resolución 24/2026

| Requisito de la Resolución              | Cómo lo cubre GDR                                      |
|-----------------------------------------|--------------------------------------------------------|
| Diagnóstico de madurez digital          | Org-Diagnóstico (10 dimensiones)                       |
| “Primero los datos…”                    | Org-InventarioDatos obligatorio antes de recomendaciones de IA |
| Expediente Digital de Arranque          | Org-Expediente genera documento versionado             |
| Estrategia 2026-2027                    | Org-Estrategia produce roadmap priorizado              |
| Trazabilidad y soberanía                | Audit log + identidades + almacenamiento local         |
| Condiciones de baja conectividad        | Diseño offline-first                                   |

---

## Licencia

Apache License 2.0  
Copyright © 2026 yecharlot / Guantánamo Digital Ready

---

## Estado del proyecto

**MVP 1.0 en construcción** — base funcional lista para uso local y demostraciones.

**Construido para Guantánamo. Pensado para Cuba. Orientado a la soberanía digital.**
