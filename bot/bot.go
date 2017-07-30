package bot

import "fmt"
import "../util"
import "../command"

func Init (config util.Configuration) {
    // load commands
    commands := command.Load_commands(config.Commands)
    // 
}
