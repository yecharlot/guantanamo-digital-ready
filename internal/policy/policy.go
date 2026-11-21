package policy

// Engine is a minimal policy engine for GDR.
// In this MVP it enforces basic rules aligned with "Primero los datos".
type Engine struct{}

func New() *Engine { return &Engine{} }

// CanGenerateStrategy returns true only if a diagnosis and inventory exist.
func (e *Engine) CanGenerateStrategy(hasDiagnostico, hasInventario bool) bool {
	return hasDiagnostico && hasInventario
}

// CanExportExpediente requires at least a diagnosis.
func (e *Engine) CanExportExpediente(hasDiagnostico bool) bool {
	return hasDiagnostico
}
