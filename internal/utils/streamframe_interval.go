package utils

import "github.com/perseus101/quic-go/internal/protocol"

// ByteInterval is an interval from one ByteCount to the other
type ByteInterval struct {
	Start protocol.ByteCount
	End   protocol.ByteCount
}
