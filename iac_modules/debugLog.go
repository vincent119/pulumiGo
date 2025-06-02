package iac_modules

import (
    "log"
)

// Debug mode flag
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