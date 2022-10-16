package manifestentity

type IRepoManifest interface {
	Save(ManifestEntity Manifest) error
	Get(ManifestEntity Manifest) (Manifest, error)
}

type IServiceManifest interface {
	Create(ManifestEntity Manifest) (Manifest, error)
	Get(ManifestEntity Manifest) (Manifest, error)
}
