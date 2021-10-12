package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {

	client, err := ssh.Dial("tcp", "172.17.0.12:22", &ssh.ClientConfig{
		User:            "test",
		Auth:            []ssh.AuthMethod{ssh.Password("pass")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	}
	if err = session.RequestPty("linux", 32, 160, modes); err != nil {
		log.Fatalf("request pty error: %s", err.Error())
	}
	if err = session.Shell(); err != nil {
		log.Fatalf("start shell error: %s", err.Error())
	}
	if err = session.Wait(); err != nil {
		log.Fatalf("return error: %s", err.Error())
	}
}
