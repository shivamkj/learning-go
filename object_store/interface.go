package object_store

type (
	ObjectStore interface {
		Download() (File, error)
		Upload()
		Delete()
	}

	File struct {
		Name string
	}
)
