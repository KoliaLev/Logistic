package models

import "time"

type Transport struct {
	TransportType TransportType
	CreatedAt     time.Time
}

type TransportType struct {
	NameType     string
	Speed        int
	Volume       int
	Acceleration int
}
