# xk6-neo4j

A K6 Extension for Neo4j

---

## Development

### From Source

- go 1.17+
- k6 v0.38.3
- xk6 v0.7.0
- neo4j community 4.4.x
- openjdk 11

1\. Install Go 1.17+ < 1.18:

> Go 1.18.x breaks a number of libraries; atm `k6` builds with 1.17.9

2\. Install `xk6` for building custom K6 binaries:

`go install go.k6.io/xk6/cmd/xk6@latest`

3\. Install [https://golangci-lint.run/usage/install/](https://golangci-lint.run/usage/install/):

`go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2`

4\. Install [Neo4j community edition](https://neo4j.com/download-center/):

`wget https://dist.neo4j.org/neo4j-community-4.4.8-unix.tar.gz`

5\. Configure Prometheus Metrics for Neo4j

[How to monitor Neo4j with Prometheus](https://neo4j.com/developer/kb/how-to-monitor-neo4j-with-prometheus/)

`neo4j.conf`

```
# Bolt connector
dbms.connector.bolt.enabled=true
dbms.connector.bolt.tls_level=DISABLED
# bind to 0.0.0.0 when using WSL otherwise localhost
dbms.connector.bolt.listen_address=0.0.0.0:7687
dbms.connector.bolt.advertised_address=:7687

# HTTP Connector. There can be zero or one HTTP connectors.
dbms.connector.http.enabled=true
# bind to 0.0.0.0 when using WSL otherwise localhost
dbms.connector.http.listen_address=0.0.0.0:7474
dbms.connector.http.advertised_address=:7474

# Enable the Prometheus endpoint. Default is 'false'.
metrics.prometheus.enabled=true
# The IP and port the endpoint will bind to in the format <hostname or IP address>:<port number>.
# The default is localhost:2004.
# bind to 0.0.0.0 when using WSL otherwise localhost
metrics.prometheus.endpoint=0.0.0.0:2004
```

❗ Don't bind to `0.0.0.0` on a VPS like Digital Ocean or EC2!

> Be sure to restart the neo4j proc for your changes to take effect

`prometheus.yml`

```
# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  - job_name: 'neo4j-prometheus'

    # metrics_path: /metrics
    # scheme defaults to 'http'.

    static_configs:
      - targets: ['localhost:2004']
```

6/. Install OpenJDK 11 [Neo4j dep]

```
java --version
openjdk 11.0.15 2022-04-19
OpenJDK Runtime Environment Temurin-11.0.15+10 (build 11.0.15+10)
OpenJDK 64-Bit Server VM Temurin-11.0.15+10 (build 11.0.15+10, mixed mode)
```

7/. Start Neo4j

`/usr/local/neo4j/bin/neo4j start`

8/. Raise the `ulimit` if you encounter: 

> WARNING: Max 4096 open files allowed, minimum of 40000 recommended. See the Neo4j manual.

https://neo4j.com/developer/kb/number-of-open-files-on-linux/

❗ `ulimit` on WSL is tricky: https://github.com/microsoft/WSL/discussions/6226

WSL for Windows

Add the following to the top of your `~/.zshrc`:

```
# https://github.com/microsoft/WSL/discussions/6226
ULIMIT=65536
if [[ "$(ulimit -n)" != $ULIMIT ]]; then
    sudo prlimit --nofile=$ULIMIT:$ULIMIT --pid $$
    exec /bin/zsh
fi
```

then run the following:

```
source ~/.zshrc
ulimit -n
```

should return `$ULIMIT`

Restart `neo4j` once more and the warning should resolve:

```
./bin/neo4j restart
Neo4j is not running.
Directories in use:
home:         /usr/local/neo4j
config:       /usr/local/neo4j/conf
logs:         /usr/local/neo4j/logs
plugins:      /usr/local/neo4j/plugins
import:       /usr/local/neo4j/import
data:         /usr/local/neo4j/data
certificates: /usr/local/neo4j/certificates
licenses:     /usr/local/neo4j/licenses
run:          /usr/local/neo4j/run
Starting Neo4j.
Started neo4j (pid:1324). It is available at http://localhost:7474
There may be a short delay until the server is ready.
```
