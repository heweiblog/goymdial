package main

import (
	"github.com/jlaffaye/ftp"
	"log"
	"time"
)

func ftp_login(dst string) {
	//c, err := ftp.Dial(dst, ftp.DialWithTimeout(5*time.Second))
	c, err := ftp.Dial(dst, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	ftp_login("192.168.8.253:21")
	ftp_login("ftp.isc.org:21")
	ftp_login("192.168.8.121:21")
}
