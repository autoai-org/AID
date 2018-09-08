default:
	@echo "Usage:"
	@echo "\tmake test"
	@echo "\tmake format"
	@echo "\tmake docs"

test:


format:
	autoflake -i cvpm/*.py
	# autoflake -i cvpm/**/*.py

	isort -i cvpm/*.py
	# isort -i cvpm/**/*.py 

	yapf -i cvpm/*.py
	# yapf -i cvpm/**/*.py

docs:
	cd docs && npm run docs:build

.PHONY: docs