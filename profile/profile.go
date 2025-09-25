package profile

import (
	"os"
	"sort"
	"time"

	"github.com/google/pprof/profile"
	"github.com/jedib0t/go-pretty/table"
)

func Parse(data []byte) {
	p, _ := profile.ParseUncompressed(data)
	smpl := top(p)
	printProfile(smpl)
}

type Sample struct {
	Func  string
	Value int64
}

type Report map[profile.ValueType][]Sample

func top(p *profile.Profile) Report {
	repData := make(map[profile.ValueType]map[string]int64)
	for _, s := range p.Sample {
		if len(s.Location) == 0 || len(s.Location[0].Line) == 0 {
			continue
		}
		fn := s.Location[0].Line[0].Function.Name
		for i, t := range p.SampleType {
			_, ok := repData[*t]
			if !ok {
				repData[*t] = map[string]int64{fn: s.Value[i]}
				continue
			}
			repData[*t][fn] += s.Value[i]
		}
	}

	rep := make(Report)
	for t, s := range repData {
		smpl := make([]Sample, 0, len(s))
		for fn, v := range s {
			smpl = append(smpl, Sample{fn, v})
		}
		sort.Slice(smpl, func(i, j int) bool {
			return smpl[i].Value > smpl[j].Value
		})
		rep[t] = smpl
	}
	return rep
}

const (
	unitNanoseconds = "nanoseconds"
	unitBytes       = "bytes"
)

func printProfile(rep Report) {
	for v, s := range rep {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{v.Type, v.Unit})
		for _, sv := range s {
			var val interface{}
			switch v.Unit {
			case unitNanoseconds:
				val = time.Duration(sv.Value) * time.Nanosecond
			case unitBytes:
				val = sv.Value
			default:
				val = sv.Value
			}

			t.AppendRow(table.Row{sv.Func, val})
		}

		t.Render()
	}
}
