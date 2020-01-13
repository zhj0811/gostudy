
package main

import (
	"os"
	"time"
	"context"
	"log"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/alecthomas/kingpin.v2"
	"google.golang.org/grpc/connectivity"
)

const (
	defaultTimeout = 3 * time.Second
	maxRecvMsgSize = 100 * 1024 * 1024
	maxSendMsgSize = 100 * 1024 * 1024
)

var (
	versions = "1.0.0"
	app = kingpin.New("grpc-validate", "Utility for generating Hyperledger Fabric key material")

	con     = app.Command("connect", "Connect remote port with tls")
	ver     = app.Command("version", "Show the version of grpc-validate")
	add     = con.Flag("address", "Remote peer's address with port").String()
	tls     = con.Flag("tls", "Whether tls is applicable").Bool()
	tlsfile = con.Flag("tlsfile", "The tls file to use").String()
	checkTime    = con.Flag("checkTime", "The internal time to check the connection's state").Int()
)

func main() {
	kingpin.Version("0.0.1")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	// "generate" command
	case con.FullCommand():
		connect()
	case ver.FullCommand():
		version()
	}
}

func connect() {
	var opts []grpc.DialOption
	var color int

	if *tls {
		creds, err := credentials.NewClientTLSFromFile(*tlsfile, "")
		if err != nil {
			log.Printf("Load tls cert failed: %s", err.Error())
			return
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithTimeout(defaultTimeout))
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxRecvMsgSize),
		grpc.MaxCallSendMsgSize(maxSendMsgSize)))
	conn, err := grpc.Dial(*add, opts...)
	if err != nil {
		log.Printf("Grpc dial %s failed: %s", *add, err.Error())
		return
	}
	log.Println("Connection has been established ...")
	if conn.GetState() == connectivity.Ready {
		color = 32
	}else  {
		color = 31
	}
	log.Printf("The first connection status is %c[1;40;%dm%s%c[0m\n\n", 0x1B, color, conn.GetState(), 0x1B)

	go func() {
		log.Println("Check status change: Detect connection status")
		for {
			iswait := conn.WaitForStateChange(context.Background(), conn.GetState())
			if conn.GetState() == connectivity.Ready {
				color = 32
			}else  {
				color = 31
			}
			if iswait {
				log.Printf("Check status change: Connection status changed: %c[1;40;%dm%s%c[0m\n\n", 0x1B, color, conn.GetState(), 0x1B)
			} else {
				log.Println("Check status change: The connection status has not changed, perhaps the context expires")
			}
		}

	}()

	for {
		time.Sleep(time.Duration(*checkTime) * time.Second)
		if conn.GetState() == connectivity.Ready {
			color = 32
		}else  {
			color = 31
		}
		log.Printf("Check current status: Current connection status: %c[1;40;%dm%s%c[0m\n\n", 0x1B, color, conn.GetState(), 0x1B)
	}
}

func version() {
	fmt.Printf("grpc-validate version is %s\n", versions)
}
