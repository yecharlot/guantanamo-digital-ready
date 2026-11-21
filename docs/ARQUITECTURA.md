# Arquitectura de Guantánamo Digital Ready

## Principios

1. **Soberanía** – Los datos de la entidad permanecen bajo control local.
2. **Offline-first** – El sistema funciona sin conexión permanente.
3. **Trazabilidad** – Toda acción relevante queda registrada en el audit log.
4. **Primero los datos** – No se generan recomendaciones de IA sin inventario de datos.
5. **Organismos digitales** – Cada capacidad principal se modela como un Digital Organism (inspirado en PrismaTec-Core).

## Capas

- **Frontend** (`web/`): HTML + CSS + JS. Interfaz clara, animaciones ligeras y gráficos (Chart.js).
- **API** (`internal/api`): HTTP handlers en Go.
- **Organismos** (`internal/organism/...`): Lógica de dominio de cada capacidad.
- **Policy** (`internal/policy`): Reglas de negocio (ej. exigir diagnóstico + inventario antes de estrategia).
- **Audit** (`internal/audit`): Registro append-only de acciones.
- **Store** (`internal/store`): Persistencia JSON local (preparada para sincronización diferida).

## Flujo típico

1. Crear / evaluar diagnóstico de madurez.
2. Registrar inventario de datos.
3. Generar Expediente Digital de Arranque.
4. Generar Roadmap de estrategia 2026-2027.
