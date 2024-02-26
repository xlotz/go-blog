package global

import (
	"github.com/sirupsen/logrus"
	"go-blog/server/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
