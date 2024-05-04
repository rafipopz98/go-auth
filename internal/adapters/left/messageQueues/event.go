package MessageQueue

import (
	"auth-service/internal/ports/app"
	"encoding/json"
)

type message struct {
	Operation   string          `json:"operation"`
	FromService string          `json:"fromService"`
	Data        json.RawMessage `json:"data"`
}

type messageHandler func(message, app.App) bool
