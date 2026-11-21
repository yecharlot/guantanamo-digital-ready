# Instalación y uso

## Requisitos

- Go 1.22+
- Navegador moderno

## Arranque

```bash
cd guantanamo-digital-ready
go mod tidy
go run ./cmd/gdr
```

Abrir: http://localhost:8090

Variables de entorno opcionales:

- `GDR_ADDR` – dirección de escucha (default `:8090`)
- `GDR_DATA_DIR` – directorio de datos (default `data/runtime`)

## Uso básico

1. Indicar el nombre de la entidad en la barra superior.
2. Ir a **Diagnóstico** y evaluar las 10 dimensiones.
3. Ir a **Inventario de Datos** y registrar los conjuntos de datos principales.
4. Generar el **Expediente Digital de Arranque**.
5. Generar el **Roadmap de Estrategia**.
