# Docker commands for single node
single_node_up:
	docker-compose -f docker-compose.single.yml up -d

single_node_down:
	docker-compose -f docker-compose.single.yml down

single_node_attach:
	docker exec -it manticore mysql

single_node_clear:
	rm -rf ./data/single/*



## API commands
api_create_table:
	curl -X POST 'http://127.0.0.1:9308/sql' -d 'mode=raw&query=CREATE TABLE testrt ( title text, content text, gid integer)'

api_insert_data:
	curl -X POST 'http://127.0.0.1:9308/json/insert' -d'{"index":"testrt","id":1,"doc":{"title":"Hello","content":"world","gid":1}}'

api_search:
	curl -X POST 'http://127.0.0.1:9308/json/search' -d '{"index":"testrt","query":{"match":{"*":"hello world"}}}'