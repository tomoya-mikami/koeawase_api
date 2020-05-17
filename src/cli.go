package src

import (
	"log"
	Task "local.packages/task"
)

type CLI struct {
	voiceTask Task.TaskInterface
}

func NewCLI(voiceTask *Task.VoiceTask) *CLI {
	cli := new(CLI)
	cli.voiceTask = voiceTask
	return cli
}

func (c CLI) Execute(args []string) {
	args = args[1:]
	if args[0] != "cli" || len(args) < 3 {
		log.Fatal("cli mode usage: go main.go cli $domain $taskName $args")
	}

	domain := args[1]
	task := args[2]
	provideArgs := args[3:]
	switch domain {
	case "voice":
		c.voiceTask.Execute(task, provideArgs)
	}
}
