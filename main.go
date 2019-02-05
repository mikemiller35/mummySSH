package main

import (
		"io/ioutil"
		"net"
		"golang.org/x/crypto/ssh"
		"github.com/helloyi/go-sshclient"
       )

func main() {
	//Need to get the following values from my https request
	client, err := DialWithKey("host:port", "username", "pem")
		if err != nil {
			handleErr(err)
		}
	defer client.Close()

		// default terminal
		if err := client.Terminal(nil).Start(); err != nil {
			handleErr(err)
		}

	// terminal config
config := &sshclient.TerminalConfig {
Term: "xterm",
	      Height: 50,
	      Weight: 120,
	      Modes: ssh.TerminalModes {
		      ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		      ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	      },
	}
	if err := client.Terminal(config).Start(); err != nil {
		handleErr(err)
	}
}
