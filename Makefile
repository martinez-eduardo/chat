
push:
	git status
	git add .
	git commit -m "$$(date)"
	git pull origin main 
	git push origin main

compile:
	go build main.go

run:
	go run main.go

