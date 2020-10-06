package main

import (
        "flag"
        "fmt"
        "os"
        "io/ioutil"
        "time"
        "log"
        "golang.org/x/crypto/ssh"

        "github.com/pkg/sftp"
)

var (
        USER = flag.String("user", os.Getenv("USER"), "ssh username")
        HOST = flag.String("host", os.Getenv("HOST"), "ssh server hostname")
        PORT = flag.Int("port", 22, "ssh server port")
        PASS = flag.String("pass", os.Getenv("SOCKSIE_SSH_PASSWORD"), "ssh password")
        LOCALFILE = flag.String("localfile", "localhost", "localfile full path")
)

// func sendFile(ip, username, pass, port){
func sendFile(HOST, USER, PASS, LOCALFILE string, PORT int) (err error) {

        sshConfig := &ssh.ClientConfig{
                User: USER,
                Auth: []ssh.AuthMethod{ssh.Password(PASS)},
                HostKeyCallback: ssh.InsecureIgnoreHostKey(),
                Timeout: 30 * time.Second,
        }
        hostPort := fmt.Sprintf("%s:%d", HOST, PORT)

        connection, err := ssh.Dial("tcp", hostPort, sshConfig)
        if err != nil {
                return err
        }

        fmt.Println(">> SSH Connection Created!")

        // Creating an SFPT connection over SSH
        sftp, err := sftp.NewClient(connection)
        if err != nil {
                return err
        }
        defer sftp.Close()

        fmt.Println(">> SFTP Client Created!")

        // Uploading the file
        remoteFileName := "/home/subrata/newfile" // TODO: Make this name configurable

        // remoteFile, err := sftp.Create(remoteFileName)
        remoteFile, err := sftp.OpenFile(remoteFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)

        if err != nil {
                return err
        }

        fmt.Println(">> SFTP File Created!")

        srcFile, err := ioutil.ReadFile(LOCALFILE)

        if err != nil {
          return err
        }

//        fmt.Println(string(srcFile))

        if _, err := remoteFile.Write(srcFile); err != nil {
                return err
        }
        return
}

func main(){
        flag.Parse()

        fmt.Println(*USER)
        fmt.Println(*HOST)
        fmt.Println(*PORT)
        fmt.Println(*PASS)
        fmt.Println(*LOCALFILE)

        err := sendFile(*HOST, *USER, *PASS, *LOCALFILE, *PORT)

        if err != nil {
            log.Fatal(err)
        }
}
