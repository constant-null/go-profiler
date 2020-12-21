package storage

type Profile interface {
	Store(app, ptype string, profData []byte) (int64, error)
}
