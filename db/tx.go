package db

import "github.com/boltdb/bolt"

// Tx represents a BoltDB transaction
type Tx struct {
	*bolt.Tx
}

// Page retrieves a Page from the database with the given name.
func (tx *Tx) Page(name []byte) (*Page, error) {
	p := &Page{
		Tx:   tx,
		Name: name,
	}

	return p, p.Load()
}

func (tx *Tx) Pages() ([]string, error) {
	p := &Page{
		Tx:   tx,
		Name: []byte(""),
	}
	names, err := p.list()
	return names, err
}
