package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
)

type FileService struct {
	host host.Host
	dest peer.ID
}

func NewFileService(h host.Host, dest peer.ID) *FileService {
	h.SetStreamHandler(Protocol, streamHandler)

	fmt.Println("File peer is ready")
	fmt.Println("libp2p-peer addresses:")
	for _, a := range h.Addrs() {
		fmt.Printf("%s/ipfs/%s\n", a, peer.Encode(h.ID()))
	}

	return &FileService{
		host: h,
		dest: dest,
	}
}

func (p *FileService) Serve() {
	stdReader := bufio.NewReader(os.Stdin)
	stream, err := p.host.NewStream(context.Background(), p.dest, Protocol)
	w := bufio.NewWriter(stream)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer stream.Close()
	r := bufio.NewReader(stream)
	go func(r *bufio.Reader) {
		for {
			str, _ := r.ReadString('\n')

			if str == "" {
				return
			}
			if str != "\n" {
				// Green console colour: 	\x1b[32m
				// Reset console colour: 	\x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
			}

		}
	}(r)

	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		w.WriteString(fmt.Sprintf("%s\n", strings.Trim(sendData, "\n")))
		w.Flush()
	}
}
