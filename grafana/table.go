package grafana

import "encoding/json"

type CollType string

const (
	CollTime      CollType = "time"
	CollTimestamp CollType = "timestamp"
	CollString    CollType = "string"
	CollNumber    CollType = "number"
	CollUrl       CollType = "url"
)

type Column struct {
	Text string   `json:"text"`
	Type CollType `json:"type"`
}

type Row []interface{}

type Table struct {
	Columns []Column `json:"columns"`
	Rows    []Row    `json:"rows"`
}

func (t Table) MarshalJSON() ([]byte, error) {
	d := []interface{}{map[string]interface{}{
		"columns": t.Columns,
		"rows":    t.Rows,
		"type":    "table",
	}}
	return json.Marshal(d)
}
