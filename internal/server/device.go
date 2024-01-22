package server

import (
	"fmt"
	"github.com/disism/karma/ent/devices"
	"github.com/disism/karma/ent/users"
	"net/http"
	"strconv"
	"time"

	"github.com/disism/karma/ent"
	"github.com/gin-gonic/gin"
)

type Device struct {
	ID         string    `json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	IP         string    `json:"ip"`
	Device     string    `json:"device"`
}

const (
	DeviceNotFound = "device not found"
)

func (s *Server) GetDevices() error {
	defer s.client.Close()

	all, err := s.client.Devices.
		Query().
		Where(
			devices.
				HasUserWith(
					users.IDEQ(
						GetUserID(s.ctx),
					),
				),
		).All(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNotFound(s.ctx, DeviceNotFound)
			return nil
		}
		return fmt.Errorf("query devices: %w", err)
	}
	r := make([]Device, len(all))
	for i, v := range all {
		r[i] = Device{
			ID:         strconv.FormatUint(v.ID, 10),
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			IP:         v.IP,
			Device:     v.Device,
		}
	}

	Success(s.ctx, gin.H{
		"code":    http.StatusOK,
		"devices": r,
	})
	return nil
}

func (s *Server) DeleteDevice() error {
	defer s.client.Close()

	id, err := strconv.ParseUint(s.ctx.Param("id"), 10, 64)
	if err != nil {
		return fmt.Errorf("parse device id: %w", err)
	}

	exist, err := s.client.Devices.
		Query().
		Where(
			devices.IDEQ(id),
			devices.HasUserWith(
				users.IDEQ(
					GetUserID(s.ctx),
				),
			),
		).Only(s.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			ErrorNotFound(s.ctx, DeviceNotFound)
			return nil
		}
		return fmt.Errorf("exist device: %w", err)
	}

	if err := s.client.Devices.DeleteOne(exist).Exec(s.ctx); err != nil {
		return fmt.Errorf("delete device error: %w", err)
	}

	Success(s.ctx, gin.H{
		"code": http.StatusOK,
	})
	return nil
}
