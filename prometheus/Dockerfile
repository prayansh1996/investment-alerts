# Use the official Prometheus image as a base
FROM prom/prometheus

USER root

# Copy your Prometheus configuration file into the image
COPY ./prometheus.yaml /etc/prometheus/prometheus.yaml

# Set the working directory inside the container
WORKDIR /prometheus
