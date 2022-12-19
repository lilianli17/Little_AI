# Random sentence generator

[![Go Report Card](https://goreportcard.com/badge/github.com/lilianli17/random_sentence_generator)](https://goreportcard.com/report/github.com/lilianli17/random_sentence_generator)


> Let AI generate your own words!

## Introduction

Final project for CIS 1930. An n-gram language model that can be used to generate random text resembling a source document 

## Usage Instructions

``` go
go get github.com/lilianli17/random_sentence_generator
```

## Simple Usages
```
var model NGramModel
model = NewNGram(5, "frankenstein.txt")
result := model.GetRandomText(6)
// "It was morning, I remember"
```

## Documentation

Documentation for the package can be found here. (https://pkg.go.dev/github.com/lilianli17/random_sentence_generator)

## License

MIT © 2022 Lilian Li
