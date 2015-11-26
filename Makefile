clean:
	@rm -rf ./bin

dist-tools:
	go get -u github.com/pkieltyka/fresh
	go get -v github.com/robfig/glock

deps:
	@glock -v sync github.com/alinz/releasifier

build:
	@mkdir -p ./bin
	cd cmd/releasifier && GOGC=off go build -i -o ../../bin/releasifier

dev: kill
	@(export CONFIG=$$PWD/etc/releasifier.conf && \
		cd ./cmd/releasifier && \
		fresh -c ../../etc/fresh-runner.conf -w=../..)

##
## Database mgmt
##
reset-db:
	bash scripts/init_db.sh

kill-fresh:
	ps -ef | grep 'f[r]esh' | awk '{print $$2}' | xargs kill

kill-by-port:
	lsof -t -i:7331 | xargs kill

kill: kill-fresh kill-by-port
	
