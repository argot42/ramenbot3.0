package listener

import (
    "fmt"
    "net"
)

func Pong (socket net.Conn, txt string) {
    fmt.Fprintf(socket, "PONG :%s\r\n", txt)
}
