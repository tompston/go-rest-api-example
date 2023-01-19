package user_module

import (
	"time"

	"github.com/google/uuid"
)

// Custom Response for endpoint
type GetUserDetailsWithAuthResponse struct {
	UserID          uuid.UUID     `json:"user_id"`
	CreatedAt       time.Time     `json:"created_at"`
	Username        string        `json:"username"`
	Balance         int64         `json:"balance"`
	TimeUntillBonus time.Duration `json:"time_untill_bonus"`
}
