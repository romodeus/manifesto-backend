package service

import (
	manifestentity "manifesto/domains/manifest/entities"
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

func (s *manifestService) Create(ManifestEntity manifestentity.Manifest) error {
	ManifestEntity.CustomURL = helpers.URIFormat(ManifestEntity.CustomURL)
	return s.Repo.Save(ManifestEntity)
}

func (s *manifestService) Get(ManifestEntity manifestentity.Manifest) (manifestentity.Manifest, error) {
	return s.Repo.Get(ManifestEntity)
}
