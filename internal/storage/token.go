package storage

type Token interface {
	Get() (string, error)
	Set(string) error
}
