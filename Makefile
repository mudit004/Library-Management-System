Scripts = config.sh virtualHosting.sh

all: $(Scripts)

$(SCRIPTS):
    .scripts/$@

goSetup:
	go mod tidy
	go mod vendor

script: 
	.scripts/script.sh


.PHONY: all script