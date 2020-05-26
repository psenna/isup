# isup

[![Go Report Card](https://goreportcard.com/badge/github.com/psenna/isup)](https://goreportcard.com/report/github.com/psenna/isup)
[![Build Status](https://travis-ci.org/psenna/isup.svg?branch=master)](https://travis-ci.org/psenna/isup)
[![codecov](https://codecov.io/gh/psenna/isup/branch/master/graph/badge.svg)](https://codecov.io/gh/psenna/isup)


Check if systems are running

# testing 

go test ./... -coverprofile=cover.out -covermode=atomic

go tool cover -func=cover.out

go tool cover -html=cover.out -o cover.html
