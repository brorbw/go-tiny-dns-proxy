# DNS Proxy Server (Round-Robin Forwarding)

This Go-based DNS Proxy Server forwards incoming DNS requests to a configurable list of upstream DNS servers using a round-robin algorithm. It helps in distributing DNS queries across multiple DNS servers to improve redundancy, load balancing, and fault tolerance.

## Features

- **Round-Robin DNS forwarding:** Distributes DNS queries evenly across a configurable list of upstream DNS servers.
- **Customizable configuration:** Configurable via a YAML file or command-line argument.
- **UDP support:** Handles DNS queries over UDP, typically used for DNS resolution.
- **Timeout settings:** Allows for configurable timeouts for DNS queries to upstream servers.

## Requirements

- Go 1.18+ (or later)
- Access to a working network with DNS servers

## Installation

### Step 1: Clone the repository

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/dns-proxy.git
cd dns-proxy
```

### Step 2: Install dependencies

Run the following command to install required Go modules:

```bash
go mod tidy
```

### Step 3: Install the application

Build the application with the following command:

```bash
go build -o dns_proxy
```

This will generate an executable named `dns_proxy`.

## Configuration

The DNS proxy uses a YAML configuration file to specify the upstream DNS servers and the timeout duration for DNS queries. You can provide the path to the YAML configuration file using a command-line argument.

### Sample configuration (`config.yaml`)

```yaml
upstream_servers:
  - "8.8.8.8:53"  # Google DNS
  - "1.1.1.1:53"  # Cloudflare DNS
  - "9.9.9.9:53"  # Quad9 DNS

timeout: 5  # Timeout in seconds for upstream DNS queries
```

### Configuration File Options:

- `upstream_servers`: A list of DNS servers to which the proxy will forward DNS queries. The proxy will use these servers in a round-robin fashion.
- `timeout`: The timeout (in seconds) for DNS queries to upstream servers. This helps control how long the proxy will wait for a response from the upstream server before giving up and trying the next server in the round-robin list.

You can modify this file to add/remove DNS servers or adjust the timeout as needed.

## Usage

### Running the DNS Proxy Server

You can run the DNS Proxy server by specifying the path to your configuration file using the `-config` flag.

```bash
sudo ./dns_proxy -config /path/to/config.yaml
```

The DNS proxy will listen on UDP port `53` by default, which is the standard DNS port. Running this application typically requires elevated privileges (e.g., using `sudo`) because it binds to a privileged port.

### Testing

Once the DNS Proxy server is running, you can test it by sending DNS queries using a tool like `dig` or `nslookup`.

#### Using `dig`:

```bash
dig @127.0.0.1 google.com
```
### Command-Line Arguments

- `-config`: Path to the YAML configuration file containing the list of upstream DNS servers and the timeout.

## Troubleshooting

- **Firewall issues:** Ensure that UDP traffic on port 53 is allowed in your firewall.
- **Permission issues:** If you encounter permission errors when binding to port 53, make sure to run the program with elevated privileges (e.g., using `sudo`).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [miekg/dns](https://github.com/miekg/dns): A DNS library for Go that helps handle DNS queries.
- [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2): A YAML package for Go used to read configuration files.
