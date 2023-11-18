package main

import (
	"flag"
	"log"
	"time"

	"github.com/Kartochnik010/test-cat-breeds/config"
	"github.com/Kartochnik010/test-cat-breeds/internal/api"
	"github.com/Kartochnik010/test-cat-breeds/utils"
)

func main() {
	showTime := flag.Bool("t", false, "Display the execution time of the application")
	flag.Parse()
	defer func(t time.Time) {
		if *showTime {
			log.Printf("Executed in %vs\n", time.Since(t).Seconds())
		}
	}(time.Now())

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}

	breeds, err := api.FetchBreeds(cfg)
	if err != nil {
		log.Fatalln("Error fetching:", err)
	}

	groupedBreeds := utils.GroupByCountry(breeds)
	sortedBreeds := utils.SortBreeds(groupedBreeds)

	err = utils.SaveToFileJSON(sortedBreeds, cfg.OutFileName)
	if err != nil {
		log.Fatalln("Error saving file:", err)
	}
}
