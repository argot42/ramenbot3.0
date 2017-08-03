package parser

import "strings"

type Msg struct {
    Ircom string
    Sender string
    Body []string
}

// stolen from twisted matrix :3
func Parse_msg(msg string) Msg {
    var prefix string
    var args []string
    var message Msg

    if (len(msg) == 0) { return message }

    if (msg[0] == ':') {
        splited := strings.SplitN(msg[1:], " ", 2)
        prefix = splited[0]
        msg = splited[1]
    }

    if (strings.Index(msg, " :") != -1) {
        splited := strings.SplitN(msg, " :", 2)
        args = strings.Split(splited[0], " ")
        args = append(args, splited[1])
    } else {
        args = strings.Split(msg, " ")
    }

    message.Ircom = args[0]
    message.Sender = prefix
    message.Body = args[1:]

    return message
}
