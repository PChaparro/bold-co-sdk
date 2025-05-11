package definitions

// GetBindedTerminalsForIntegrationsAPIResponse represents the API response
// for retrieving available terminals for integration.
type GetBindedTerminalsForIntegrationsAPIResponse struct {
	// Payload contains the information of the available terminals.
	Payload GetBindedTerminalsPayload `json:"payload"`
	// Errors contains any errors that occurred during the request.
	Errors []ErrorField `json:"errors,omitempty"`
}

// GetBindedTerminalsPayload contains the information of available terminals.
type GetBindedTerminalsPayload struct {
	AvailableTerminals *[]TerminalInfo `json:"available_terminals,omitempty"`
}

// TerminalInfo contains the information of a specific terminal.
type TerminalInfo struct {
	TerminalModel  string         `json:"terminal_model"`
	TerminalSerial string         `json:"terminal_serial"`
	Status         TerminalStatus `json:"status"`
	Name           string         `json:"name"`
}

// TerminalStatus constants for the status of the terminals.
type TerminalStatus string

const (
	// TerminalStatusBinded represents a terminal binded to the integration.
	TerminalStatusBinded TerminalStatus = "BINDED"
)
