package Task

import (
	"log"
	"os"
	"fmt"

	Voice "local.packages/voice"
)

type VoiceTask struct {
	voiceService Voice.ServiceInterface
}

func NewVoiceTask(voiceService Voice.ServiceInterface) *VoiceTask {
	cli := new(VoiceTask)
	cli.voiceService = voiceService
	return cli
}

func (v VoiceTask) Execute(taskName string, args []string) {
	switch taskName {
		case "add":
			addVoice(v, args)
			break
		case "calculateSimilarity":
			calculateSimilarity(v, args)
			break
		default:
			log.Fatal("not found task")
	}
}

func addVoice(v VoiceTask, args []string) {
	if len(args) != 2 {
		log.Fatal("provide filename usage: go run main.go cli voice add $name $filename")
	}

	name := args[0]
	fileName := args[1]

	file, err := os.Open("./media/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	powerSpectrum := v.voiceService.CalculatePowerSpectrum(file)
	voice, err := v.voiceService.Add(name, powerSpectrum)
	if err != nil {
		log.Fatal(err)
	}

	str := fmt.Sprintf(
		"add voice id:%s name:%s power spectrum frequency:%d",
		voice.ID,
		voice.Name,
		len(voice.PowerSpectrum),
	)

	fmt.Println(str)
}

func calculateSimilarity(v VoiceTask, args []string) {
	if len(args) != 2 {
		log.Fatal("provide filename usage: go run main.go cli voice calculateSimilarity $id1 $id2")
	}
	id1 := args[0]
	id2 := args[1]

	sample, err := v.voiceService.Get(id1)
	training, err := v.voiceService.Get(id2)
	if err != nil {
		log.Fatal(err)
	}

	calcResult := v.voiceService.CosSimilarity(
		sample.PowerSpectrum,
		training.PowerSpectrum,
	)
	str := fmt.Sprintf(
		"%sさんと%sさんの声の類似度は%fです",
		sample.Name,
		training.Name,
		calcResult,
	)

	fmt.Println(str)
}
