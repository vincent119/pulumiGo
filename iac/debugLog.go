// Package iac 提供基礎設施即代碼(IaC)相關的模組與工具函式。
// 包含調試日誌功能、堆疊檢查等實用工具。
// 此 package 被 main.go 及其他模組使用,以增強 IaC 管理功能。
package iac

import (
	"fmt"

	"github.com/vincent119/zlogger"
)

func init() {
	// Initialize at warn level to suppress the internal "logger initialized" info message.
	// Level will be raised to debug via SetDebugMode when --debug flag is set.
	zlogger.Init(&zlogger.Config{Level: "warn"})
}

// SetDebugMode enables or disables debug-level logging.
func SetDebugMode(enabled bool) {
	if enabled {
		zlogger.SetLevel("debug")
		zlogger.Debug("debug mode enabled")
	} else {
		zlogger.SetLevel("info")
	}
}

// DebugLog emits a formatted message at debug level.
// It is a no-op when the current log level is above debug.
func DebugLog(format string, v ...interface{}) {
	zlogger.Debug(fmt.Sprintf(format, v...))
}
