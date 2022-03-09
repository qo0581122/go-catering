package response

import "catering/model"

type IntegrationLog struct {
	*model.UserIntegrationLog
	LevelName string `json:"level_name"`
	Level     int    `json:"level"`
}

type Integration struct {
	*model.UserIntegration
	LevelName string `json:"level_name"`
	Level     int    `json:"level"`
}

type VipLevelLog struct {
	*model.UserVipLevelLog
	BeforeLevelName string `json:"before_level_name"`
	BeforeLevel     int    `json:"before_level"`
	AfterLevelName  string `json:"after_level_name"`
	AfterLevel      int    `json:"after_level"`
}
