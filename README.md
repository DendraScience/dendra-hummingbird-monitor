# Dendra Hummingbird Monitor

A lightweight system monitoring agent for Linux servers that collects metrics and sends them to a cloud endpoint (Google BigQuery via Cloud Functions).

## Features

- **System Metrics**: Memory, CPU load, disk usage, network traffic
- **Package Management**: Tracks installed packages and available updates (supports apt, pacman, dnf, apk, and more via [snack](https://github.com/gogrlx/snack))
- **Container Monitoring**: Docker container stats (CPU, memory, uptime)
- **Kubernetes Support**: Optional control plane metrics collection
- **Low Overhead**: Minimal resource usage with configurable collection intervals

## Metrics Collected

| Category | Metrics |
|----------|---------|
| **System** | Hostname, uptime, kernel version |
| **Memory** | Total, available, free, cached, usage percentage |
| **CPU** | Processor count, load averages, CPU percentage |
| **Disk** | Free space, usage percentage per mount point |
| **Network** | Bytes up/down for WAN and LAN interfaces, IP addresses |
| **Packages** | Installed count, available updates |
| **Containers** | Image, name, CPU/memory utilization, uptime |

## Requirements

- Linux (amd64)
- Go 1.26+ (for building)
- Network access to the metrics endpoint
- Optional: Docker (for container metrics)

## Building

```bash
# Build for the current system
make

# Build with custom output filename
FILENAME=hummingbird make

# Cross-compile for Linux
GOOS=linux make
```

## Configuration

Create a configuration file at `/etc/dendra/hummingbird.toml`:

```toml
# Authentication key for the metrics endpoint
authkey = "your-secret-key"

# URL of the metrics ingestion endpoint
endpoint = "https://your-cloud-function-url"

# Network interface names
wan = "eth0"
lan = "eth1"

# Collection interval in minutes
sleeplooptime = 15
```

### Configuration Options

| Option | Description | Example |
|--------|-------------|---------|
| `authkey` | Secret key for authenticating with the endpoint | `"abc123"` |
| `endpoint` | URL where metrics are POSTed | `"https://us-central1-project.cloudfunctions.net/ingest"` |
| `wan` | WAN network interface name | `"eth0"`, `"ens3"` |
| `lan` | LAN network interface name | `"eth1"`, `"ens4"` |
| `sleeplooptime` | Minutes between metric collections | `15` |

## Deployment

### Manual Installation

1. Build the binary:
   ```bash
   GOOS=linux make
   ```

2. Copy to the target server:
   ```bash
   scp dendra-hummingbird-monitor user@server:/usr/local/bin/
   ```

3. Create the configuration directory and file:
   ```bash
   ssh user@server "sudo mkdir -p /etc/dendra"
   scp hummingbird.toml user@server:/tmp/
   ssh user@server "sudo mv /tmp/hummingbird.toml /etc/dendra/"
   ```

4. Install the systemd service:
   ```bash
   scp dendra-hummingbird-monitor.service user@server:/tmp/
   ssh user@server "sudo mv /tmp/dendra-hummingbird-monitor.service /etc/systemd/system/"
   ssh user@server "sudo systemctl daemon-reload"
   ssh user@server "sudo systemctl enable --now dendra-hummingbird-monitor"
   ```

### Systemd Service

The included `dendra-hummingbird-monitor.service` file runs the monitor as a system service:

```ini
[Unit]
Description=Dendra Hummingbird Monitor
After=network-online.target

[Service]
Type=simple
User=root
ExecStart=/usr/local/bin/dendra-hummingbird-monitor
Restart=always
RestartSec=30

[Install]
WantedBy=multi-user.target
```

### Service Management

```bash
# Start the service
sudo systemctl start dendra-hummingbird-monitor

# Stop the service
sudo systemctl stop dendra-hummingbird-monitor

# Check status
sudo systemctl status dendra-hummingbird-monitor

# View logs
sudo journalctl -u dendra-hummingbird-monitor -f
```

## Cloud Function Backend

The metrics are ingested by a Google Cloud Function that writes to BigQuery. See the `cloud_function/` directory for the backend implementation.

### Environment Variables

The Cloud Function requires:

- `HUMMINGBIRD_KEY`: Authentication key (must match client `authkey`)

### BigQuery Schemas

Schema definitions for the BigQuery tables are in `cloud_function/bq/`:

- `metrics_schema.json`: Main metrics table
- `containers_schema.json`: Container metrics
- `disk_schema.json`: Disk usage metrics

## Command Line Options

```bash
# Show version information and exit
./dendra-hummingbird-monitor --version

# Run as Kubernetes control plane (collects cluster-wide container metrics)
./dendra-hummingbird-monitor --isPrimary
```

## Architecture

```
┌─────────────────────┐         ┌─────────────────────┐
│   Linux Server      │         │   Google Cloud      │
│                     │         │                     │
│  ┌───────────────┐  │  POST   │  ┌───────────────┐  │
│  │  Hummingbird  │──┼────────►│  │Cloud Function │  │
│  │    Monitor    │  │         │  └───────┬───────┘  │
│  └───────────────┘  │         │          │          │
│                     │         │          ▼          │
│                     │         │  ┌───────────────┐  │
│                     │         │  │   BigQuery    │  │
│                     │         │  └───────────────┘  │
└─────────────────────┘         └─────────────────────┘
```

## Development

### Project Structure

```
.
├── bin/main.go          # Main entry point
├── config/              # Configuration loading
├── cpu/                 # CPU metrics
├── disk/                # Disk metrics
├── k8s/                 # Kubernetes/Docker integration
├── pkg/                 # Package manager metrics
├── proc/                # /proc filesystem parsing
├── publish/             # HTTP client for posting metrics
├── types/               # Shared data structures
├── cloud_function/      # GCP Cloud Function backend
│   └── bq/              # BigQuery integration
└── wg/                  # WireGuard metrics (if applicable)
```

### Running Locally

```bash
# Build and run (requires Linux or cross-compilation)
make
./dendra-hummingbird-monitor

# Run with version info only
./dendra-hummingbird-monitor --version
```

## License

See repository for license information.
