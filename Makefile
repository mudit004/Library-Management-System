Scripts = config.sh virtualHosting.sh

virtualHosting:

	sudo chmod +x ./scripts/virtualHosting.sh
	./scripts/virtualHosting.sh

goSetup:
	go mod tidy
	go mod vendor

script: 
	./scripts/script.sh


.PHONY: all script