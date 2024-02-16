# Create persistent volume data
docker volume create prometheus-data

# Start Prometheus container
docker run \
    -p 9091:9091 \
    -v ./prometheus.yml:/etc/prometheus/prometheus.yml \
    -v ./prometheus-data:/prometheus \
    prom/prometheus