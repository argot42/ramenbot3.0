package bot

import (
    "fmt"
    "net"
    "log"
    "time"
    "bufio"
    "crypto/tls"
    "../util"
    "../command"
    "../parser"
    "../listener"
)

func Init (config util.Configuration) {
    // load commands
    commands := command.Load_commands(config.Commands)
    fmt.Println(commands)

    // main loop
    retry := 1
    for retry <= util.MaxRetry {
        // try to connect
        fmt.Println("Connecting...")
        socket,err := connect(config)
        defer socket.Close()

        if (err != nil) {
            log.Println(err)

        } else {
            // reset retry counter
            retry = 1

            // register (first ping, name, NickServ)
            register(socket, config)
            // join channels
            join(socket, config)
            time.Sleep(time.Second * 2)

            // start bot //
            // communication channel
            msg_in := make(chan util.Msg)
            // spawn goroutines
            go listener.Listener(socket, msg_in, config)
            go timer.Timer(socket, msg_in, config)
            // execute answerer
            response := answerer.Answerer(socket, msg_in, config)

            // check answerer return
            if (response == util.Shutdown) { break }
        }

        // try to reconnect
        fmt.Printf("Reconnecting in %ds...", util.RetryDelay * retry)
        // sleep
        time.Sleep(time.Second * time.Duration(util.RetryDelay) * time.Duration(retry))
        retry++
    }

    fmt.Println("Goodbye ❤")
}

func connect (config util.Configuration) (net.Conn, error) {
    if (config.SSL) { return tls.Dial("tcp", config.Host + ":" + config.Port, nil) }

    return net.Dial("tcp", config.Host + ":" + config.Port)
}

func register (socket net.Conn, config util.Configuration) {
    // register nick and user
    time.Sleep(time.Millisecond * 5)
    fmt.Fprintf(socket, "USER %s %s %s :howdy\r\n", config.Nick, config.Nick, config.Nick)
    time.Sleep(time.Millisecond * 5)
    fmt.Fprintf(socket, "NICK %s\r\n", config.Nick)

    reader := bufio.NewReader(socket)
    for {
        message,_ := reader.ReadString('\n')
        fmt.Printf(message)

        // parse msg and check if it's ping
        parsed_msg := parser.Parse_msg(message)
        if (parsed_msg.Ircom == "PING") {
            // answer ping
            listener.Pong(socket, parsed_msg.Body[0])
            time.Sleep(time.Millisecond * 5)

            // register with NickServ
            if (config.Havepass) { fmt.Fprintf(socket, "PRIVMSG NickServ :IDENTIFY %s\r\n", config.Password) }

            return
        }
    }
}

func join (socket net.Conn, config util.Configuration) {
    for _,channel := range config.Channels {
        fmt.Fprintf(socket, "JOIN %s\r\n", channel)
        time.Sleep(time.Millisecond * 5)
    }
}
