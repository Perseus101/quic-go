package congestion

import "github.com/perseus101/quic-go/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}
