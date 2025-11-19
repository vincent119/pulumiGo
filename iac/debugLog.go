// Package iac 提供基礎設施即代碼(IaC)相關的模組與工具函式。
// 包含調試日誌功能、堆疊檢查等實用工具。
// 此 package 被 main.go 及其他模組使用,以增強 IaC 管理功能。
package iac

import (
	"log"
)

// DebugMode 指示是否啟用調試日誌
var DebugMode bool = false

// SetDebugMode enables or disables debug logging
func SetDebugMode(enabled bool) {
    DebugMode = enabled
    if enabled {
        log.Printf("Debug mode enabled")
    }
}

// DebugLog logs messages only when debug mode is enabled
func DebugLog(format string, v ...interface{}) {
    if DebugMode {
        log.Printf("[DEBUG] "+format, v...)
    }
}
