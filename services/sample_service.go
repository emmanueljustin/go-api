package services

import "github.com/emmanueljustin/go-api/repositories"

func GetSampleData() (string, error) {
	return repositories.GetSample()
}
