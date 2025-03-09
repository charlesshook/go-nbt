# NBT

A NBT library written in Go (Golang). Currently this library only works for Minecraft Bedrock level.dat files.

## Install
```
go get -u github.com/charlesshook/go-nbt
```

## Usage

Reading Minecraft Bedrock level.dat and printing the output.
```
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charlesshook/go-nbt"
)

func main() {
	file, err := os.Open("level.dat")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	nbtData, err := nbt.Read(file)
	if err != nil {
		fmt.Println("Error reading NBT:", err)
		return
	}
	prettyData, _ := json.MarshalIndent(nbtData, "", "  ")
	fmt.Println("NBT Data:", string(prettyData))
}
```