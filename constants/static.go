package constants

import (
	"time"
)

var TimeZone, _ = time.LoadLocation("Asia/Taipei")

// RedisExpireTime 5 days
var RedisExpireTime = time.Second * 432000
