package util

import "os"
import "fmt"

// version
const Version = "3.0"

// config structure
type Configuration struct {
    Host string
    Port int
    Nick string
    Channels []string
    Prefix string
    Report_word string
    Havepass bool
    Password string
    SSL bool
    DB string
    Commands string
}

// usage msg
func Usage() {
    fmt.Printf("ramenbot v%s\n", Version)
    fmt.Printf("usage: %s <config.json>\n", os.Args[0])
}

// emisor - receptor structs
/*** have to change to reflect real data ***/
type Emisor struct {
    Handle string
}
type Receptor struct {
    Handle string
}
