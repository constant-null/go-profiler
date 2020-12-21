package storage

import "github.com/google/profile_storage/profile"

type ProfileData interface {
	StoreTop(app, sampleType string, sample profile.Sample) error
}
