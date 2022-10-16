package manifestentities

type IRepoManifest interface {
	Save(ManifestEntity Manifest) error
	Get(ManifestEntity Manifest) (Manifest, error)
	Update(ManifestEntity Manifest) error
}

type IServiceManifest interface {
	Create(ManifestEntity Manifest) error
	Get(ManifestEntity Manifest) (Manifest, error)
	Update(ManifestEntity Manifest) error
}
