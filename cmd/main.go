package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/invilliafelipeflores/notfiy/internal/config"
	"github.com/invilliafelipeflores/notfiy/internal/github/service"
	"github.com/invilliafelipeflores/notfiy/pkg/notify"
)

func main() {
	var configFile string

	// -c options set configuration file path, but can be overwritten by GAPI_CONFIG_FILE environment variable
	flag.StringVar(&configFile, "c", "env.json", "config file path")
	flag.Parse()

	// Load configuration yaml file using -c location/GAPI_CONFIG_FILE and merging environments variables with higher precedence
	sc, err := config.LoadServiceConfig(configFile)
	if err != nil {
		log.Fatalf("main: could not load service configuration [%v]", err)
	}

	config.Dump("env.json", sc)

	gitHubService := service.NewGithubService(sc.AccessToken)

	for _, repo := range sc.Repos {
		prs, err := gitHubService.GetPulls(repo.Repo)
		if err != nil {
			log.Fatal(err)
		}

		for _, pr := range prs {
			notify.Notify(repo.Name, fmt.Sprintf("%s New Pull Request", repo.Name), pr.URL, "")
		}
	}

}
