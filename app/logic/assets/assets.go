package assets

import (
	"ControlCenter/model/mongoModel"
)

type Assets interface {
	Get() (mongoModel.Model, error)
	Add(assets mongoModel.Model) error
	Remove() error
	Edit(assets mongoModel.Model) error
	checkAuthority(authorityType int) bool
}

type DefaultAssets struct{}

var _ Assets = (*DefaultAssets)(nil)

func (d *DefaultAssets) Get() (mongoModel.Model, error) {
	return &mongoModel.DefaultModel{}, nil
}

func (d *DefaultAssets) Add(assets mongoModel.Model) error {
	return nil
}

func (d *DefaultAssets) Edit(assets mongoModel.Model) error {
	return nil
}

func (d *DefaultAssets) Remove() error {
	return nil
}

func (d *DefaultAssets) checkAuthority(authorityType int) bool {
	return true
}
