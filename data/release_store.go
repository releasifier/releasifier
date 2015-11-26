package data

import "upper.io/bond"

//ReleaseStore store for release
type ReleaseStore struct {
	bond.Store
}

func (s ReleaseStore) CreateRelease(version Version, platform Platform, note string, userID, appID int64) (*Release, error) {
	return nil, nil
}
