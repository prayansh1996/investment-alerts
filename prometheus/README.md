Commands Ran:

Docker:

sudo groupadd -f docker

sudo usermod -aG docker $USER

// After logging back in:
sudo service docker start

newgrp docker


Docker Compose:
sudo curl -L https://github.com/docker/compose/releases/download/1.22.0/docker-compose-Linux-x86_64 -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose


