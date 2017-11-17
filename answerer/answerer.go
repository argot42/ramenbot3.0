package answerer

import (
    "net"
    "database/sql"
    "../util"
)

func Answerer (socket net.Conn, in chan util.Directive, config util.Configuration, db sql.DB) int {
}
