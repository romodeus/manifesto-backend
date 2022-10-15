docker compose --file ./docker-compose-dev.yaml down
docker rmi shortly
docker compose --file ./docker-compose-dev.yaml --env-file .env up -d