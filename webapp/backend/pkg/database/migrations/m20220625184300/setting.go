package m20220625184300

import (
	"gorm.io/gorm"
)

type Setting struct {
	//GORM attributes, see: http://gorm.io/docs/conventions.html
	gorm.Model

	SettingKeyName  string `json:"setting_key_name"`
	SettingDataType string `json:"setting_data_type"`

	SettingValueNumeric int64  `json:"setting_value_numeric"`
	SettingValueString  string `json:"setting_value_string"`
}
