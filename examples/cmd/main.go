package main

import (
	"log/slog"
	"os/exec"
)

func main() {
	command := "kruise-game-api"
	args := []string{"--filter={\"opsState\": \"None\"}", "--kubeconfig=~/.kube/config"}

	cmd := exec.Command(command, args...)

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	slog.Info(string(output))
}
