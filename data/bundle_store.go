package data

import "upper.io/bond"

//BundleStore store for bundle
type BundleStore struct {
	bond.Store
}

func (s BundleStore) FindByHash(hash string) (*Bundle, error) {
	var bundle *Bundle

	return bundle, nil
}
