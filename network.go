package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"

	ma "github.com/multiformats/go-multiaddr"
)

const Protocol = "/sharewire/0.0.1"

func streamHandler(stream network.Stream) {
	defer stream.Close()

	buf := bufio.NewReader(stream)
	w := bufio.NewWriter(stream)
	for {
		str, _ := buf.ReadString('\n')
		fmt.Printf("\x1b[32m%s\x1b[0m", str)
		s := strings.Trim(str, "\n")
		if s[:3] == DIR {
			dirList(w)
		} else if s[:2] == CD {
			chdir(w, s[3:])
		} else if s[:3] == PWD {
			pwd(w)
		} else if s[:4] == RETRIEVE {
			retrieve(w, s[5:])
		}
		w.Flush()
	}
}

func makeRandomHost(port int) host.Host {
	host, err := libp2p.New(libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port)))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return host
}

func addAddrToPeerstore(h host.Host, addr string) peer.ID {
	ipfsaddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		log.Fatalln(err.Error())
	}

	peerid, err := peer.Decode(pid)
	if err != nil {
		log.Fatalln(err.Error())
	}

	targetPeerAddr, _ := ma.NewMultiaddr(
		fmt.Sprintf("/ipfs/%s", peer.Encode(peerid)))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	h.Peerstore().AddAddr(peerid, targetAddr, peerstore.PermanentAddrTTL)
	return peerid
}
