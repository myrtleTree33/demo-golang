default:
	echo "hi there!"

start:	
	make down
	make up

# setup cluster of dependencies in docker
up:
	docker compose -f "compose.yml" up -d --build --wait

# teardown dev setup
down:
	docker compose -f "compose.yml" down
	(echo "y" | docker volume prune)
	clear
	