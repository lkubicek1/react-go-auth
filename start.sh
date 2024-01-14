export COMPOSE_FILES="$(find . | grep "docker-compose.yml")"

for COMPOSE_FILE in $COMPOSE_FILES
do 
	docker-compose -f $COMPOSE_FILE up -d --build
done
