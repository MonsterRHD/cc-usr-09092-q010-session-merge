package session

import "time"

type Config struct {
	DatabasePath      string
	DefaultGapSeconds int
}

type BoundaryEvent struct {
	EventID        string
	SubjectHint    string
	DeviceID       string
	Kind           string
	LogicalVersion uint64
	OccurredAt     time.Time
}

type ViewingSession struct {
	SessionID    string
	SubjectID    string
	StartedAt    time.Time
	EndedAt      time.Time
	State        string
	LogicVersion string
	Revision     uint64
}

type Correction struct {
	CorrectionID string
	SessionID    string
	Kind         string
	Reason       string
	CreatedAt    time.Time
}
