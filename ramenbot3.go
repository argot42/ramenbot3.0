package main

import (
    "os"
    "log"
    "encoding/json"
    "io/ioutil"
    "./util"
    "./bot"
)

func main() {
    // Check arguments
    if (len(os.Args) < 2) {
        util.Usage()
        return
    }

    // read configuration
    info,err := ioutil.ReadFile(os.Args[1])
    if (err != nil) {
        log.Println(err)
        os.Exit(1)
    }
    // decode json
    var config util.Configuration
    if err := json.Unmarshal(info, &config); err != nil {
        log.Println(err)
        os.Exit(1)
    }

    // start bot
    bot.Init(config)
}
