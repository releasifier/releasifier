clean:
	@rm -rf ./bin
	@mkdir -p ./bin/temp
	@mkdir -p ./bin/bundle

dist-tools:
	go get -u github.com/pkieltyka/fresh
	go get -v github.com/robfig/glock

deps:
	@glock -v sync github.com/alinz/releasifier

build: clean
	@mkdir -p ./bin
	cd cmd/releasifier && GOGC=off go build -i -o ../../bin/releasifier

build-all: clean
	cd cmd/releasifier; \
	for GOOS in darwin linux windows; do \
		for GOARCH in 386 amd64; do \
			echo "building $$GOOS $$GOARCH ..."; \
			export GOOS=$$GOOS; \
			export GOARCH=$$GOARCH; \
			go build -o ../../bin/releasifier-$$GOOS-$$GOARCH; \
		done \
	done

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


docker-rm-all-images: docker-rm-existing-ps
	docker images | grep "^<none>" | awk '{print $$3}' | xargs docker rmi

docker-rm-existing-ps:
	docker ps -a | awk 'NR>1' | awk '{print $$1}' | xargs docker rm
