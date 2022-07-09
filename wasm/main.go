package main

import (
	"fmt"
	js "syscall/js"
)

func main() {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		fmt.Println("Unable to get document object")
	}
	inputTextArea := jsDoc.Call("getElementById", "input")
	if !inputTextArea.Truthy() {
		fmt.Println("Unable to get input text area")
	}
	outputTextArea := jsDoc.Call("getElementById", "board")
	if !outputTextArea.Truthy() {
		fmt.Println("Unable to get output text area")
	}
	outputTextArea.Set("value", "Go on...\n")
	jsClickEnterHandler := js.FuncOf(func(this js.Value, args []js.Value) any {
		event := args[0]
		if event.Get("key").String() == "Enter" {
			event.Call("preventDefault")
			boardContent := outputTextArea.Get("value").String()
			boardContent += "me: " + inputTextArea.Get("value").String() + "\n"
			inputTextArea.Set("value", js.Null())
			outputTextArea.Set("value", js.ValueOf(boardContent))
		}
		return nil
	})
	if !jsClickEnterHandler.Truthy() {
		fmt.Println("Unable to render the clickEnterHandler function")
	}
	defer jsClickEnterHandler.Release()
	inputTextArea.Call("addEventListener", "keypress", jsClickEnterHandler)
	// fmt.Println("Hi!")
	<-make(chan bool)
}
