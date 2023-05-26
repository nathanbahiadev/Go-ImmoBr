package models

import "gorm.io/gorm"

type RealState struct {
	gorm.Model
	Name string `json:"name"`
}

func Create(db *gorm.DB, name string) *RealState {
	realState := &RealState{Name: name}
	db.Create(realState)
	return realState
}
