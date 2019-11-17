# markov-go-wasm
This project build a Markov Chain Playground on WASM

## Goal

This project was born with the purpose of loading a basic implementation of the Markov chain completely on the browser, compiles in WASM and is loaded directly on the github-pages through the directory *docs*.

The result is available at: [https://blog.m4dfry.space/markov-go-wasm/](https://blog.m4dfry.space/markov-go-wasm/)

## Makefile

Compile main.go with WASM and load on docs folder
> make wasm

Run serve of docs folder to test locally
> make serv

## Notes

Go compiler for WASM is on early stage an library like "syscall/js" are still in experimental state.
This project is compiled with Go 1.13.4.

## TODO
 - Add sentence creation from seed
 - Check if seed is OK with order
 - Export JSON on file

## License

Used software (all software is licensed under appropriate license of authors)
*  Go Programming Language [https://golang.org/](https://golang.org/) BSD 3-clause "New" or "Revised" License
*  gomarkov - Go implementation of markov chains for textual data [https://github.com/mb-14/gomarkov](https://github.com/mb-14/gomarkov) MIT License

---

Have a request, suggestion, errata, critic or question? Open an Issue!

---
