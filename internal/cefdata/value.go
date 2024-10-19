package cefdata

type Value struct {
	ParentXID string // guid
	Key       string
	Index     int
	XID       string // guid
	Value     string
	Numeric   float64
	Min, Max  float64
	Type      int
	Unit      int
	ObjectRef string
	Make      string
	Product   string
	Date      string // YYYY-MM-DD
	Private   string
}
