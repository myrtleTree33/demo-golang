package models

import "github.com/segmentio/ksuid"

type Car struct {
	Id    ksuid.KSUID `json:"id" db:"id"`
	Name  string      `json:"name" db:"name"`
	Color string      `json:"color" db:"color"`
}
