package redoc

import (
	"bytes"
	_ "embed"
	"errors"
	"go-start/internal/pkg/log"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Redoc struct {
	Title, DocPath, SpecPath, SpecFile, Desc string
}

var (
	//go:embed static/index.html
	RedocHTML string
	//go:embed static/redoc.standalone.js
	RedocJavaScript string
	// SpecNotFountErr Spec文件缺失
	SpecNotFountErr = errors.New("spec not found")
	// RedocRenderErr Redoc渲染失败
	RedocRenderErr = errors.New("redoc render failed")
)

func (rd *Redoc) Handler() http.HandlerFunc {
	by, err := rd.Render()
	if err != nil {
		log.Logger.Error(RedocRenderErr)
		return nil
	}
	if rd.SpecPath == "" || rd.SpecFile == "" {
		log.Logger.Error(SpecNotFountErr)
		return nil
	}
	return func(w http.ResponseWriter, r *http.Request) {
		allowedMethod := strings.ToUpper(r.Method)
		if allowedMethod != "GET" && allowedMethod != "HEAD" {
			return
		}
		if strings.HasSuffix(r.URL.Path, rd.SpecPath) {
			spec, err := os.ReadFile(rd.SpecFile)
			if err != nil {
				log.Logger.Error(SpecNotFountErr)
				return
			}
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write(spec)
		}
		if rd.DocPath == "" || rd.DocPath == r.URL.Path {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/html")
			w.Write(by)
		}
	}
}

func (rd *Redoc) Render() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	tpl, err := template.New("Redoc").Parse(RedocHTML)
	if err != nil {
		return nil, err
	}
	err = tpl.Execute(buf, map[string]string{
		"js":    RedocJavaScript,
		"title": rd.Title,
		"url":   rd.SpecPath,
		"desc":  rd.Desc,
	})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
