package main

import (
        "github.com/tftp"
        "fmt"
        "os"
        "time"
)

func main(){
path := "/home/vagrant/send.go"

c, err := tftp.NewClient("192.168.50.14:8000")
wt, err := c.Receive("/home/vagrant/test.txt", "octet")
file, err := os.Create(path)
// Optionally obtain transfer size before actual data.
if n, ok := wt.(IncomingTransfer).Size(); ok {
	fmt.Printf("Transfer size: %d\n", n)
}
n, err := wt.WriteTo(file)
fmt.Printf("%d bytes received\n", n)

if err != nil {
 _ = err
 //fmt.Printf(err)
}
}
