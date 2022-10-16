package manifestentity

type IRepoManifest interface {
	Save(ManifestEntity Manifest) error
	Get(ManifestEntity Manifest) (Manifest, error)
}

type IServiceManifest interface {
	CustomURLCreate(ManifestEntity Manifest) (Manifest, error)
	Get(ManifestEntity Manifest) (Manifest, error)
}
