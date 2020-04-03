package models

import (
	"github.com/dctewi/tewi-hwboard/core/database"
)

// HistoryModel for homecontroller
type HistoryModel struct {
	Model
	Historys []database.UploadLog
}
