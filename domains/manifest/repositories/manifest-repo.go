package manifestrepo

import (
	manifestentity "manifesto/domains/manifest/entities"
	"manifesto/exceptions"

	"github.com/go-redis/redis"
)

type manifestRepo struct {
	Rds *redis.Client
}

func New(rediscli *redis.Client) *manifestRepo {
	return &manifestRepo{
		Rds: rediscli,
	}
}

func (r *manifestRepo) Save(ManifestEntity manifestentity.Manifest) error {
	if err := r.Rds.Set(ManifestEntity.CustomURL, ManifestEntity.RealURL, 0); err.Err() != nil {
		return exceptions.NewInternalServerError(err.String())
	}
	return nil
}

func (r *manifestRepo) Get(ManifestEntity manifestentity.Manifest) (manifestentity.Manifest, error) {
	URLString, err := r.Rds.Get(ManifestEntity.CustomURL).Result()
	if err == redis.Nil {
		return manifestentity.Manifest{}, exceptions.NewNotFoundError(err.Error())
	} else if err != nil {
		return manifestentity.Manifest{}, exceptions.NewInternalServerError(err.Error())
	}

	return manifestentity.Manifest{
		CustomURL: ManifestEntity.CustomURL,
		RealURL:   URLString,
	}, nil
}
