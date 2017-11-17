package trigger

import (
    "testing"
    "fmt"
    "time"
    "./state"
    "./definition"
)

func TestCheckTime (t *testing.T) {
    tman := Trigger_manager{
        // User trigger
        []state.Trigger_state{
            state.Trigger_state{ 1, definition.Trigger{ 0, "", 1, time.Second * 2 }, 1, time.Now() },
            state.Trigger_state{ 2, definition.Trigger{ 0, "", 1, time.Second * 3 }, 1, time.Now() },
        },
        // System trigger
        //[]state.Trigger_state{
        //}
        nil,
    }

    fmt.Println(tman.Check_time())
    time.Sleep(2 * time.Second)
    fmt.Println(tman.Check_time())
    time.Sleep(1 * time.Second)
    fmt.Println(tman.Check_time())
}
