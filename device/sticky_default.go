//go:build !linux

package device

import (
	"github.com/RostamVPN/wireguard/conn"
	"github.com/RostamVPN/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
