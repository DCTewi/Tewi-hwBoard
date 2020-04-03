package models

import (
	"github.com/dctewi/tewi-hwboard/core/database"
)

// HomeModel for homecontroller
type HomeModel struct {
	Model
	Tasks []database.TaskInfo
}
