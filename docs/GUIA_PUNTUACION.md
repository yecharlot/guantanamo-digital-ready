# Guía de puntuación del diagnóstico de madurez digital

Cada una de las **10 dimensiones** se puntúa de **0 a 5**. Esta guía está implementada en:

- Backend: `internal/organism/diagnostico/criterios.go`
- API: `GET /api/diagnostico/criterios`
- Interfaz: formulario de diagnóstico (cada renglón muestra criterios expandibles)

## Escala común

| Nivel | Etiqueta | Idea general |
|------:|----------|--------------|
| 0 | Inexistente | No hay práctica ni evidencia |
| 1 | Inicial | Intentos aislados, sin formalizar |
| 2 | Básico | Formalización mínima / parcial |
| 3 | Intermedio | Práctica sistemática con responsables |
| 4 | Avanzado | Medición, integración y mejora |
| 5 | Óptimo | Madurez institucional sostenible |

## Cómo usar la interfaz

1. Lea la **pregunta guía** de la dimensión.
2. Abra **«Cómo puntuar esta dimensión»** y revise *por qué importa* y *cómo evaluar*.
3. Compare la realidad de la entidad con cada bloque **0–5** (*cuándo elegir*, *ejemplos*, *evidencia sugerida*).
4. Elija la puntuación; el sistema resalta el criterio del nivel elegido.
5. Anote la **evidencia** en el campo de notas (documentos, sistemas, hechos observados).

## Principio

Se puntúa el **estado actual demostrado**, no la intención ni el plan futuro. Ante duda entre dos niveles, elija el **inferior** hasta disponer de evidencia del superior.
