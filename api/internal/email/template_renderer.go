package email

import (
	"bytes"
	"html/template"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(globPattern string) (*TemplateRenderer, error) {
	templates, err := template.ParseGlob(globPattern)
	if err != nil {
		return nil, err
	}
	return &TemplateRenderer{
		templates: templates,
	}, nil
}

func (r *TemplateRenderer) RenderTemplate(templateName string, data interface{}) (string, error) {
	var buf bytes.Buffer
	err := r.templates.ExecuteTemplate(&buf, templateName, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
