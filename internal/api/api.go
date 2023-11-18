package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Kartochnik010/test-cat-breeds/config"
	"github.com/Kartochnik010/test-cat-breeds/internal/models"
)

func FetchBreeds(config *config.Config) ([]models.Breed, error) {
	allBreeds := []models.Breed{}
	wg := &sync.WaitGroup{}

	breedsChannel := make(chan []models.Breed, 1)
	errChannel := make(chan error, 1)
	doneChannel := make(chan bool, 1)

	var firstPageResult models.Response
	resp, err := http.Get(fmt.Sprintf("%s?page=1", config.BaseURL))
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&firstPageResult); err != nil {
		resp.Body.Close()
		return nil, err
	}
	if len(firstPageResult.Data) == 0 {
		return nil, errors.New("didn't decode as it should have")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Request returned non 200 status")
	}
	resp.Body.Close()

	totalPages := firstPageResult.LastPage

	for page := 1; page <= totalPages; page++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			resp, err := http.Get(fmt.Sprintf("%s?page=%d", config.BaseURL, page))
			if err != nil {
				errChannel <- err
				return
			}
			if resp.StatusCode != http.StatusOK {
				errChannel <- errors.New("Request returned non 200 status")
				return
			}
			defer resp.Body.Close()

			var result models.Response
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				errChannel <- err
				return
			}
			if len(result.Data) == 0 {
				errChannel <- errors.New("didn't decode as it should have")
				return
			}

			breedsChannel <- result.Data
		}(page)

		time.Sleep(time.Millisecond * time.Duration(config.RequestInterval))
	}

	go func() {
		wg.Wait()
		doneChannel <- true
	}()

	for {
		select {
		case breeds := <-breedsChannel:
			allBreeds = append(allBreeds, breeds...)
			// fmt.Println("got:", len(breeds))
		case err := <-errChannel:
			if err == nil {
				continue
			}
			return nil, err
		case <-doneChannel:
			close(errChannel)
			close(breedsChannel)
			return allBreeds, nil
		}
	}
}
