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
file, err := os.Open(path)
c.SetTimeout(5 * time.Second) // optional
rf, err := c.Send("/home/vagrant/send.go", "octet")
n, err := rf.ReadFrom(file)
fmt.Printf("%d bytes sent\n", n)
if err != nil {
 _ = err
 //fmt.Printf(err)
}
}
