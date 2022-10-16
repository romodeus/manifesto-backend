package service

import (
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

	result, _ := s.Repo.Get(ManifestEntity)
	if result.CustomURL != "" {
		return manifestentity.Manifest{}, exceptions.NewBadRequestError("custom url has been used")
	}

	err := s.Repo.Save(ManifestEntity)
	if err != nil {
		return manifestentity.Manifest{}, err
	}

	return s.Repo.Get(ManifestEntity)
}

func (s *manifestService) Get(ManifestEntity manifestentity.Manifest) (manifestentity.Manifest, error) {
	return s.Repo.Get(ManifestEntity)
}
