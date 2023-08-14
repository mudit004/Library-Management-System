Scripts = config.sh virtualHosting.sh

all: $(Scripts)

$(SCRIPTS):
	cd ./scripts
    ./$@

goSetup:
	go mod tidy
	go mod vendor

script: 
	cd ./scripts
	./script.sh


.PHONY: all