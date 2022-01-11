package services

import (
    "blog/common"
    "gorm.io/gorm"
)

func db() *gorm.DB {
    return common.DB()
}
