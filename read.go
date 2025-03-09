package nbt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// readNbtByte takes in a bytes reader and return a int8 and an error.
// It reads a byte in little endian and converts to int8.
func readNbtByte(r *bytes.Reader) (int8, error) {
	var value int8
	return value, binary.Read(r, binary.LittleEndian, &value)
}

// readNbtShort takes in a bytes reader and return a int16 and an error.
// It reads bytes in little endian and converts to int8.
func readNbtShort(r *bytes.Reader) (int16, error) {
	var value int16
	return value, binary.Read(r, binary.LittleEndian, &value)
}

func readNbtInt(r *bytes.Reader) (int32, error) {
	var value int32
	return value, binary.Read(r, binary.LittleEndian, &value)
}

func readNbtLong(r *bytes.Reader) (int64, error) {
	var value int64
	return value, binary.Read(r, binary.LittleEndian, &value)
}

func readNbtFloat(r *bytes.Reader) (float32, error) {
	var value float32
	return value, binary.Read(r, binary.LittleEndian, &value)
}

func readNbtDouble(r *bytes.Reader) (float64, error) {
	var value float64
	return value, binary.Read(r, binary.LittleEndian, &value)
}

func readNbtByteArray(r *bytes.Reader) ([]byte, error) {
	var length int32
	err := binary.Read(r, binary.LittleEndian, &length)

	if err != nil {
		return nil, err
	}

	data := make([]byte, length)
	_, err = r.Read(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func readNbtString(r *bytes.Reader) (string, error) {
	var length uint16
	err := binary.Read(r, binary.LittleEndian, &length)

	if err != nil {
		return "", err
	}

	strBytes := make([]byte, length)
	_, err = r.Read(strBytes)

	if err != nil {
		return "", err
	}

	return string(strBytes), nil
}

func readNbtList(r *bytes.Reader) ([]interface{}, error) {
	var tagType byte
	var length int32
	err := binary.Read(r, binary.LittleEndian, &tagType)

	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.LittleEndian, &length)

	if err != nil {
		return nil, err
	}

	list := make([]interface{}, length)
	for i := 0; i < int(length); i++ {
		value, err := readNbtTag(tagType, r)
		if err != nil {
			return nil, err
		}
		list[i] = value
	}

	return list, nil
}

func readNbtCompound(r *bytes.Reader) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	for {
		tagType, err := r.ReadByte()

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		if tagType == TagEnd {
			break
		}

		name, err := readNbtString(r)

		if err != nil {
			return nil, fmt.Errorf("error reading tag name: %v", err)
		}

		value, err := readNbtTag(tagType, r)
		if err != nil {
			return nil, fmt.Errorf("error reading tag value for %s: %v", name, err)
		}

		data[name] = value
	}

	return data, nil
}

func readNbtIntArray(r *bytes.Reader) ([]int32, error) {
	var length int32
	err := binary.Read(r, binary.LittleEndian, &length)

	if err != nil {
		return nil, err
	}

	data := make([]int32, length)
	err = binary.Read(r, binary.LittleEndian, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func readNbtLongArray(r *bytes.Reader) ([]int64, error) {
	var length int32
	err := binary.Read(r, binary.LittleEndian, &length)

	if err != nil {
		return nil, err
	}

	data := make([]int64, length)
	err = binary.Read(r, binary.LittleEndian, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// readNbtTag takes in a tagType and a bytes.Reader and returns interface{} which will be
// a parsed NBT tag value.
func readNbtTag(tagType byte, r *bytes.Reader) (interface{}, error) {
	switch tagType {
	case TagByte:
		return readNbtByte(r)
	case TagShort:
		return readNbtShort(r)
	case TagInt:
		return readNbtInt(r)
	case TagLong:
		return readNbtLong(r)
	case TagFloat:
		return readNbtFloat(r)
	case TagDouble:
		return readNbtDouble(r)
	case TagByteArray:
		return readNbtByteArray(r)
	case TagString:
		return readNbtString(r)
	case TagList:
		return readNbtList(r)
	case TagCompound:
		return readNbtCompound(r)
	case TagIntArray:
		return readNbtIntArray(r)
	case TagLongArray:
		return readNbtLongArray(r)
	default:
		return nil, fmt.Errorf("unsupported NBT tag type: %d", tagType)
	}
}
