//go:build !linux

package device

import (
	"github.com/RostamVPN/RostamWG/conn"
	"github.com/RostamVPN/RostamWG/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
