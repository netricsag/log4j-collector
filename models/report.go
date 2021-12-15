package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	ServerName      string           `json:"serverName"`
	VulnerableFiles []VulnerableFile `json:"vulnerableFiles"`
	// VulnerableFiles []string `json:"vulnerableFiles"`
}

type VulnerableFile struct {
	gorm.Model
	ReportID uint
	FileName string `json:"fileName"`
}
