package grpc

import (
	"medium/structs"
	"time"

	structsAPI "medium/structs/api"
)

type (
	// TypeHeader ...
	TypeHeader struct {
		NewRequest     bool
		RequestService string
		XDomain        string
	}

	// TypeHeaderRPC ...
	TypeHeaderRPC struct {
		ReqID       string                `json:"x-request-id"`
		Date        time.Time             `json:"date"`
		ContentType string                `json:"content-type"`
		RoundTrip   string                `json:"x-roundtrip"`
		Error       structs.TypeGRPCError `json:"error"`
	}

	// TypeResponseRPC ...
	TypeResponseRPC struct {
		Header   []byte
		Body     []byte
		Metadata structsAPI.HeaderTracer
	}
)
