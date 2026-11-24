/* Guantánamo Digital Ready – Frontend MVP */

const state = {
  entityId: "gtm-mun-001",
  entityName: "Gobierno Municipal de Guantánamo",
  diagnosticoId: null,
  resultado: null,
  inventarioId: null,
  conjuntos: [],
  expediente: null,
  roadmap: null,
  respuestas: {}
};


/* Guías detalladas por dimensión y nivel (alineadas a internal/organism/diagnostico/criterios.go) */
const GUIAS = [
  {
    nombre: "Gobernanza y liderazgo digital",
    pregunta_guia: "¿Existe dirección formal, roles y decisiones sobre transformación digital en la entidad?",
    por_que_importa: "Sin liderazgo y gobernanza, los proyectos digitales se fragmentan, no se priorizan ni se sostiene el cambio.",
    como_evaluar: "Revise organigramas, actas, designaciones de CTO/CDTO o equivalentes, planes aprobados y evidencias de seguimiento por la dirección.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"No hay ninguna instancia, rol ni decisión formal sobre digitalización.", ejemplos:"Nadie es responsable; no aparece en planes institucionales.", evidencia_sugerida:"Ausencia de documentos o designaciones."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Hay interés o iniciativas aisladas de personas, sin mandato formal.", ejemplos:"Un jefe de informática impulsa mejoras por cuenta propia; menciones verbales en reuniones.", evidencia_sugerida:"Correos o notas informales; sin resolución/acta."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Existe un responsable designado o un grupo informal con algún plan escrito limitado.", ejemplos:"Designación por escrito del responsable TIC; plan anual breve sin presupuesto claro.", evidencia_sugerida:"Resolución de designación; borrador de plan."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Hay estructura de gobernanza activa (comité o equivalente) y plan alineado a objetivos de la entidad.", ejemplos:"Comité de transformación digital que se reúne; plan con metas y responsables.", evidencia_sugerida:"Actas de comité; plan aprobado; indicadores."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"La dirección usa indicadores digitales en la gestión y hay políticas de datos/TIC aplicadas.", ejemplos:"Tablero de seguimiento en consejos de dirección; política de datos firmada y comunicada.", evidencia_sugerida:"Informes periódicos a dirección; políticas vigentes."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Gobernanza madura, transversal, con mejora continua y rendición de cuentas pública o institucional.", ejemplos:"Estrategia digital institucional integrada al plan de desarrollo; auditoría de cumplimiento.", evidencia_sugerida:"Estrategia publicada; evaluación externa o autoevaluación formal."}
    ]
  },
  {
    nombre: "Infraestructura tecnológica",
    pregunta_guia: "¿Qué capacidad real tienen equipos, redes, servidores y continuidad de servicio?",
    por_que_importa: "La infraestructura es la base física/lógica; sin ella no hay servicios estables ni datos confiables.",
    como_evaluar: "Inventario de equipos, conectividad, centros de datos o hosting, respaldos y tiempos de indisponibilidad conocidos.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"Casi no hay equipos útiles o la entidad depende totalmente de medios ajenos no controlados.", ejemplos:"Puestos sin PC; sin red institucional.", evidencia_sugerida:"Inventario vacío o equipos fuera de servicio."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Equipos obsoletos o insuficientes; conectividad intermitente; sin estándares.", ejemplos:"PCs compartidos; Internet solo en una oficina.", evidencia_sugerida:"Listado informal de equipos."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Hay PCs y red local básica en áreas prioritarias; sin alta disponibilidad.", ejemplos:"LAN en sede principal; impresora compartida; antivirus puntual.", evidencia_sugerida:"Inventario básico; diagrama simple de red."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Infraestructura documentada, con respaldos regulares y soporte definido.", ejemplos:"Servidor de archivos; backups semanales; contrato o equipo de soporte.", evidencia_sugerida:"Política de backup; tickets de soporte."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Monitoreo, segmentación razonable, continuidad planificada y actualización controlada.", ejemplos:"Monitoreo de red; UPS; plan de continuidad; ventanas de mantenimiento.", evidencia_sugerida:"Informes de monitoreo; plan DR/BCP resumido."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Arquitectura resiliente, escalable, con métricas de disponibilidad y mejora continua.", ejemplos:"Alta disponibilidad en servicios críticos; métricas SLA internas cumplidas.", evidencia_sugerida:"SLA; reportes de disponibilidad; auditorías técnicas."}
    ]
  },
  {
    nombre: "Disponibilidad y calidad de datos",
    pregunta_guia: "¿Los datos institucionales están identificados, accesibles, actualizados y con calidad conocida?",
    por_que_importa: "La Resolución prioriza «primero los datos»; sin datos de calidad no hay IA ni decisiones confiables.",
    como_evaluar: "Catálogo de conjuntos de datos, dueños, frecuencia de actualización, errores conocidos y medios de acceso.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"No se sabe qué datos existen ni dónde están.", ejemplos:"Información solo en papel disperso o en personas.", evidencia_sugerida:"Sin inventario de datos."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Datos en archivos aislados (Excel, USB) sin dueño claro ni control de versiones.", ejemplos:"Cada área tiene sus hojas de cálculo distintas.", evidencia_sugerida:"Muestras de archivos sin estándar."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Algunos conjuntos prioritarios identificados; calidad desigual; poco documentados.", ejemplos:"Padrones parciales; bases Access locales.", evidencia_sugerida:"Lista inicial de conjuntos de datos."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Catálogo básico con responsables; reglas mínimas de actualización y validación.", ejemplos:"Diccionario de datos parcial; controles de duplicados en sistemas clave.", evidencia_sugerida:"Catálogo; responsables asignados."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Calidad medida (completitud, exactitud); datos maestros definidos; acceso controlado.", ejemplos:"Indicadores de calidad; proceso de depuración periódico.", evidencia_sugerida:"Informes de calidad; políticas de datos maestros."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Gobernanza de datos madura, trazabilidad del ciclo de vida y datos listos para analítica/IA.", ejemplos:"Lineaje de datos; datasets certificados para uso analítico.", evidencia_sugerida:"Marco de gobernanza de datos; certificaciones internas."}
    ]
  },
  {
    nombre: "Seguridad informática y ciberseguridad",
    pregunta_guia: "¿Cómo se protegen sistemas, datos y usuarios ante incidentes y amenazas?",
    por_que_importa: "La digitalización sin seguridad expone datos de ciudadanos y la continuidad del servicio público.",
    como_evaluar: "Controles de acceso, respaldos, actualizaciones, sensibilización, gestión de incidentes y cumplimiento básico.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"Sin controles básicos; contraseñas compartidas o ausentes; sin conciencia del riesgo.", ejemplos:"Equipos sin usuario/contraseña; USB libre sin control.", evidencia_sugerida:"Hallazgos evidentes de ausencia de controles."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Medidas ad hoc (un antivirus, alguna contraseña) sin política.", ejemplos:"Antivirus en algunos PCs; WiFi con clave compartida antigua.", evidencia_sugerida:"Instalaciones puntuales."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Controles mínimos documentados: cuentas por usuario, antivirus institucional, respaldos ocasionales.", ejemplos:"Política breve de contraseñas; backups mensuales.", evidencia_sugerida:"Documento de normas básicas."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Gestión de accesos, actualizaciones regulares, canales de reporte de incidentes.", ejemplos:"Altas/bajas de usuarios; parches; registro de incidentes.", evidencia_sugerida:"Procedimiento de incidentes; registros."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Enfoque preventivo: clasificación de información, pruebas, formación periódica.", ejemplos:"Clasificación de datos; simulacros; capacitación anual.", evidencia_sugerida:"Plan de seguridad; evidencias de formación."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Sistema de gestión de seguridad alineado a buenas prácticas, con mejora continua y auditoría.", ejemplos:"SGSI o equivalente adaptado; auditorías y tratamiento de riesgos.", evidencia_sugerida:"Informes de auditoría; matriz de riesgos actualizada."}
    ]
  },
  {
    nombre: "Procesos de trabajo digitalizados",
    pregunta_guia: "¿Qué proporción de los procesos prioritarios se ejecuta y registra en medios digitales?",
    por_que_importa: "Digitalizar procesos reduce tiempos, errores y habilita datos operativos reales.",
    como_evaluar: "Mapee procesos clave (trámites, control interno, logística) y vea si hay sistema, plantillas o solo papel.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"Casi todo es papel o oral; no hay flujos digitales.", ejemplos:"Expedientes solo físicos; colas sin turno digital.", evidencia_sugerida:"Observación de operación 100% manual."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Uso esporádico de ofimática; procesos siguen siendo manuales.", ejemplos:"Se imprime todo lo que se escribe en Word.", evidencia_sugerida:"Plantillas sueltas sin flujo."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Algunos procesos parciales en sistemas o formularios digitales, con retrabajo a papel.", ejemplos:"Registro digital que luego se imprime y firma en papel.", evidencia_sugerida:"Capturas de pantallas de sistemas parciales."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Procesos prioritarios mayormente digitales, con reglas claras de operación.", ejemplos:"Trámite interno en sistema de workflow; menos papel.", evidencia_sugerida:"Mapa de procesos; sistema en producción."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Automatizaciones, integraciones entre áreas y medición de tiempos/calidad del proceso.", ejemplos:"Alertas automáticas; tableros de tiempos de trámite.", evidencia_sugerida:"KPI de procesos; integraciones documentadas."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Procesos optimizados de punta a punta, medidos y en mejora continua basada en datos.", ejemplos:"Rediseño continuo; eliminación de pasos sin valor.", evidencia_sugerida:"Ciclo de mejora documentado; resultados medibles."}
    ]
  },
  {
    nombre: "Servicios al ciudadano",
    pregunta_guia: "¿Cómo interactúa el ciudadano/usuario con la entidad por canales digitales?",
    por_que_importa: "La transformación digital pública se valida en la experiencia real del ciudadano.",
    como_evaluar: "Presencia web/app, trámites en línea, citas, información actualizada, accesibilidad y soporte.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"Sin canal digital útil; solo presencial o teléfono informal.", ejemplos:"Sin sitio web o sitio abandonado años.", evidencia_sugerida:"Verificación de ausencia de canales."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Información estática mínima en web o redes, sin trámites.", ejemplos:"Página con dirección y teléfono desactualizados.", evidencia_sugerida:"Captura del sitio."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Información actualizada y algún formulario o descarga de modelos.", ejemplos:"PDF de requisitos; formulario de contacto.", evidencia_sugerida:"Formularios publicados."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Trámites o solicitudes digitales con seguimiento básico.", ejemplos:"Cita en línea; estado de solicitud por código.", evidencia_sugerida:"Sistema de trámites en uso."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Servicios transaccionales integrados, con identidad del usuario y métricas de uso.", ejemplos:"Pagos o validaciones en línea según marco legal; panel de uso.", evidencia_sugerida:"Estadísticas de uso; integración con backends."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Experiencia omnicanal, accesible, con mejora continua basada en satisfacción y datos.", ejemplos:"Múltiples canales coherentes; encuestas; ajustes frecuentes.", evidencia_sugerida:"Medición de satisfacción; roadmap de servicio."}
    ]
  },
  {
    nombre: "Capacidades del personal",
    pregunta_guia: "¿El personal tiene competencias digitales y de datos acordes a su rol?",
    por_que_importa: "La tecnología sin personas capacitadas no se adopta ni se sostiene.",
    como_evaluar: "Planes de formación, nivel de uso real de herramientas, perfiles TIC y cultura de aprendizaje.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"Personal sin habilidades digitales básicas para el puesto.", ejemplos:"Dependencia total de una sola persona para cualquier tarea en PC.", evidencia_sugerida:"Observación / entrevistas."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Uso básico autodidacta; sin plan de formación.", ejemplos:"Algunos usan WhatsApp o Word a nivel elemental.", evidencia_sugerida:"Sin plan de capacitación."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Capacitaciones puntuales; usuarios operan sistemas prioritarios con apoyo.", ejemplos:"Curso ocasional; manuales internos breves.", evidencia_sugerida:"Registros de asistencia a cursos."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Plan de formación digital por roles; competencias definidas.", ejemplos:"Rutas de aprendizaje; evaluación básica de competencias.", evidencia_sugerida:"Plan anual de capacitación."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Competencias de datos/TIC extendidas; multiplicadores internos; actualización continua.", ejemplos:"Referentes digitales por área; comunidades de práctica.", evidencia_sugerida:"Evidencias de formación continua."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Cultura de aprendizaje digital consolidada; atracción y retención de talento especializado.", ejemplos:"Perfiles de datos/IA o avanzados TIC; carrera profesional.", evidencia_sugerida:"Política de talento digital; indicadores de clima/retención."}
    ]
  },
  {
    nombre: "Marco normativo y cumplimiento",
    pregunta_guia: "¿La entidad aplica y documenta el marco legal y normativo digital y de datos aplicable?",
    por_que_importa: "La legitimidad y la protección de derechos dependen del cumplimiento normativo.",
    como_evaluar: "Políticas internas, registros exigidos, protección de datos personales, alineación a resoluciones ministeriales.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"Desconocimiento o incumplimiento evidente de normas digitales/datos aplicables.", ejemplos:"Sin políticas; datos personales sin control.", evidencia_sugerida:"Hallazgos de incumplimiento."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Conocimiento parcial; aplicación inconsistente.", ejemplos:"Se conoce la norma pero no hay procedimientos.", evidencia_sugerida:"Entrevistas; ausencia de procedimientos."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Algunas políticas o procedimientos escritos alineados a normas prioritarias.", ejemplos:"Política interna breve de uso de TIC o de datos.", evidencia_sugerida:"Documento aprobado."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Cumplimiento sistemático de obligaciones principales con responsables claros.", ejemplos:"Registros, consentimientos o controles según marco aplicable.", evidencia_sugerida:"Evidencias de cumplimiento rutinario."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Seguimiento proactivo de cambios normativos y ajustes oportunos.", ejemplos:"Revisión periódica del marco; checklist de cumplimiento.", evidencia_sugerida:"Actas de revisión normativa."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Cultura de cumplimiento, auditoría interna y mejora ante hallazgos.", ejemplos:"Auditorías; planes correctivos cerrados a tiempo.", evidencia_sugerida:"Informes de auditoría y seguimiento."}
    ]
  },
  {
    nombre: "Interoperabilidad",
    pregunta_guia: "¿Los sistemas y áreas intercambian datos de forma acordada, segura y reutilizable?",
    por_que_importa: "Sin interoperabilidad se duplican esfuerzos y se fragmenta la visión del ciudadano y de la gestión.",
    como_evaluar: "Interfaces, estándares, convenios de intercambio, silos conocidos y reutilización de datos entre sistemas.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"Sistemas y áreas totalmente aislados; intercambio solo por papel o USB informal.", ejemplos:"Cada sistema es un silo; reescritura manual de datos.", evidencia_sugerida:"Arquitectura de sistemas sin enlaces."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Intercambios esporádicos manuales o archivos enviados por correo.", ejemplos:"Exportar CSV y reimportar a mano.", evidencia_sugerida:"Ejemplos de archivos intercambiados."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Algunas interfaces puntuales o bases compartidas limitadas entre sistemas propios.", ejemplos:"Dos sistemas de la entidad comparten una tabla o carpeta controlada.", evidencia_sugerida:"Descripción de la interfaz puntual."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Estándares internos de intercambio; convenios con al menos un actor externo o sistema nacional/sectorial.", ejemplos:"API o formato acordado; convenio de intercambio.", evidencia_sugerida:"Documento de interfaz; convenio."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Interoperabilidad planificada en la arquitectura; reutilización sistemática de datos.", ejemplos:"Catálogo de servicios de datos; menos duplicación.", evidencia_sugerida:"Mapa de integraciones."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Ecosistema interoperable medido, seguro y alineado a marcos nacionales/sectoriales de intercambio.", ejemplos:"Participación estable en plataformas de interoperabilidad; monitoreo de intercambios.", evidencia_sugerida:"Métricas de intercambio; certificaciones o adhesiones formales."}
    ]
  },
  {
    nombre: "Preparación para IA (Pirámide Funcional)",
    pregunta_guia: "¿Están sentadas las bases (datos, procesos, ética, casos de uso) antes de adoptar IA?",
    por_que_importa: "La Resolución insiste en no saltar a IA sin datos y capacidades previas; evita proyectos fallidos.",
    como_evaluar: "Madurez de datos y procesos previos, casos de uso priorizados, riesgos éticos y capacidad de evaluar proveedores/modelos.",
    niveles: [
      {nivel:0, etiqueta:"Inexistente", cuando_elegir:"No hay bases de datos/procesos; se habla de IA solo como moda, sin casos reales.", ejemplos:"Propuestas de IA sin inventario de datos ni problema claro.", evidencia_sugerida:"Ausencia de casos de uso documentados."},
      {nivel:1, etiqueta:"Inicial", cuando_elegir:"Interés exploratorio; experimentos personales sin gobierno ni datos preparados.", ejemplos:"Pruebas con herramientas genéricas sin datos institucionales.", evidencia_sugerida:"Demos aisladas."},
      {nivel:2, etiqueta:"Básico", cuando_elegir:"Se reconocen precondiciones (datos, calidad, procesos) y hay al menos un caso de uso hipotético priorizado.", ejemplos:"Documento que pospone IA hasta mejorar datos; lista de ideas de casos de uso.", evidencia_sugerida:"Nota de priorización."},
      {nivel:3, etiqueta:"Intermedio", cuando_elegir:"Casos de uso seleccionados con datos disponibles y criterios de éxito; riesgos identificados.", ejemplos:"Pilot piloto definido (ej. clasificación de documentos) con dataset y métrica.", evidencia_sugerida:"Ficha de caso de uso; dataset identificado."},
      {nivel:4, etiqueta:"Avanzado", cuando_elegir:"Piloto(s) evaluados, con gobernanza, supervisión humana y medición de resultados.", ejemplos:"Piloto en producción limitada; revisión ética/operativa.", evidencia_sugerida:"Informe de resultados del piloto."},
      {nivel:5, etiqueta:"Óptimo", cuando_elegir:"Capacidad institucional para incorporar IA de forma responsable, medible y sostenible.", ejemplos:"Portafolio de casos; políticas de uso responsable; mejora continua.", evidencia_sugerida:"Política de IA responsable; portafolio activo."}
    ]
  }
];

const dimensiones = GUIAS.map(g => g.nombre);

function renderDimensionForm() {
  const form = document.getElementById("dimensiones-form");
  form.innerHTML = GUIAS.map((g, i) => {
    const nivelesHTML = g.niveles.map(n => `
      <div class="nivel-criterio" data-nivel="${n.nivel}">
        <div class="nivel-criterio-head">
          <span class="nivel-badge n${n.nivel}">${n.nivel}</span>
          <strong>${n.etiqueta}</strong>
        </div>
        <p><span class="crit-label">Cuándo elegir:</span> ${n.cuando_elegir}</p>
        <p><span class="crit-label">Ejemplos:</span> ${n.ejemplos}</p>
        <p><span class="crit-label">Evidencia sugerida:</span> ${n.evidencia_sugerida}</p>
      </div>
    `).join("");
    return `
    <article class="dim-card" data-dim-index="${i}">
      <header class="dim-card-head">
        <span class="dim-num">${i + 1}</span>
        <div>
          <h4>${g.nombre}</h4>
          <p class="dim-pregunta">${g.pregunta_guia}</p>
        </div>
      </header>
      <details class="dim-details" open>
        <summary>Cómo puntuar esta dimensión (guía para especialistas)</summary>
        <div class="dim-guide">
          <p><span class="crit-label">Por qué importa:</span> ${g.por_que_importa}</p>
          <p><span class="crit-label">Cómo evaluar:</span> ${g.como_evaluar}</p>
          <div class="niveles-grid">${nivelesHTML}</div>
        </div>
      </details>
      <div class="dim-score-row">
        <label for="sel-dim-${i}">Puntuación elegida</label>
        <select id="sel-dim-${i}" data-dim="${g.nombre}" class="dim-select">
          ${g.niveles.map(n => `<option value="${n.nivel}" ${n.nivel===2?"selected":""}>${n.nivel} – ${n.etiqueta}</option>`).join("")}
        </select>
        <p class="dim-score-hint" id="hint-dim-${i}"></p>
      </div>
      <label class="dim-evidencia-label" for="ev-dim-${i}">Evidencia / notas del evaluador</label>
      <textarea id="ev-dim-${i}" class="dim-evidencia" rows="2" placeholder="Cite documentos, sistemas o hechos observados que justifican la puntuación…"></textarea>
    </article>`;
  }).join("");

  form.querySelectorAll(".dim-select").forEach(sel => {
    const updateHint = () => {
      const card = sel.closest(".dim-card");
      const idx = parseInt(card.dataset.dimIndex, 10);
      const nivel = parseInt(sel.value, 10);
      const c = GUIAS[idx].niveles.find(n => n.nivel === nivel);
      const hint = card.querySelector(".dim-score-hint");
      if (c && hint) {
        hint.innerHTML = `<strong>Por qué este nivel:</strong> ${c.cuando_elegir}`;
      }
      card.querySelectorAll(".nivel-criterio").forEach(el => {
        el.classList.toggle("is-selected", parseInt(el.dataset.nivel, 10) === nivel);
      });
    };
    sel.addEventListener("change", updateHint);
    updateHint();
  });
}

renderDimensionForm();



// Evaluate diagnosis
document.getElementById("btn-evaluar").addEventListener("click", async () => {
  const respuestas = [];
  document.querySelectorAll("#dimensiones-form select.dim-select").forEach(sel => {
    const card = sel.closest(".dim-card");
    const ev = card ? card.querySelector(".dim-evidencia") : null;
    respuestas.push({
      dimension: sel.dataset.dim,
      nivel: parseInt(sel.value, 10),
      evidencia: ev ? ev.value.trim() : "",
      notas: ""
    });
  });

  // Create organism first
  let orgId = state.diagnosticoId;
  if (!orgId) {
    const res = await fetch("/api/diagnostico", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ entity_id: state.entityId, entity_name: state.entityName })
    });
    const org = await res.json();
    orgId = org.id;
    state.diagnosticoId = orgId;
  }

  const res = await fetch("/api/diagnostico/evaluar", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      organism_id: orgId,
      entity_id: state.entityId,
      entity_name: state.entityName,
      respuestas
    })
  });
  const result = await res.json();
  state.resultado = result;

  const box = document.getElementById("resultado-diagnostico");
  box.classList.remove("hidden");
  box.innerHTML = `
    <h4>Resultado del diagnóstico</h4>
    <p><strong>Puntaje:</strong> ${result.puntaje_total.toFixed(1)} / 100</p>
    <p><strong>Nivel global:</strong> ${result.nivel_global}</p>
    <p style="margin-top:0.5rem;color:#64748b;font-size:0.9rem;">Fecha: ${new Date(result.fecha).toLocaleString()}</p>
  `;

  updateDashboard();
  updateRadar(respuestas);
});

// Inventory
document.getElementById("btn-add-dato").addEventListener("click", () => {
  const nombre = document.getElementById("dato-nombre").value.trim();
  if (!nombre) return alert("Escriba el nombre del conjunto de datos");
  const sensibilidad = document.getElementById("dato-sensibilidad").value;
  const calidad = document.getElementById("dato-calidad").value;
  state.conjuntos.push({
    id: "d-" + Date.now(),
    nombre,
    descripcion: "",
    propietario: state.entityName,
    sensibilidad,
    calidad,
    formato: "por definir",
    ubicacion: "local",
    actualizado: new Date().toISOString().slice(0, 10)
  });
  document.getElementById("dato-nombre").value = "";
  renderListaDatos();
});

function renderListaDatos() {
  const ul = document.getElementById("lista-datos");
  ul.innerHTML = state.conjuntos.map(c => `
    <li>
      <span><strong>${c.nombre}</strong> · ${c.sensibilidad} · calidad ${c.calidad}</span>
    </li>
  `).join("");
  document.getElementById("dash-datos").textContent = state.conjuntos.length;
}

document.getElementById("btn-guardar-inventario").addEventListener("click", async () => {
  if (state.conjuntos.length === 0) return alert("Añada al menos un conjunto de datos");
  const res = await fetch("/api/inventario", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      entity_id: state.entityId,
      entity_name: state.entityName,
      conjuntos: state.conjuntos
    })
  });
  const org = await res.json();
  state.inventarioId = org.id;
  alert("Inventario guardado correctamente");
  updateDashboard();
});

// Expediente
document.getElementById("btn-generar-expediente").addEventListener("click", async () => {
  if (!state.resultado) return alert("Primero realice el diagnóstico de madurez");
  const hallazgos = [
    `Nivel de madurez global: ${state.resultado.nivel_global}`,
    `Puntaje obtenido: ${state.resultado.puntaje_total.toFixed(1)} puntos`
  ];
  const brechas = state.resultado.respuestas
    .filter(r => r.nivel <= 2)
    .map(r => r.dimension);
  const pasos = [
    "Completar inventario de datos institucionales",
    "Designar responsables de Transformación Digital",
    "Elaborar política de datos de la entidad"
  ];

  const res = await fetch("/api/expediente", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      entity_id: state.entityId,
      entity_name: state.entityName,
      diagnostico_ref: state.diagnosticoId || "",
      inventario_ref: state.inventarioId || "",
      puntaje: state.resultado.puntaje_total,
      nivel: state.resultado.nivel_global,
      hallazgos,
      brechas: brechas.length ? brechas : ["Sin brechas críticas identificadas"],
      pasos
    })
  });
  if (!res.ok) {
    const t = await res.text();
    return alert(t);
  }
  const exp = await res.json();
  state.expediente = exp;
  const box = document.getElementById("expediente-preview");
  box.classList.remove("hidden");
  box.innerHTML = `
    <h4>Expediente Digital de Arranque v${exp.version}</h4>
    <p><strong>Entidad:</strong> ${exp.entity_name}</p>
    <p><strong>Fecha:</strong> ${new Date(exp.fecha_generacion).toLocaleString()}</p>
    <p><strong>Madurez:</strong> ${exp.puntaje_madurez.toFixed(1)} – ${exp.nivel_madurez}</p>
    <p style="margin-top:0.75rem"><strong>Hallazgos clave:</strong></p>
    <ul>${exp.hallazgos_clave.map(h => `<li>${h}</li>`).join("")}</ul>
    <p style="margin-top:0.75rem"><strong>Brechas críticas:</strong></p>
    <ul>${exp.brechas_criticas.map(b => `<li>${b}</li>`).join("")}</ul>
    <p style="margin-top:0.75rem"><strong>Próximos pasos:</strong></p>
    <ul>${exp.proximos_pasos.map(p => `<li>${p}</li>`).join("")}</ul>
  `;
  document.getElementById("dash-exp").textContent = "Generado";
  updateDashboard();
});

// Estrategia
document.getElementById("btn-generar-estrategia").addEventListener("click", async () => {
  const res = await fetch("/api/estrategia", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      entity_id: state.entityId,
      entity_name: state.entityName,
      nivel: state.resultado ? state.resultado.nivel_global : "Básico",
      has_diagnostico: !!state.resultado,
      has_inventario: state.conjuntos.length > 0
    })
  });
  if (!res.ok) {
    const t = await res.text();
    return alert(t);
  }
  const rm = await res.json();
  state.roadmap = rm;
  const box = document.getElementById("estrategia-preview");
  box.classList.remove("hidden");
  box.innerHTML = `
    <h4>Roadmap 2026-2027</h4>
    <p style="margin-bottom:1rem;color:#64748b">${rm.resumen}</p>
    ${rm.recomendaciones.map(r => `
      <div class="rec-item ${r.prioridad}">
        <h4>${r.titulo}</h4>
        <p>${r.descripcion}</p>
        <p style="margin-top:0.35rem;font-size:0.8rem">
          <strong>Prioridad:</strong> ${r.prioridad} · <strong>Horizonte:</strong> ${r.horizonte}
        </p>
      </div>
    `).join("")}
  `;
  document.getElementById("dash-recs").textContent = rm.recomendaciones.length;
  updateDashboard();
});

function updateDashboard() {
  if (state.resultado) {
    document.getElementById("dash-puntaje").textContent = state.resultado.puntaje_total.toFixed(0);
    document.getElementById("dash-nivel").textContent = state.resultado.nivel_global;
  }
  document.getElementById("dash-datos").textContent = state.conjuntos.length;
}

let radarChart = null;
function updateRadar(respuestas) {
  const ctx = document.getElementById("radarChart");
  const labels = respuestas.map(r => r.dimension.split(" ")[0] + "…");
  const data = respuestas.map(r => r.nivel);
  if (radarChart) radarChart.destroy();
  radarChart = new Chart(ctx, {
    type: "radar",
    data: {
      labels,
      datasets: [{
        label: "Nivel de madurez (0-5)",
        data,
        fill: true,
        backgroundColor: "rgba(14, 165, 233, 0.2)",
        borderColor: "#0ea5e9",
        pointBackgroundColor: "#6366f1",
        pointBorderColor: "#fff",
        pointHoverBackgroundColor: "#fff",
        pointHoverBorderColor: "#6366f1"
      }]
    },
    options: {
      scales: {
        r: {
          min: 0,
          max: 5,
          ticks: { stepSize: 1 }
        }
      },
      plugins: {
        legend: { display: false }
      }
    }
  });
}

// Initial empty radar
updateRadar(dimensiones.map(d => ({ dimension: d, nivel: 0 })));
