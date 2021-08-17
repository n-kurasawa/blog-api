TAG=latest

build:
	docker buildx build --platform linux/arm/v7 -t naohirokurasawa/blog-api:$(TAG) --push .

setup:
	docker buildx create --use --name mybuilder