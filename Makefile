Scripts = config.sh virtualHosting.sh

all: $(Scripts)

$(Scripts):
	sudo chmod +x ./scripts/$@
	./scripts/$@

goSetup:
	go mod tidy
	go mod vendor

script: 
	./scripts/script.sh


.PHONY: all script