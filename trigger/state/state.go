package state

import (
    "time"
    "../definition"
)

type Trigger_state struct {
    Id int
    Tr definition.Trigger
    Remaining_uses int
    Start_time time.Time
}

// methods
func (ts *Trigger_state) Reset () {
    ts.Start_time = time.Now()
}
func (ts Trigger_state) Check_counter () int {
    return ts.Remaining_uses
}
func (ts *Trigger_state) Update_counter (n int) int {
    if (ts.Remaining_uses <= 0) { return ts.Remaining_uses }

    ts.Remaining_uses -= n
    if (ts.Remaining_uses < 0) { ts.Remaining_uses = 0 }

    return ts.Remaining_uses
}
