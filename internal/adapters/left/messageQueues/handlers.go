package MessageQueue

import (
	"auth-service/internal/ports/app"
	"log"
)

func trialHandler(body message, app app.App) bool {
	log.Println(body)
	return true
}
