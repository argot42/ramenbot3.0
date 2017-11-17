package listener

import (
    "fmt"
    "net"
    "bufio"
    "../parser"
    "../util"
)

func Pong (socket net.Conn, txt string) {
    fmt.Fprintf(socket, "PONG :%s\r\n", txt)
}

func Listener (socket net.Conn, in chan util.Directive, config util.Configuration) {
    reader := bufio.NewReader(socket)
    for {
        message,err := reader.ReadString('\n')
        if (err != nil) { // lost connection to server restart
            in<- Directive{ Comtype: 1, Sender: "", Receiver: "", Com: nil }
        }

        parsed_msg := parser.Parse_msg(message)
        switch parsed_msg.Ircom {
        case "PRIVMSG":
            // send check trigger
            // parse usr comm
            // if user comm send ucom
        case "JOIN":
            // send check trigger
        case "PART","QUIT":
            // send check trigger
        default:
            continue
        }
    }
}
