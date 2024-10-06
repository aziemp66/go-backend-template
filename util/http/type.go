package util_http

type (
	Response struct {
		Data     any      `json:"data,omitempty"`
		Metadata Metadata `json:"metadata"`
	}

	Metadata struct {
		Timestamp string `json:"timestamp"`
		Message   string `json:"message"`
	}

	Error struct {
		Message string `json:"message"`
	}

	TraceType string
)
