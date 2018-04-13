package testutil

import (
	"time"

	"github.com/golang/protobuf/ptypes"
)

var (
	timeZero          = time.Time{}
	timeEarly         = time.Unix(100, 0)
	timeNow           = time.Unix(200, 0)
	timeLate          = time.Unix(200, 0)
	protoTimeZero, _  = ptypes.TimestampProto(timeZero)
	protoTimeEarly, _ = ptypes.TimestampProto(timeEarly)
	protoTimeNow, _   = ptypes.TimestampProto(timeNow)
	protoTimeLate, _  = ptypes.TimestampProto(timeLate)
)
