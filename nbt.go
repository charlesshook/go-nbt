// Package nbt provides functions for reading Minecraft NBT data.
package nbt

import (
	"bytes"
	"fmt"
	"io"
)

// NBT Tag constants represent the different types of tages used in the
// Minecraft Bedrock NBT (Named Binary Tag) format.
const (
	TagEnd       byte = 0x00 // TagEnd represents the end of a compound tag.
	TagByte      byte = 0x01 // TagByte represents an 8-bit signed integer.
	TagShort     byte = 0x02 // TagShort represents a 16-bit signed integer.
	TagInt       byte = 0x03 // TagInt represents a 32-bit signed integer.
	TagLong      byte = 0x04 // TagLong represents a 64-bit signed integer.
	TagFloat     byte = 0x05 // TagFloat represents a 32-bit floating-point number.
	TagDouble    byte = 0x06 // TagDouble represents a 64-bit floating-point number.
	TagByteArray byte = 0x07 // TagByteArray represents an array of bytes.
	TagString    byte = 0x08 // TagString represents a UTF-8 string.
	TagList      byte = 0x09 // TagList represents list of NBT values of the same type.
	TagCompound  byte = 0x0A // TagCompound represents a collection of key-value pairs.
	TagIntArray  byte = 0x0B // TagIntArray represents an array of 32-bit numbers.
	TagLongArray byte = 0x0C // TagLongArray represents an array of 64-bit numbers.
)

// Read takes in an io.Reader and returns a map[string]interface{} that has the NBT parsed
// values.
func Read(r io.Reader) (map[string]interface{}, error) {
	data, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	// Manually remove the 8-byte header specific to Bedrock level.dat
	// Show allow the user to set an option for if Bedrock or Java.
	data = data[8:]

	nbtReader := bytes.NewReader(data)

	parsedNbt, err := readNbtCompound(nbtReader)

	if err != nil {
		return nil, fmt.Errorf("error reading root compound: %v", err)
	}

	return parsedNbt, nil
}
