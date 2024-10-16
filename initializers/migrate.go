package initializers

import (
	"github.com/raunak173/go-todo/models"
)

func SyncDB() {
	Db.AutoMigrate(&models.Task{})
	Db.AutoMigrate(&models.User{})
}
