package main

import (
	"fmt"

	"cross-device-viewing-session/internal/session"
)

func main() {
	cfg := session.Config{DatabasePath: "data/sessions.db", DefaultGapSeconds: 180}
	fmt.Printf("会话合并服务配置：存储=%s，默认间隔=%d秒\n", cfg.DatabasePath, cfg.DefaultGapSeconds)
}
