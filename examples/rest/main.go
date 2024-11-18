package main

import (
	"fmt"
	kruisegameapiclient "github.com/CloudNativeGame/kruise-game-api/facade/rest/client"
	"github.com/CloudNativeGame/kruise-game-api/pkg/filter"
	jsonpatchbuilder "github.com/CloudNativeGame/kruise-game-api/pkg/jsonpatches/builder"
	"log/slog"
)

func main() {
	client := kruisegameapiclient.NewKruiseGameApiHttpClient()
	filter := filter.NewGsFilterBuilder().OpsState("None")
	gameServers, err := client.GetGameServers(filter.NewGsFilterBuilder().OpsState("None"))
	if err != nil {
		slog.Error("get filtered GameServers failed", "error", err)
		panic(err)
	}

	slog.Info(fmt.Sprintf("%d GameServers matched filter %s", len(gameServers), filter.Build()))
	for _, gs := range gameServers {
		slog.Info("filtered GS", "gs", gs)
	}

	results, err := client.UpdateGameServers(filter.NewGsFilterBuilder().UpdatePriority(1),
		jsonpatchbuilder.NewGsJsonPatchBuilder())
	if err != nil {
		return
	}

	for _, result := range results {
		if result.Err != nil {
			slog.Error(fmt.Sprintf("update GameServer %s/%s failed", result.Gs.Namespace, result.Gs.Name),
				"error", result.Err.Error())
		} else {
			slog.Info(fmt.Sprintf("update GameServer %s/%s success", result.Gs.Namespace, result.Gs.Name),
				"gs", result.Gs, "updatedGs", result.UpdatedGs)
		}
	}
}
