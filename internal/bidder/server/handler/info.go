package handler

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/adgear/go-commons/pkg/buildinfo"
	"github.com/adgear/go-commons/pkg/utils/httputils"
	"github.com/valyala/fasthttp"
)

//go:embed info.html
var f embed.FS

const imageUrl = "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTRBNIhqvYNCMoRGOSuOQ7U1R422dQlHRbmcQ&usqp=CAU"

// InfoHandler handles service info requests and respond html web page with
// build information - git commit hash, git release tag, go language reference.
func infoHandler(serviceName string) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		if ctx == nil {
			return
		}

		tmpl := template.Must(template.ParseFS(f, "info.html"))
		data := buildinfo.NewServiceInfo(serviceName, imageUrl)
		var resp bytes.Buffer
		_ = tmpl.Execute(&resp, data)
		httputils.SetResponseCorsHeaders(ctx)
		ctx.Response.AppendBody(resp.Bytes())
		ctx.Response.Header.Set(fasthttp.HeaderContentType, "text/html")
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
	}
}
