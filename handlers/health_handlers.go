package handlers

import (
	. "github.com/tbxark/g4vercel"
	"github.com/umarkotak/animapu-lite-lambda/utils"
)

func GetHealth(context *Context) {
	utils.RenderResponse(
		context,
		map[string]interface{}{
			"status": "healthy",
		},
		nil,
		200,
	)
}
