package traceid

import (
	"github.com/aidarkhanov/nanoid"
)

type TraceID struct {
	id string
}

func (t *TraceID) IsZero() bool {
	return t.id == ""
}

func (t *TraceID) Value() string {
	return t.id
}

func GenerateTraceID() *TraceID {
	return &TraceID{id: nanoid.New()}
}

func GenerateEmptyTraceID() *TraceID {
	return &TraceID{id: ""}
}

func GenerateTraceIDByID(id string) *TraceID {
	return &TraceID{id: id}
}
