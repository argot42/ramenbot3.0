# Ideas for ramenbot3.0

## Basic structure
[goroutines]

+ Main
    1. Decode configuration
    2. Load commands
    3. Connect to server
        - If it fails retry x times with increasing wait time
        - If it success continue with 4
    4. Spawn **Comunication**, **Execution** and **Timer**
    5. Block until wake up
    6. Check messages from routines
        - If restart message go to 3
        - If shutdown message end program

+ Communication
    1. Register
    2. Join channels
    3. Listen for messages
        - If message
            + If server shutdown send restart to **Execution**
            + Send check for command to **Execution**
            + Send check for trigger to **Execution**
            + If IRC command answer right away
        - If no message 
            + Continue with 4
    4. Check responses
        - If type shutdown kill routine
        - If type restart kill routine
        - If type response send response back to server
    5. Sleep?

+ Execution
    1. Check messages
        - If no messages -> block
        - If messages
            + If restart send restart answer to **Communication**, **Timer** and back to **Main** routine
            + If shutdown send shutdown answer to **Communication**, **Timer** and back to **Main** routine
            + If command send answer to **Communication**
            + If trigger send answer to **Communication**

+ Timer
    1. Calcute time elapsed
    2. Check time triggers
        - If time trigger
            + Send trigger command to **Execution**
        - If no time trigger
            + Continue to 3 
    3. Check answers
        - If type shutdown kill routine
        - If type restart kill routine
    4. Sleep

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
