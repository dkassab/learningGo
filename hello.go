package main

import (
    "fmt"
    "os"
    "github.com/olekukonko/tablewriter"
    )

func main() {
    fmt.Printf("hello, world\n")
    fmt.Println("مرحبا من ديمة!")

    data:=[][]string{
		[]string{"dima","go","4"},
		[]string{"junaid","go","0"},
    }
    table:=tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Name","PL", "Level"})
    table.AppendBulk(data)
    table.Render()
}
