default:
	@echo "Make sure you are inside components/mlserve. This Makefile is ONLY for mlserve."
	@echo "Usage:"
	@echo "\tmake format"
	@echo "\tmake clean"
	@echo "\tmake package"
	@echo "\tmake test"
	@echo "\tmake build"
	@echo "\tmake publish-test"
	@echo "\tmake publish-prod"

format:
	autoflake -i **/*.py
	isort -i **/*.py
	yapf -i **/*.py

clean:
	rm -rf build
	rm -rf dist
	rm -rf mlpm.egg-info

build:
	python3 setup.py sdist bdist_wheel

test:
	PYTHONPATH=./ python3 tests/server.py

publish-test:
	twine upload --repository-url https://test.pypi.org/legacy/ dist/*

publish:
	twine upload dist/*