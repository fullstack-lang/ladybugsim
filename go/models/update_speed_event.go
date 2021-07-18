package models

import (
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

// UpdateSpeedEvent is an event whose role is to
// ask the agant to compute its new speed
type UpdateSpeedEvent struct {
	gongsim_models.Event
}
