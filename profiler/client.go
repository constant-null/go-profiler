package profiler

import (
	"context"

	"gitlab.corp.mail.ru/otvetmailru/profiler/profile"
)

// Type contains profile types
type Type string

// Types supported by profiler
const (
	CPU    Type = "cpu"
	Memory      = "memory"
	Mutex       = "mutex"
)

// StorageClient base interface for accessing remote profile storage
type StorageClient interface {
	Upload(context.Context, Type, *profile.Profile) error
}
