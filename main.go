package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/mb-14/gomarkov"
)

var chain *gomarkov.Chain

func echo(this js.Value, i []js.Value) interface{} {
	return i[0].String()
}

func newChain(this js.Value, i []js.Value) interface{} {
	order, err := strconv.Atoi(i[0].String())
	if err != nil {
		println("Error: invalid order.")
		return -1
	}
	chain = gomarkov.NewChain(order)

	return 1
}

func load(this js.Value, i []js.Value) interface{} {
	text := i[0].String()
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		chain.Add(strings.Split(line, " "))
	}
	return 1
}

func generate(this js.Value, i []js.Value) interface{} {
	in := i[0].String()
	lenght := i[1].Int()
	text := in
	prefix := strings.Split(in, " ")
	for i := 0; i < lenght; i++ {
		next, err := chain.Generate(prefix)
		if err != nil  {
			return text
		}
		text += " " + next
		prefix = append(prefix[1:], next)
	}
	
	return text
}

func getJSONChain(this js.Value, i []js.Value) interface{} {
	json, err := json.Marshal(chain)
	if err != nil {
		return 0
	}
	return string(json)
}

func registerCallbacks() {
	js.Global().Set("wEcho", js.FuncOf(echo))
	js.Global().Set("wLoad", js.FuncOf(load))
	js.Global().Set("wGenerate", js.FuncOf(generate))
	js.Global().Set("wNewChain", js.FuncOf(newChain))
	js.Global().Set("wGetJsonChain", js.FuncOf(getJSONChain))
	
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
