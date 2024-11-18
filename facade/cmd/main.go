package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/CloudNativeGame/kruise-game-api/internal/updater"
	"github.com/CloudNativeGame/kruise-game-api/pkg/options"
	"log/slog"
	"os"
	"time"
)

// kruise-game-api --filter={} --patch={} --kubeconfig=~/.kube/config
func main() {
	var filter string
	var jsonPatch string
	var kubeConfigPath string

	flag.StringVar(&filter, "filter", "", "filter for the game servers")
	flag.StringVar(&jsonPatch, "patch", "", "jsonPatch for the game servers")
	flag.StringVar(&kubeConfigPath, "kubeconfig", "", "path of the kube config")

	flag.Parse()

	kubeOption := options.KubeOption{
		KubeConfigPath:          kubeConfigPath,
		InformersReSyncInterval: time.Second * 30,
	}

	f := filter.NewFilter(&filter.FilterOption{
		KubeOption: kubeOption,
	})

	gameServers, err := f.GetFilteredGameServers(filter)
	if err != nil {
		slog.Error("get filtered GameServers failed", "error", err)
		os.Exit(1)
	}

	if jsonPatch != "" {
		u := updater.NewUpdater(&updater.UpdaterOption{
			KubeOption: kubeOption,
		})

		results := u.Update(gameServers, []byte(jsonPatch))
		resultsJson, err := json.Marshal(results)
		if err != nil {
			slog.Error("marshal GameServers update results failed", "error", err)
			os.Exit(1)
		}
		fmt.Println(resultsJson)
	} else {
		resultsJson, err := json.Marshal(gameServers)
		if err != nil {
			slog.Error("marshal GameServers failed", "error", err)
			os.Exit(1)
		}

		fmt.Println(resultsJson)
	}
}
