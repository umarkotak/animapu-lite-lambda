package utils

import (
	. "github.com/tbxark/g4vercel"
)

func RenderResponse(geeCtx *Context, bodyPayload interface{}, err error, status int) {
	success := true
	if status != 200 {
		success = false
	}

	geeCtx.SetHeader("Access-Control-Allow-Origin", "*")
	geeCtx.SetHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	geeCtx.SetHeader(
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Animapu-User-Uid, X-Visitor-Id, X-From-Path, Server, X-Vercel-Cache, X-Vercel-Id",
	)

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	geeCtx.JSON(status, H{
		"success": success,
		"data":    bodyPayload,
		"error":   errMsg,
	})
}
