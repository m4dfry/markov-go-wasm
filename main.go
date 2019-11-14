package main

import (
	"encoding/json"
	"strings"
	"syscall/js"

	"github.com/mb-14/gomarkov"
)

var chain *gomarkov.Chain

func newChain(this js.Value, i []js.Value) interface{} {
	order := i[0].Int()
	chain = gomarkov.NewChain(order)

	return 1
}

func read(this js.Value, i []js.Value) interface{} {
	return i[0].String()
}

func getJSONChain(this js.Value, i []js.Value) interface{} {
	json, err := json.Marshal(chain)
	if err != nil {
		return 0
	}
	//println(json)
	return string(json)
}

func load(this js.Value, i []js.Value) interface{} {
	line := i[0].String()
	chain.Add(strings.Split(line, " "))

	json, err := json.Marshal(chain)
	if err != nil {
		return 0
	}
	//println(json)
	return string(json)
}

func registerCallbacks() {
	js.Global().Set("wRead", js.FuncOf(read))
	js.Global().Set("wLoad", js.FuncOf(load))
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
