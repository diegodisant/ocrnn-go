package storage

type Exportable interface {
	Export(model ModelFiles)
}
