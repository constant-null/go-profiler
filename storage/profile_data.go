package storage

import "github.com/go-profiler/profiler/profile"

type ProfileData interface {
	StoreTop(app, sampleType string, sample profile.Sample) error
}
