package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"
	"sort"
	"github.com/cheggaaa/pb/v3"
)

func worker(ports, results chan uint, url string) {
	for p := range ports {
		address := fmt.Sprintf("%v:%d", url, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {

	url := flag.String("url", "127.0.0.1", "URL or IP on wich on want to run the test.")
	startPort := flag.Uint("startPort", 1, "Start scanning from this port")
	endPort := flag.Uint("endPort", 65535, "Makes  curl  verbose  during the operation.")
	workersCount := flag.Uint("workersCount", 1000, "Makes  curl  verbose  during the operation.")

	flag.Parse()

	if *startPort >= *endPort {
		fmt.Print("The starting port must be lower than the ending port.\n")
		os.Exit(0)
	}

	fmt.Printf("Url: %v\nstartPort: %v\nendPort: %v\nworkersCount: %v\n", *url, *startPort, *endPort, *workersCount)

	ports := make(chan uint, *workersCount)
	results := make(chan uint)
	var openports []uint


	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, *url)
	}

	go func() {
		for i := *startPort; i <= *endPort; i++ {
			ports <- i
		}
	}()

	diff := int(math.Abs(float64(*endPort) - float64(*startPort)))
	bar := pb.StartNew(diff)

	for i := 0; i < diff; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
		bar.Increment()
	}

	close(ports)
	close(results)

	bar.Finish()

	sort.Slice(openports, func(i, j int) bool {
		return openports[i] < openports[j]
	})

	for _, port := range openports {
		fmt.Printf("The port %d is open !\n", port)
	}
}
