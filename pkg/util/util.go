package util

import "go-gin-backend-admin/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = string(setting.AppSetting.JwtSecret)
}
