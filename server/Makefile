clean:
	@rm -rf ./bin

dist-tools:
	go get -u github.com/pilu/fresh
	go get -u github.com/robfig/glock
	go get -u github.com/pressly/gv

deps:
	@glock sync github.com/alinz/releasifer/server

build:
	@mkdir -p ./bin
	cd main && GOGC=off go build -i -o ../bin/updater-server

dev:
	@(export CONFIG=$$PWD/etc/server.conf && cd ./main && fresh -c ../etc/fresh-runner.conf -w=../)

reset-db:
	bash scripts/init_db.sh
