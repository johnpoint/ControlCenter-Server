package assets

import (
	"ControlCenter/model/mongomodel"
)

type Assets interface {
	Get() (mongomodel.Model, error)
	Add(assets mongomodel.Model) error
	Remove() error
	Edit(assets mongomodel.Model) error
	checkAuthority(authorityType int) bool
}

type DefaultAssets struct{}

var _ Assets = (*DefaultAssets)(nil)

func (d *DefaultAssets) Get() (mongomodel.Model, error) {
	return &mongomodel.DefaultModel{}, nil
}

func (d *DefaultAssets) Add(assets mongomodel.Model) error {
	return nil
}

func (d *DefaultAssets) Edit(assets mongomodel.Model) error {
	return nil
}

func (d *DefaultAssets) Remove() error {
	return nil
}

func (d *DefaultAssets) checkAuthority(authorityType int) bool {
	return true
}
