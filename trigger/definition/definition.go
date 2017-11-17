package definition

import "time"

type Trigger struct {
    Ttype uint8
    Pattern string
    Uses int
    Time time.Duration
}
