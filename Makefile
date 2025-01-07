build :
	go mod tidy
	go build -o LineWalker

run :
	go run .

push :
	git add .
	git commit -a
	git push

pull :
	git pull
	make build