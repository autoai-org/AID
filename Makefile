default:
	@echo "Usage:"
	@echo "\tmake test"
	@echo "\tmake format"
	@echo "\tmake docs"
	@echo "\tmake package"
	@echo "\tmake clean"
	@echo "\tmake publish"
test:

format:
	autoflake -i cvpm/*.py
	# autoflake -i cvpm/**/*.py

	isort -i cvpm/*.py
	# isort -i cvpm/**/*.py 

	yapf -i cvpm/*.py
	# yapf -i cvpm/**/*.py

docs:
	npm run docs:build

package:
	python setup.py sdist bdist_wheel

clean:
	rm -rf build
	rm -rf dist
	rm -rf cvpm.egg-info

publish-test:
	twine upload --repository-url https://test.pypi.org/legacy/ dist/*

.PHONY: docs