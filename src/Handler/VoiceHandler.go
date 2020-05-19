package Handler

import (
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

func registerPostVoice(h VoiceHandler, w http.ResponseWriter, r *http.Request) (* Voice.Voice, error) {
	file, _, err := r.FormFile("voice")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	defer file.Close()

	name := r.FormValue("name")

	powerSpectrum := h.voiceService.CalculatePowerSpectrum(file)
	return h.voiceService.Add(name, powerSpectrum)
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
	voice, err := registerPostVoice(h, w, r)
	if err != nil {
		panic(err)
	}

	template, err := createTemplate("register.html")
	if err != nil {
		panic(err)
	}

	response := new(RegisterResponse)
	response.ID = voice.ID
	response.Name = voice.Name
	response.Host = r.Header.Get("HOST")

	err = template.Execute(w, response)
	if err != nil {
		panic(err)
	}
}

type SimilarityResponse struct {
	ID string
	Name string
	TrainingName string
	Similarity int
	Host string
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
	} else if r.Method == http.MethodPost {
		voice, err := registerPostVoice(h, w, r)
		if err != nil {
			panic(err)
		}

		// cosine類似度の計算
		ID := r.FormValue("id")
		training, err := h.voiceService.Get(ID)
		if err != nil {
			panic(err)
		}

		similarity := h.voiceService.CosSimilarity(voice.PowerSpectrum, training.PowerSpectrum)
		response := new(SimilarityResponse)
		response.ID = voice.ID
		response.Name = voice.Name
		response.TrainingName = training.Name
		response.Similarity = int(similarity * 100)
		response.Host = r.Header.Get("HOST")

		template, err := createTemplate("result.html")
		if err != nil {
			panic(err)
		}

		err = template.Execute(w, response)
		if err != nil {
			panic(err)
		}
	}
}