package main

import "../util"
import "fmt"
import "strings"

func Test(em util.Emisor, re util.Receptor, args []string) []string{
    var res []string
    res = append(res, strings.ToUpper(args[0]))
    fmt.Println(em, re)

    return res
}
