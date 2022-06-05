package godate

import "time"

type (
	Date1 struct {
		time.Time
	}

	Date2 struct {
		datetime time.Time
	}

	Date3 struct {
		*time.Time
	}

	Date4 struct {
		datetime *time.Time
	}

	Date5 time.Time
	Date6 *time.Time
)
