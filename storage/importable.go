package storage

type Importable interface {
	Import(model ModelFiles)
}
