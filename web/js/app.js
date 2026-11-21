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

const dimensiones = [
  "Gobernanza y liderazgo digital",
  "Infraestructura tecnológica",
  "Disponibilidad y calidad de datos",
  "Seguridad informática y ciberseguridad",
  "Procesos de trabajo digitalizados",
  "Servicios al ciudadano",
  "Capacidades del personal",
  "Marco normativo y cumplimiento",
  "Interoperabilidad",
  "Preparación para IA (Pirámide Funcional)"
];

// Navigation
document.querySelectorAll(".nav-item").forEach(btn => {
  btn.addEventListener("click", () => {
    document.querySelectorAll(".nav-item").forEach(b => b.classList.remove("active"));
    document.querySelectorAll(".view").forEach(v => v.classList.remove("active"));
    btn.classList.add("active");
    const view = btn.dataset.view;
    document.getElementById("view-" + view).classList.add("active");
    document.getElementById("view-title").textContent = btn.textContent;
  });
});

// Entity name
document.getElementById("entity-name").addEventListener("input", e => {
  state.entityName = e.target.value || "Entidad sin nombre";
});

// Build diagnosis form
const form = document.getElementById("dimensiones-form");
dimensiones.forEach((d, i) => {
  const row = document.createElement("div");
  row.className = "dim-row";
  row.innerHTML = `
    <label>${i + 1}. ${d}</label>
    <select data-dim="${d}">
      <option value="0">0 – Inexistente</option>
      <option value="1">1 – Inicial</option>
      <option value="2" selected>2 – Básico</option>
      <option value="3">3 – Intermedio</option>
      <option value="4">4 – Avanzado</option>
      <option value="5">5 – Óptimo</option>
    </select>
  `;
  form.appendChild(row);
});

// Evaluate diagnosis
document.getElementById("btn-evaluar").addEventListener("click", async () => {
  const respuestas = [];
  document.querySelectorAll("#dimensiones-form select").forEach(sel => {
    respuestas.push({
      dimension: sel.dataset.dim,
      nivel: parseInt(sel.value, 10),
      evidencia: "",
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
