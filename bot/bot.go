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
            return
            // join channels
            //join(socket, config)

            // start bot //
            // communication channel
            //msg_in := make(chan util.Msg)
            // spawn goroutines
            //go listener.Listener(socket, msg_in, config)
            //go timer.Timer(socket, msg_in, config)
            // execute answerer
            //response := answerer.Answerer(socket, msg_in, config)

            // check answerer return
            //if (response == util.Shutdown) { break }
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
    message,_ := bufio.NewReader(socket).ReadString('\n')
    fmt.Printf(string(message))
}

//func join (socket net.Conn, config util.Configuration) {
//}
