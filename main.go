package main

import (
	"flag"
	"fmt"
)

const help = `
This is a peer to peer stripped down FTP implementation
Usage: Start file peer first with:   ./sharewire
       Then start the client peer with: ./sharewire -d <remote-peer-multiaddress>
Then you can type something into the terminal like: > DIR.
This runs the FTP command on the file peer and sends the response back.`

func main() {
	flag.Usage = func() {
		fmt.Println(help)
		flag.PrintDefaults()
	}

	destPeer := flag.String("d", "", "destination peer address")
	port := flag.Int("l", 5000, "libp2p listen port")
	flag.Parse()

	if *destPeer != "" {
		host := makeRandomHost(*port + 1)
		destPeerID := addAddrToPeerstore(host, *destPeer)
		fservice := NewFileService(host, destPeerID)
		fservice.Serve()
	} else {
		host := makeRandomHost(*port)
		_ = NewFileService(host, "")
		<-make(chan struct{})
	}
}
