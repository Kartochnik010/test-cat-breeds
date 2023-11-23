package utils

import (
	"sort"

	"github.com/Kartochnik010/test-cat-breeds/internal/models"
)

func GroupByCountry(breeds []models.Breed) map[string][]models.Breed {
	grouped := map[string][]models.Breed{}
	for _, breed := range breeds {
		grouped[breed.Country] = append(grouped[breed.Country], breed)
	}
	return grouped
}

func SortBreeds(groupedBreeds map[string][]models.Breed) map[string][]models.Breed {
	for country, breeds := range groupedBreeds {
		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i].Breed) < len(breeds[j].Breed)
		})
		groupedBreeds[country] = breeds
	}
	return groupedBreeds
}
