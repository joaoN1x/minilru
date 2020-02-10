package debugger

type Debug struct {
	Service   string `json:"service"`
	File      string `json:"file"`
	Line      string `json:"line"`
	Function  string `json:"function"`
	Message   string `json:"message"`
	Error     string `json:"error"`
	Timestamp string `json:"timestamp"`
}
