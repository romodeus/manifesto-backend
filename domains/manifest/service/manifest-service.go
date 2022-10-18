package service

import (
	"fmt"
	manifestentity "manifesto/domains/manifest/entities"
	"manifesto/exceptions"
	"manifesto/utils/helpers"
)

type manifestService struct {
	Repo manifestentity.IRepoManifest
}

func New(repo manifestentity.IRepoManifest) *manifestService {
	return &manifestService{
		Repo: repo,
	}
}

func (s *manifestService) CustomURLCreate(ManifestEntity manifestentity.Manifest) (manifestentity.Manifest, error) {
	ManifestEntity.CustomURL = helpers.URIFormat(ManifestEntity.CustomURL)

	result, err := s.Repo.Get(ManifestEntity)
	if err != nil {
		fmt.Println("get first", err)
	}
	if result.CustomURL != "" {
		return manifestentity.Manifest{}, exceptions.NewBadRequestError("custom url has been used")
	}

	err = s.Repo.Save(ManifestEntity)
	if err != nil {
		fmt.Println("error after save", err)
		return manifestentity.Manifest{}, err
	}

	result1, err := s.Repo.Get(ManifestEntity)
	if err != nil {
		fmt.Println(result1)
		fmt.Println("error get second", err)
	}

	return result1, err
}

func (s *manifestService) Get(ManifestEntity manifestentity.Manifest) (manifestentity.Manifest, error) {
	return s.Repo.Get(ManifestEntity)
}
