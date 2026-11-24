package diagnostico

// CriterioNivel describe cuándo elegir cada puntuación (0–5).
type CriterioNivel struct {
	Nivel       Nivel  `json:"nivel"`
	Etiqueta    string `json:"etiqueta"`
	CuandoElegir string `json:"cuando_elegir"`
	Ejemplos    string `json:"ejemplos"`
	EvidenciaSugerida string `json:"evidencia_sugerida"`
}

// GuiaDimension orienta al especialista sobre una dimensión completa.
type GuiaDimension struct {
	Nombre      string         `json:"nombre"`
	PreguntaGuia string        `json:"pregunta_guia"`
	PorQueImporta string       `json:"por_que_importa"`
	ComoEvaluar string         `json:"como_evaluar"`
	Niveles     []CriterioNivel `json:"niveles"`
}

// Guias contiene la orientación detallada por dimensión (orden = Dimensiones).
var Guias = []GuiaDimension{
	{
		Nombre: "Gobernanza y liderazgo digital",
		PreguntaGuia: "¿Existe dirección formal, roles y decisiones sobre transformación digital en la entidad?",
		PorQueImporta: "Sin liderazgo y gobernanza, los proyectos digitales se fragmentan, no se priorizan ni se sostiene el cambio.",
		ComoEvaluar: "Revise organigramas, actas, designaciones de CTO/CDTO o equivalentes, planes aprobados y evidencias de seguimiento por la dirección.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "No hay ninguna instancia, rol ni decisión formal sobre digitalización.", "Nadie es responsable; no aparece en planes institucionales.", "Ausencia de documentos o designaciones."},
			{1, "Inicial", "Hay interés o iniciativas aisladas de personas, sin mandato formal.", "Un jefe de informática impulsa mejoras por cuenta propia; menciones verbales en reuniones.", "Correos o notas informales; sin resolución/acta."},
			{2, "Básico", "Existe un responsable designado o un grupo informal con algún plan escrito limitado.", "Designación por escrito del responsable TIC; plan anual breve sin presupuesto claro.", "Resolución de designación; borrador de plan."},
			{3, "Intermedio", "Hay estructura de gobernanza activa (comité o equivalente) y plan alineado a objetivos de la entidad.", "Comité de transformación digital que se reúne; plan con metas y responsables.", "Actas de comité; plan aprobado; indicadores."},
			{4, "Avanzado", "La dirección usa indicadores digitales en la gestión y hay políticas de datos/TIC aplicadas.", "Tablero de seguimiento en consejos de dirección; política de datos firmada y comunicada.", "Informes periódicos a dirección; políticas vigentes."},
			{5, "Óptimo", "Gobernanza madura, transversal, con mejora continua y rendición de cuentas pública o institucional.", "Estrategia digital institucional integrada al plan de desarrollo; auditoría de cumplimiento.", "Estrategia publicada; evaluación externa o autoevaluación formal."},
		},
	},
	{
		Nombre: "Infraestructura tecnológica",
		PreguntaGuia: "¿Qué capacidad real tienen equipos, redes, servidores y continuidad de servicio?",
		PorQueImporta: "La infraestructura es la base física/lógica; sin ella no hay servicios estables ni datos confiables.",
		ComoEvaluar: "Inventario de equipos, conectividad, centros de datos o hosting, respaldos y tiempos de indisponibilidad conocidos.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "Casi no hay equipos útiles o la entidad depende totalmente de medios ajenos no controlados.", "Puestos sin PC; sin red institucional.", "Inventario vacío o equipos fuera de servicio."},
			{1, "Inicial", "Equipos obsoletos o insuficientes; conectividad intermitente; sin estándares.", "PCs compartidos; Internet solo en una oficina.", "Listado informal de equipos."},
			{2, "Básico", "Hay PCs y red local básica en áreas prioritarias; sin alta disponibilidad.", "LAN en sede principal; impresora compartida; antivirus puntual.", "Inventario básico; diagrama simple de red."},
			{3, "Intermedio", "Infraestructura documentada, con respaldos regulares y soporte definido.", "Servidor de archivos; backups semanales; contrato o equipo de soporte.", "Política de backup; tickets de soporte."},
			{4, "Avanzado", "Monitoreo, segmentación razonable, continuidad planificada y actualización controlada.", "Monitoreo de red; UPS; plan de continuidad; ventanas de mantenimiento.", "Informes de monitoreo; plan DR/BCP resumido."},
			{5, "Óptimo", "Arquitectura resiliente, escalable, con métricas de disponibilidad y mejora continua.", "Alta disponibilidad en servicios críticos; métricas SLA internas cumplidas.", "SLA; reportes de disponibilidad; auditorías técnicas."},
		},
	},
	{
		Nombre: "Disponibilidad y calidad de datos",
		PreguntaGuia: "¿Los datos institucionales están identificados, accesibles, actualizados y con calidad conocida?",
		PorQueImporta: "La Resolución prioriza «primero los datos»; sin datos de calidad no hay IA ni decisiones confiables.",
		ComoEvaluar: "Catálogo de conjuntos de datos, dueños, frecuencia de actualización, errores conocidos y medios de acceso.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "No se sabe qué datos existen ni dónde están.", "Información solo en papel disperso o en personas.", "Sin inventario de datos."},
			{1, "Inicial", "Datos en archivos aislados (Excel, USB) sin dueño claro ni control de versiones.", "Cada área tiene sus hojas de cálculo distintas.", "Muestras de archivos sin estándar."},
			{2, "Básico", "Algunos conjuntos prioritarios identificados; calidad desigual; poco documentados.", "Padrones parciales; bases Access locales.", "Lista inicial de conjuntos de datos."},
			{3, "Intermedio", "Catálogo básico con responsables; reglas mínimas de actualización y validación.", "Diccionario de datos parcial; controles de duplicados en sistemas clave.", "Catálogo; responsables asignados."},
			{4, "Avanzado", "Calidad medida (completitud, exactitud); datos maestros definidos; acceso controlado.", "Indicadores de calidad; proceso de depuración periódico.", "Informes de calidad; políticas de datos maestros."},
			{5, "Óptimo", "Gobernanza de datos madura, trazabilidad del ciclo de vida y datos listos para analítica/IA.", "Lineaje de datos; datasets certificados para uso analítico.", "Marco de gobernanza de datos; certificaciones internas."},
		},
	},
	{
		Nombre: "Seguridad informática y ciberseguridad",
		PreguntaGuia: "¿Cómo se protegen sistemas, datos y usuarios ante incidentes y amenazas?",
		PorQueImporta: "La digitalización sin seguridad expone datos de ciudadanos y la continuidad del servicio público.",
		ComoEvaluar: "Controles de acceso, respaldos, actualizaciones, sensibilización, gestión de incidentes y cumplimiento básico.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "Sin controles básicos; contraseñas compartidas o ausentes; sin conciencia del riesgo.", "Equipos sin usuario/contraseña; USB libre sin control.", "Hallazgos evidentes de ausencia de controles."},
			{1, "Inicial", "Medidas ad hoc (un antivirus, alguna contraseña) sin política.", "Antivirus en algunos PCs; WiFi con clave compartida antigua.", "Instalaciones puntuales."},
			{2, "Básico", "Controles mínimos documentados: cuentas por usuario, antivirus institucional, respaldos ocasionales.", "Política breve de contraseñas; backups mensuales.", "Documento de normas básicas."},
			{3, "Intermedio", "Gestión de accesos, actualizaciones regulares, canales de reporte de incidentes.", "Altas/bajas de usuarios; parches; registro de incidentes.", "Procedimiento de incidentes; registros."},
			{4, "Avanzado", "Enfoque preventivo: clasificación de información, pruebas, formación periódica.", "Clasificación de datos; simulacros; capacitación anual.", "Plan de seguridad; evidencias de formación."},
			{5, "Óptimo", "Sistema de gestión de seguridad alineado a buenas prácticas, con mejora continua y auditoría.", "SGSI o equivalente adaptado; auditorías y tratamiento de riesgos.", "Informes de auditoría; matriz de riesgos actualizada."},
		},
	},
	{
		Nombre: "Procesos de trabajo digitalizados",
		PreguntaGuia: "¿Qué proporción de los procesos prioritarios se ejecuta y registra en medios digitales?",
		PorQueImporta: "Digitalizar procesos reduce tiempos, errores y habilita datos operativos reales.",
		ComoEvaluar: "Mapee procesos clave (trámites, control interno, logística) y vea si hay sistema, plantillas o solo papel.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "Casi todo es papel o oral; no hay flujos digitales.", "Expedientes solo físicos; colas sin turno digital.", "Observación de operación 100% manual."},
			{1, "Inicial", "Uso esporádico de ofimática; procesos siguen siendo manuales.", "Se imprime todo lo que se escribe en Word.", "Plantillas sueltas sin flujo."},
			{2, "Básico", "Algunos procesos parciales en sistemas o formularios digitales, con retrabajo a papel.", "Registro digital que luego se imprime y firma en papel.", "Capturas de pantallas de sistemas parciales."},
			{3, "Intermedio", "Procesos prioritarios mayormente digitales, con reglas claras de operación.", "Trámite interno en sistema de workflow; menos papel.", "Mapa de procesos; sistema en producción."},
			{4, "Avanzado", "Automatizaciones, integraciones entre áreas y medición de tiempos/calidad del proceso.", "Alertas automáticas; tableros de tiempos de trámite.", "KPI de procesos; integraciones documentadas."},
			{5, "Óptimo", "Procesos optimizados de punta a punta, medidos y en mejora continua basada en datos.", "Rediseño continuo; eliminación de pasos sin valor.", "Ciclo de mejora documentado; resultados medibles."},
		},
	},
	{
		Nombre: "Servicios al ciudadano",
		PreguntaGuia: "¿Cómo interactúa el ciudadano/usuario con la entidad por canales digitales?",
		PorQueImporta: "La transformación digital pública se valida en la experiencia real del ciudadano.",
		ComoEvaluar: "Presencia web/app, trámites en línea, citas, información actualizada, accesibilidad y soporte.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "Sin canal digital útil; solo presencial o teléfono informal.", "Sin sitio web o sitio abandonado años.", "Verificación de ausencia de canales."},
			{1, "Inicial", "Información estática mínima en web o redes, sin trámites.", "Página con dirección y teléfono desactualizados.", "Captura del sitio."},
			{2, "Básico", "Información actualizada y algún formulario o descarga de modelos.", "PDF de requisitos; formulario de contacto.", "Formularios publicados."},
			{3, "Intermedio", "Trámites o solicitudes digitales con seguimiento básico.", "Cita en línea; estado de solicitud por código.", "Sistema de trámites en uso."},
			{4, "Avanzado", "Servicios transaccionales integrados, con identidad del usuario y métricas de uso.", "Pagos o validaciones en línea según marco legal; panel de uso.", "Estadísticas de uso; integración con backends."},
			{5, "Óptimo", "Experiencia omnicanal, accesible, con mejora continua basada en satisfacción y datos.", "Múltiples canales coherentes; encuestas; ajustes frecuentes.", "Medición de satisfacción; roadmap de servicio."},
		},
	},
	{
		Nombre: "Capacidades del personal",
		PreguntaGuia: "¿El personal tiene competencias digitales y de datos acordes a su rol?",
		PorQueImporta: "La tecnología sin personas capacitadas no se adopta ni se sostiene.",
		ComoEvaluar: "Planes de formación, nivel de uso real de herramientas, perfiles TIC y cultura de aprendizaje.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "Personal sin habilidades digitales básicas para el puesto.", "Dependencia total de una sola persona para cualquier tarea en PC.", "Observación / entrevistas."},
			{1, "Inicial", "Uso básico autodidacta; sin plan de formación.", "Algunos usan WhatsApp o Word a nivel elemental.", "Sin plan de capacitación."},
			{2, "Básico", "Capacitaciones puntuales; usuarios operan sistemas prioritarios con apoyo.", "Curso ocasional; manuales internos breves.", "Registros de asistencia a cursos."},
			{3, "Intermedio", "Plan de formación digital por roles; competencias definidas.", "Rutas de aprendizaje; evaluación básica de competencias.", "Plan anual de capacitación."},
			{4, "Avanzado", "Competencias de datos/TIC extendidas; multiplicadores internos; actualización continua.", "Referentes digitales por área; comunidades de práctica.", "Evidencias de formación continua."},
			{5, "Óptimo", "Cultura de aprendizaje digital consolidada; atracción y retención de talento especializado.", "Perfiles de datos/IA o avanzados TIC; carrera profesional.", "Política de talento digital; indicadores de clima/retención."},
		},
	},
	{
		Nombre: "Marco normativo y cumplimiento",
		PreguntaGuia: "¿La entidad aplica y documenta el marco legal y normativo digital y de datos aplicable?",
		PorQueImporta: "La legitimidad y la protección de derechos dependen del cumplimiento normativo.",
		ComoEvaluar: "Políticas internas, registros exigidos, protección de datos personales, alineación a resoluciones ministeriales.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "Desconocimiento o incumplimiento evidente de normas digitales/datos aplicables.", "Sin políticas; datos personales sin control.", "Hallazgos de incumplimiento."},
			{1, "Inicial", "Conocimiento parcial; aplicación inconsistente.", "Se conoce la norma pero no hay procedimientos.", "Entrevistas; ausencia de procedimientos."},
			{2, "Básico", "Algunas políticas o procedimientos escritos alineados a normas prioritarias.", "Política interna breve de uso de TIC o de datos.", "Documento aprobado."},
			{3, "Intermedio", "Cumplimiento sistemático de obligaciones principales con responsables claros.", "Registros, consentimientos o controles según marco aplicable.", "Evidencias de cumplimiento rutinario."},
			{4, "Avanzado", "Seguimiento proactivo de cambios normativos y ajustes oportunos.", "Revisión periódica del marco; checklist de cumplimiento.", "Actas de revisión normativa."},
			{5, "Óptimo", "Cultura de cumplimiento, auditoría interna y mejora ante hallazgos.", "Auditorías; planes correctivos cerrados a tiempo.", "Informes de auditoría y seguimiento."},
		},
	},
	{
		Nombre: "Interoperabilidad",
		PreguntaGuia: "¿Los sistemas y áreas intercambian datos de forma acordada, segura y reutilizable?",
		PorQueImporta: "Sin interoperabilidad se duplican esfuerzos y se fragmenta la visión del ciudadano y de la gestión.",
		ComoEvaluar: "Interfaces, estándares, convenios de intercambio, silos conocidos y reutilización de datos entre sistemas.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "Sistemas y áreas totalmente aislados; intercambio solo por papel o USB informal.", "Cada sistema es un silo; reescritura manual de datos.", "Arquitectura de sistemas sin enlaces."},
			{1, "Inicial", "Intercambios esporádicos manuales o archivos enviados por correo.", "Exportar CSV y reimportar a mano.", "Ejemplos de archivos intercambiados."},
			{2, "Básico", "Algunas interfaces puntuales o bases compartidas limitadas entre sistemas propios.", "Dos sistemas de la entidad comparten una tabla o carpeta controlada.", "Descripción de la interfaz puntual."},
			{3, "Intermedio", "Estándares internos de intercambio; convenios con al menos un actor externo o sistema nacional/sectorial.", "API o formato acordado; convenio de intercambio.", "Documento de interfaz; convenio."},
			{4, "Avanzado", "Interoperabilidad planificada en la arquitectura; reutilización sistemática de datos.", "Catálogo de servicios de datos; menos duplicación.", "Mapa de integraciones."},
			{5, "Óptimo", "Ecosistema interoperable medido, seguro y alineado a marcos nacionales/sectoriales de intercambio.", "Participación estable en plataformas de interoperabilidad; monitoreo de intercambios.", "Métricas de intercambio; certificaciones o adhesiones formales."},
		},
	},
	{
		Nombre: "Preparación para IA (Pirámide Funcional)",
		PreguntaGuia: "¿Están sentadas las bases (datos, procesos, ética, casos de uso) antes de adoptar IA?",
		PorQueImporta: "La Resolución insiste en no saltar a IA sin datos y capacidades previas; evita proyectos fallidos.",
		ComoEvaluar: "Madurez de datos y procesos previos, casos de uso priorizados, riesgos éticos y capacidad de evaluar proveedores/modelos.",
		Niveles: []CriterioNivel{
			{0, "Inexistente", "No hay bases de datos/procesos; se habla de IA solo como moda, sin casos reales.", "Propuestas de IA sin inventario de datos ni problema claro.", "Ausencia de casos de uso documentados."},
			{1, "Inicial", "Interés exploratorio; experimentos personales sin gobierno ni datos preparados.", "Pruebas con herramientas genéricas sin datos institucionales.", "Demos aisladas."},
			{2, "Básico", "Se reconocen precondiciones (datos, calidad, procesos) y hay al menos un caso de uso hipotético priorizado.", "Documento que pospone IA hasta mejorar datos; lista de ideas de casos de uso.", "Nota de priorización."},
			{3, "Intermedio", "Casos de uso seleccionados con datos disponibles y criterios de éxito; riesgos identificados.", "Pilot piloto definido (ej. clasificación de documentos) con dataset y métrica.", "Ficha de caso de uso; dataset identificado."},
			{4, "Avanzado", "Piloto(s) evaluados, con gobernanza, supervisión humana y medición de resultados.", "Piloto en producción limitada; revisión ética/operativa.", "Informe de resultados del piloto."},
			{5, "Óptimo", "Capacidad institucional para incorporar IA de forma responsable, medible y sostenible.", "Portafolio de casos; políticas de uso responsable; mejora continua.", "Política de IA responsable; portafolio activo."},
		},
	},
}

// GuiaPorNombre devuelve la guía de una dimensión o nil.
func GuiaPorNombre(nombre string) *GuiaDimension {
	for i := range Guias {
		if Guias[i].Nombre == nombre {
			return &Guias[i]
		}
	}
	return nil
}
