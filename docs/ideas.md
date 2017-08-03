# Ideas for ramenbot3.0

## Basic structure
[goroutines]

+ Main
    1. Decode configuration
    2. Load commands
    3. Try to connect to server
        - If it fails retry x times with wait times
        - If it success continue with 4
    4. Register to server (first ping, register name, NickServ)
    5. Join channels
    6. Spawn **Listener** and **Timer**
    7. Execute **Answerer**
    8. Check **Answerer** return value
        - If 0 -> close program (shutdown)
        - If 1 -> goto 3 (restart)

+ Listener
    1. Listen for messages
        - If IRC command -> send response to server right away
        - If user command -> send command to **Answerer**
        - If EOF -> send restart to **Answerer**
        - If no msg -> goto 2
    2. Send check for triggers to **Answerer**
    3. Goto 1

+ Answerer
    1. Check channel
        - If no messages -> block
        - If messages
            + If user command -> execute and send answer to server
            + If trigger check -> check trigger, execute command associated to it and send answer to server
            + If restart -> return 0
            + If shutdown -> return 1

+ Timer
    1. Calculate time elapsed
    2. Check time triggers
        - If trigger or signal -> send trigger or signal associated response to **Answerer**
        - If no trigger -> goto 3
    3. Sleep 1s/.5s

## Msg definition (parsing)
* irc command - string
* sender - string 
* body - list:string

## Bot's configuration file
[.json file]

* host - string
* port - string
* nick - string
* channels - list:string
* prefix - string
* report_word - string
* havepass - boolean
* password - string
* ssl - boolean
* db - string
* commands - string

## Type of triggers
+ time
+ word
+ nickname
+ login, logout

## Command definition in Go
function signature -> func <command_name> (util.Emisor \[emisor info\], util.Receptor \[receptor info\], []string \[arguments\]) []string \[response\]

* com - function
* isSystem - boolean :: users only can execute command with this flag to false ::
* autoLoad - boolean :: execute command at start ::
* autoLoad_args - list :: arguments for the autoload initialization ::

## Command definition file
[.json file]

* command_name - string
* command_path - string
* isSystem - boolean
* autoLoad - boolean
* autoLoad_args - list
