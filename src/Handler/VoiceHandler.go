package Handler

import (
	"os"
	"log"
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

type RegisterResponse struct {
	ID string
	Name string
	Host string
}

func (h VoiceHandler) Register(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("voice")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	name := r.FormValue("name")
	log.Print(name)

	powerSpectrum := h.voiceService.CalculatePowerSpectrum(file)
	voice, err := h.voiceService.Add(name, powerSpectrum)
	if err != nil {
		panic(err)
	}

	template, err := createTemplate("register.html")
	if err != nil {
		panic(err)
	}

	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	response := new(RegisterResponse)
	response.ID = voice.ID
	response.Name = voice.Name
	response.Host = hostName

	err = template.Execute(w, response)
	if err != nil {
		panic(err)
	}
}

func (h VoiceHandler) Similarity(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		samples, ok := r.URL.Query()["sample"]

		if !ok || len(samples[0]) < 1 {
			log.Println("Url Param 'sample' is missing")
			return
		}

		sample := samples[0]

		voice, err := h.voiceService.Get(sample)
		if err != nil {
			panic(err)
		}

		template, err := createTemplate("similarity.html")
		if err != nil {
			panic(err)
		}

		err = template.Execute(w, voice)
		if err != nil {
			panic(err)
		}
	}
}
