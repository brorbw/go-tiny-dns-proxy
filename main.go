package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/miekg/dns"
	"gopkg.in/yaml.v2"
)

type Config struct {
	UpstreamServers []string `yaml:"upstream_servers"`
}

var upstreamServers []string
var serverIndex = 0
var mu sync.Mutex

// Load the upstream servers from the YAML config file
func loadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return fmt.Errorf("failed to parse config file: %v", err)
	}

	if len(config.UpstreamServers) == 0 {
		return fmt.Errorf("no upstream servers configured")
	}

	upstreamServers = config.UpstreamServers
	return nil
}

// getNextUpstreamServer returns the next DNS server in round robin fashion
func getNextUpstreamServer() string {
	mu.Lock()
	defer mu.Unlock()

	upstreamServer := upstreamServers[serverIndex]
	// Move to the next server
	serverIndex = (serverIndex + 1) % len(upstreamServers)
	return upstreamServer
}

// handleDNSRequest processes an incoming DNS query and forwards it to the next available upstream server
func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	// Get the next upstream server
	upstream := getNextUpstreamServer()

	// Forward the request to the upstream DNS server
	client := new(dns.Client)
	response, _, err := client.Exchange(r, upstream)
	if err != nil {
		log.Printf("Failed to forward query to %s: %v\n", upstream, err)
		return
	}
	log.Printf("Forwarding response from %s to %s", upstream, w.RemoteAddr())

	// Send the response back to the original client
	w.WriteMsg(response)
}

// startDNSServer starts the DNS server and listens for incoming queries
func startDNSServer() {
	// Create a new DNS server
	server := &dns.Server{
		Addr: ":53", // Listen on port 53 for DNS requests
		Net:  "udp", // UDP protocol
	}

	// Set the handler for DNS queries
	dns.HandleFunc(".", handleDNSRequest)

	// Start the DNS server
	log.Println("Starting DNS proxy server on port 53...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func main() {
	// Command-line argument for the config file path
	configFile := flag.String("config", "config.yaml", "Path to the YAML config file")
	flag.Parse()

	// Load the config file
	err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Start the DNS server
	startDNSServer()
}
