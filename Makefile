default:
	@echo "Usage:"
	@echo "\tmake test"
	@echo "\tmake format"
	@echo "\tmake docs"
	@echo "\tmake package"
	@echo "\tmake clean"
	@echo "\tmake publish"

test-cli:
	cd core && go test -race -coverprofile c.out -v ./...
	cd core/cmd && go build

format-py:
	autoflake -i cvpm/*.py
	# autoflake -i cvpm/**/*.py

	isort -i cvpm/*.py
	# isort -i cvpm/**/*.py 

	yapf -i cvpm/*.py
	# yapf -i cvpm/**/*.py

format-go:
	gofmt -l -s -w **/*.go

package:
	python3 setup.py sdist bdist_wheel

clean:
	rm -rf build
	rm -rf dist
	rm -rf cvpm.egg-info

publish-test:
	twine upload --repository-url https://test.pypi.org/legacy/ dist/*

publish-prod:
	twine upload dist/*

build-discovery:
	cd discovery && yarn run build

publish-discovery:
	git subtree push --prefix discovery heroku master

.PHONY: docs