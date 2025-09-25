package local

import (
	"compress/gzip"
	"os"
	"path"
	"time"

	"github.com/go-profiler/profiler"
	"github.com/go-profiler/profiler/profile"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// Storage stores files in local file systems
type Storage struct {
	BaseDir string
}

// Upload saves profile in local file system
func (s *Storage) Upload(t profiler.Type, p *profile.Profile) error {
	pd := path.Join(s.BaseDir, string(t))
	os.MkdirAll(pd, 0755)

	fn := time.Now().Format("02-Jan-06_15-04-05") + ".prof"
	f, err := os.Create(path.Join(pd, fn))
	if err != nil {
		return errors.Wrap(err, "unable to create profile file")
	}

	data, err := proto.Marshal(p)
	if err != nil {
		return errors.Wrap(err, "unable to marshal profile")
	}

	gzw := gzip.NewWriter(f)
	defer gzw.Close()
	if _, err = gzw.Write(data); err != nil {
		return errors.Wrap(err, "error while writing profile")
	}

	return gzw.Close()
}
