package profiler

import (
	"bytes"
	"context"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/pkg/errors"

	"gitlab.corp.mail.ru/otvetmailru/profiler/profile"
)

// Profiler collect profiles from application and sends them to storage
type Profiler struct {
	Client      StorageClient
	Interval    time.Duration
	MaxDuration time.Duration

	start sync.Once
}

// Start profiler
// By default MaxDuration is set to Interval value
// If Interval is not specified default 10 seconds interval is used
func (p *Profiler) Start() error {
	if p.MaxDuration > p.Interval {
		return errors.New("max duration should not be bigger then interval")
	}

	p.start.Do(func() {
		if p.Interval <= 0 {
			p.Interval = 10 * time.Second
		}
		if p.MaxDuration <= 0 {
			p.MaxDuration = p.Interval
		}

		runtime.SetBlockProfileRate(100)

		for range time.Tick(p.Interval) {
			ctx, _ := context.WithTimeout(context.Background(), p.MaxDuration)
			p.collectAndUpload(ctx)
		}
	})

	return nil
}

func (p *Profiler) collectAndUpload(ctx context.Context) {
	var collectors = map[Type]func(context.Context) (*profile.Profile, error){
		CPU:    p.collectCPUProf,
		Memory: p.collectMemProf,
		Mutex:  p.collectMutexProf,
	}

	for t, collector := range collectors {
		prof, err := collector(ctx)
		if err != nil {
			debugLog("failed to collect profile: %v", err)
		}

		if err = p.Client.Upload(ctx, t, prof); err != nil {
			debugLog("failed to send profile to storage: %v", err)
		}
	}
}

func (p *Profiler) collectMemProf(ctx context.Context) (*profile.Profile, error) {
	var buf bytes.Buffer
	if err := pprof.WriteHeapProfile(&buf); err != nil {
		return nil, err
	}

	return profile.Parse(&buf)
}

func (p *Profiler) collectCPUProf(ctx context.Context) (*profile.Profile, error) {
	var buf bytes.Buffer
	if err := pprof.StartCPUProfile(&buf); err != nil {
		return nil, err
	}

	select {
	case <-time.After(p.MaxDuration):
	case <-ctx.Done():
	}

	pprof.StopCPUProfile()

	return profile.Parse(&buf)
}

func (p *Profiler) collectMutexProf(ctx context.Context) (*profile.Profile, error) {
	var buf bytes.Buffer
	mp := pprof.Lookup("mutex")
	if p == nil {
		return nil, errors.New("mutex profiling is not supported")
	}
	if err := mp.WriteTo(&buf, 0); err != nil {
		return nil, err
	}

	return profile.Parse(&buf)
}
