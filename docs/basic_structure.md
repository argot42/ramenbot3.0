# INFO
## Basic steps

1. Startup
    - Load commands
    - Spawn threads for timer, communication and execution
    - Set system triggers (timeout)
2. Connect to server
    - Join channels
3. Listen
4. Exit

## Basic structure
+ Timer
    1. Calculate time elapsed
    2. Check time for time triggers
    3. Send time trigger command to execution stack
    4. Check other triggers
    5. Send other trigger command to execution stack

+ Communication
    1. Receive message
    2. Parse message
        - If IRC command -> answer right away
        - If user command -> send command to execution
    3. Send message info to timer
    4. check exit
    5. check responses
        - Send responses back to server

+ Execution
    1. Pop commands list
    2. Pop trigger command list
    3. Execute trigger command
        - Send result to communication
    4. Execute user command
        - Send result to communication

## Type of triggers
+ time
+ word
+ nickname
+ login, logout

## Command definition in Go
* command - function
* arguments - list
* isSystem - boolean :: users only can execute command with this flag to false ::
* autoLoad - boolean :: execute command at start ::
* autoLoad_args - list :: arguemnts for the autoload initialization ::

## Command definition file
