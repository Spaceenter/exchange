package testutil

import (
	"time"

	"github.com/golang/protobuf/ptypes"
)

var (
	TimeZero          = time.Time{}
	TimeEarly         = time.Unix(100, 0)
	TimeNow           = time.Unix(200, 0)
	TimeLate          = time.Unix(200, 0)
	ProtoTimeZero, _  = ptypes.TimestampProto(TimeZero)
	ProtoTimeEarly, _ = ptypes.TimestampProto(TimeEarly)
	ProtoTimeNow, _   = ptypes.TimestampProto(TimeNow)
	ProtoTimeLate, _  = ptypes.TimestampProto(TimeLate)
)
