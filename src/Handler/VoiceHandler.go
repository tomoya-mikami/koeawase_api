package Handler

import (
	"log"
	"net/http"
	"text/template"
	Voice "local.packages/voice"
	Similarity "local.packages/similarity"
)

type VoiceHandler struct {
	voiceService Voice.ServiceInterface
	similarityService Similarity.ServiceInterface
}

const TEMPLATE_PATH = "template/"
const ERROR_FORMAT = "[ERROR] %s"

func createTemplate(templateName string) (*template.Template, error) {
	return template.ParseFiles(TEMPLATE_PATH + templateName)
}

func error500(w http.ResponseWriter) {
	template, _ := createTemplate("500.html")
	template.Execute(w, nil)
}

func error404(w http.ResponseWriter) {
	template, _ := createTemplate("404.html")
	template.Execute(w, nil)
}

func registerPostVoice(h VoiceHandler, w http.ResponseWriter, r *http.Request) (* Voice.Voice, error) {
	file, _, err := r.FormFile("voice")
	if err != nil {
		error500(w)
		return nil, err
	}
	defer file.Close()

	name := r.FormValue("name")

	powerSpectrum := h.voiceService.CalculatePowerSpectrum(file)
	return h.voiceService.Add(name, powerSpectrum)
}

func NewVoiceHandler(
	voiceService Voice.ServiceInterface,
	similarityService Similarity.ServiceInterface,
) *VoiceHandler {
	handler := new(VoiceHandler)
	handler.voiceService = voiceService
	handler.similarityService = similarityService
	return handler
}

func (h VoiceHandler) Index(w http.ResponseWriter, r *http.Request) {
	template, err := createTemplate("index.html")
	if err != nil {
		log.Printf(ERROR_FORMAT, err.Error())
		error500(w)
	}

	err = template.Execute(w, nil)
	if err != nil {
		log.Printf(ERROR_FORMAT, err.Error())
		error500(w)
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
		log.Printf(ERROR_FORMAT, err.Error())
		error500(w)
		return
	}

	template, err := createTemplate("register.html")
	if err != nil {
		log.Printf(ERROR_FORMAT, err.Error())
		error500(w)
		return
	}

	response := new(RegisterResponse)
	response.ID = voice.ID
	response.Name = voice.Name
	response.Host = r.Header.Get("HOST")

	err = template.Execute(w, response)
	if err != nil {
		log.Printf(ERROR_FORMAT, err.Error())
		error500(w)
		return
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
			log.Printf(ERROR_FORMAT, "Url Param 'sample' is missing")
			error404(w)
			return
		}

		sample := samples[0]

		voice, err := h.voiceService.Get(sample)
		if err != nil {
			log.Printf(ERROR_FORMAT, err.Error())
			error404(w)
			return
		}

		template, err := createTemplate("similarity.html")
		if err != nil {
			log.Printf(ERROR_FORMAT, err.Error())
			error500(w)
			return
		}

		err = template.Execute(w, voice)
		if err != nil {
			log.Printf(ERROR_FORMAT, err.Error())
			error500(w)
			return
		}
	} else if r.Method == http.MethodPost {
		voice, err := registerPostVoice(h, w, r)
		if err != nil {
			log.Printf(ERROR_FORMAT, err.Error())
			error500(w)
			return
		}

		// cosine類似度の計算
		ID := r.FormValue("id")
		training, err := h.voiceService.Get(ID)
		if err != nil {
			log.Printf(ERROR_FORMAT, err.Error())
			error404(w)
			return
		}

		similarity, err := h.similarityService.CalcurateSimilarity(voice, &training)
		response := new(SimilarityResponse)
		response.ID = similarity.ID
		response.Name = voice.Name
		response.TrainingName = training.Name
		response.Similarity = int(similarity.Similarity * 100)
		response.Host = r.Header.Get("HOST")

		template, err := createTemplate("result.html")
		if err != nil {
			log.Printf(ERROR_FORMAT, err.Error())
			error500(w)
			return
		}

		err = template.Execute(w, response)
		if err != nil {
			log.Printf(ERROR_FORMAT, err.Error())
			error500(w)
			return
		}
	}
}

type ShareResponse struct {
	ID string
	Name1ID string
	Name1 string
	Name2 string
	Similarity int
}

func (h VoiceHandler) Share(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ids, ok := r.URL.Query()["id"]

		if !ok || len(ids[0]) < 1 {
			log.Println("Url Param 'id' is missing")
			error404(w)
			return
		}

		id := ids[0]

		similarity, err := h.similarityService.Get(id)
		if err != nil {
			error500(w)
		}

		response := new(ShareResponse)
		response.ID = similarity.ID
		response.Name1ID = similarity.Name1ID
		response.Name1 = similarity.Name1
		response.Name2 = similarity.Name2
		response.Similarity = int(similarity.Similarity * 100)

		template, err := createTemplate("share.html")
		if err != nil {
			error500(w)
		}

		err = template.Execute(w, response)
		if err != nil {
			error500(w)
		}
	}
}
