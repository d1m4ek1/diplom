node_prod:
	npm run prod

node_watch:
	npm run watch

go_codegangsta:
	gin -i --appPort 5050 --port 3000 -d ./backend/cmd/main