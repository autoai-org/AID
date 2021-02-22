default:
	@echo "These are global shortcuts to build AID"
	@echo "Usage:"
	@echo "\tmake format\t Format all code"

format:
	cd cmd && make format
	cd dashboard && yarn run lint
	cd discovery && yarn run lint
	cd mlserve && make format