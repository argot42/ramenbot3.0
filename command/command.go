package command

import (
    "log"
    "os"
    "path/filepath"
    "io/ioutil"
    "plugin"
    "encoding/json"
)

/************************/
/****** COMMAND *********/
/************************/

// command descriptor
type descriptor struct {
    Command_name string
    Command_path string
    IsSystem bool
    AutoLoad bool
    AutoLoad_args []string
}

// command structure
type Command struct {
    Com plugin.Symbol
    IsSystem bool
    AutoLoad bool
    AutoLoad_args []string
}
// methods
func (c Command) Execute (em string, re string, args []string) []string {
    return c.Com.(func(string, string, []string) []string) (em, re, args)
}
/***********************/


/***********************/
/*** COMMAND MANAGER ***/
/***********************/
type CommMan struct {
    Commands map[string]Command
}
// methods
func (c CommMan) Lookup (name string, system bool) *Command {
    com,ok := c.Commands[name]

    // command exist
    if (!ok) { return nil }
    // don't allow users to execute system commands
    if (!system && com.IsSystem) { return nil }

    return &com
}
// get commands
func (c CommMan) GetCommands () {
    return c.Commands
}
/***********************/


// load commands
func Load_commands (path string) CommMan {
    // get filepaths of command descriptors
    files,err := filepath.Glob(path + "/*.json")
    if (err != nil) {
        log.Println(err)
        os.Exit(1)
    }

    var commands map[string]Command
    commands = make(map[string]Command)
    for _,f := range files {
        // read json file
        content,err := ioutil.ReadFile(f)
        if (err != nil) {
            log.Println(err)
            os.Exit(1)
        }

        Load(content, &commands)
    }

    //return commands
    return CommMan{commands}
}

// decode json and build command
func Load (content []byte, commands *map[string]Command) {
    // decode json
    var command_descriptor descriptor
    if err := json.Unmarshal(content, &command_descriptor); err != nil {
        log.Println(err)
        os.Exit(1)
    }

    // build command
    var new_command *Command = Build_command(command_descriptor)
    if (new_command == nil) {
        log.Printf("Command %s: not loaded!", command_descriptor.Command_name)
    } else {
        (*commands)[command_descriptor.Command_name] = *new_command
    }
}

func Build_command (command_descriptor descriptor) *Command {
    var new_command Command

    // load plugin
    plug,err := plugin.Open(command_descriptor.Command_path)
    if (err != nil) {
        log.Println(err)
        return nil
    }

    // lookup command
    new_command.Com,err = plug.Lookup(command_descriptor.Command_name)
    if (err != nil) {
        log.Println(err)
        return nil
    }
    // copy the rest of the information
    new_command.IsSystem = command_descriptor.IsSystem
    new_command.AutoLoad = command_descriptor.AutoLoad
    new_command.AutoLoad_args = command_descriptor.AutoLoad_args

    return &new_command
}
