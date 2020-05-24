# isup
Check if a list of systems are up

[![Go Report Card](https://goreportcard.com/badge/github.com/psenna/isup)](https://goreportcard.com/report/github.com/psenna/isup)

# testing 

go test ./... -coverprofile=cover.out

go tool cover -func=cover.out

go tool cover -html=cover.out -o cover.html