package handler

import (
	"fmt"
	"hash/crc32"

	"github.com/WuKongIM/WuKongIM/internal/eventbus"
)

func tail4(s string) string {
	if len(s) <= 4 {
		return s
	}
	return s[len(s)-4:]
}

// shortTraceID 统一短链路 ID：与业务侧按 uid+token 同算法，方便跨服务检索。
func shortTraceID(uid, token string) string {
	sum := crc32.ChecksumIEEE([]byte(uid + "|" + token))
	return fmt.Sprintf("u%s-%08x", tail4(uid), sum)
}

// connTraceID 连接级追踪：用于把 connect/connack/onSend 串到同一条连接生命周期。
func connTraceID(conn *eventbus.Conn) string {
	if conn == nil {
		return "conn-nil"
	}
	sum := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%s|%d|%d", conn.Uid, conn.NodeId, conn.ConnId)))
	return fmt.Sprintf("c%s-%08x", tail4(conn.Uid), sum)
}
