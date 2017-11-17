package trigger

import (
    "time"
    "./state"
)

// trigger types
const (
    Time_type = iota
    Nick_type
    Word_type
    Join_type
    Quit_part_type
)


// Trigger manager
type Trigger_manager struct {
    User_t []state.Trigger_state
    Sys_t []state.Trigger_state
}

func (tm *Trigger_manager) New_trigger() {
}

func (tm *Trigger_manager) Check_time() []state.Trigger_state {
    var triggered []state.Trigger_state
    // check user triggers
    for _,tstate := range tm.User_t {
        if (tstate.Tr.Ttype != Time_type) { continue }

        if (time.Since(tstate.Start_time) >= tstate.Tr.Time) {
            triggered = append(triggered, tstate)

            remaining := tstate.Update_counter(1)
            if (remaining == 0) {
                // if counter gets to zero, remove trigger
                tm.rm(tstate.Id)
            } else {
                // else reset time
                tstate.Reset()
            }
        }
    }

    // check system triggers
    for _,tstate := range tm.Sys_t {
        if (tstate.Tr.Ttype != Time_type) { continue }

        if (time.Since(tstate.Start_time) >= tstate.Tr.Time) {
            triggered = append(triggered, tstate)

            remaining := tstate.Update_counter(1)
            if (remaining == 0) {
                // if counter gets to zero, remove trigger
                tm.rm(tstate.Id)
            } else {
                // else reset time
                tstate.Reset()
            }
        }
    }

    return triggered
}

func check_time_aux (trigger_states []state.Trigger_state, triggered *[]state.Trigger_state) {
    for _,tstate := range trigger_states {
        if (tstate.Tr.Ttype != Time_type) { continue }

        if (time.Since(tstate.Start_time) >= tstate.Tr.Time) {
            *triggered = append(*triggered, tstate)

            remaining := tstate.Update_counter(1)
            if (remaining == 0) {
                // if counter gets to zero, remove trigger
                tm.rm(tstate.Id)
            } else {
                // else reset time
                tstate.Reset()
            }
        }
    }
}

func (tm *Trigger_manager) rm (id int) {
    return
}
