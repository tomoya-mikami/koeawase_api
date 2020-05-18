package Handler

import (
	"net/http"
	"text/template"
	Voice "local.packages/voice"
)

type VoiceHandler struct {
	voiceService Voice.ServiceInterface
}

const TEMPLATE_PATH = "template/"

func createTemplate(templateName string) (*template.Template, error) {
	return template.ParseFiles(TEMPLATE_PATH + templateName)
}

func NewVoiceHandler(voiceService Voice.ServiceInterface) *VoiceHandler {
	handler := new(VoiceHandler)
	handler.voiceService = voiceService
	return handler
}

func (h VoiceHandler) Index(w http.ResponseWriter, r *http.Request) {
	template, err := createTemplate("index.html")
	if err != nil {
		panic(err)
	}

	err = template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}
