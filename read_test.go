package nbt

import (
	"bytes"
	"testing"
)

func Test_readNbtByte(t *testing.T) {
	data := [1]byte{0x80}
	reader := bytes.NewReader(data[:])
	want := int8(-128)
	got, _ := readNbtByte(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [1]byte{0x7F}
	reader = bytes.NewReader(data[:])
	want = 127
	got, _ = readNbtByte(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [1]byte{0x32}
	reader = bytes.NewReader(data[:])
	want = 50
	got, _ = readNbtByte(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [1]byte{0xCE}
	reader = bytes.NewReader(data[:])
	want = -50
	got, _ = readNbtByte(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [1]byte{0x9C}
	reader = bytes.NewReader(data[:])
	want = -100
	got, _ = readNbtByte(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [1]byte{0x64}
	reader = bytes.NewReader(data[:])
	want = 100
	got, _ = readNbtByte(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}
}

func Test_readNbtShort(t *testing.T) {
	data := [2]byte{0x00, 0x80}
	reader := bytes.NewReader(data[:])
	want := int16(-32768)
	got, _ := readNbtShort(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [2]byte{0xFF, 0x7F}
	reader = bytes.NewReader(data[:])
	want = int16(32767)
	got, _ = readNbtShort(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [2]byte{0xD0, 0x07}
	reader = bytes.NewReader(data[:])
	want = int16(2000)
	got, _ = readNbtShort(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [2]byte{0x78, 0xEC}
	reader = bytes.NewReader(data[:])
	want = int16(-5000)
	got, _ = readNbtShort(reader)

	if got != want {
		t.Fatalf("Byte: 0x%x should result in %d but got %d.", data, want, got)
	}
}

func Test_readNbtInt(t *testing.T) {
	data := [4]byte{0xFF, 0xFF, 0xFF, 0x7F}
	reader := bytes.NewReader(data[:])
	want := int32(2147483647)
	got, _ := readNbtInt(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [4]byte{0x00, 0x00, 0x00, 0x80}
	reader = bytes.NewReader(data[:])
	want = int32(-2147483648)
	got, _ = readNbtInt(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [4]byte{0x80, 0x96, 0x98, 0x00}
	reader = bytes.NewReader(data[:])
	want = int32(10000000)
	got, _ = readNbtInt(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [4]byte{0x00, 0xD3, 0xCE, 0xFE}
	reader = bytes.NewReader(data[:])
	want = int32(-20000000)
	got, _ = readNbtInt(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}
}

func Test_readNbtLong(t *testing.T) {
	data := [10]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}
	reader := bytes.NewReader(data[:])
	want := int64(-9223372036854775808)
	got, _ := readNbtLong(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}

	data = [10]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F}
	reader = bytes.NewReader(data[:])
	want = int64(9223372036854775807)
	got, _ = readNbtLong(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %d but got %d.", data, want, got)
	}
}

func Test_readNbtFloat(t *testing.T) {
	data := [4]byte{0xFF, 0xFF, 0x7F, 0x7F}
	reader := bytes.NewReader(data[:])
	want := float32(340282346638528860000000000000000000000)
	got, _ := readNbtFloat(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %f but got %f.", data, want, got)
	}

	data = [4]byte{0x00, 0x00, 0x80, 0x00}
	reader = bytes.NewReader(data[:])
	want = float32(0.0000000000000000000000000000000000000117549435)
	got, _ = readNbtFloat(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %f but got %f.", data, want, got)
	}
}

func Test_readNbtDouble(t *testing.T) {
	data := [8]byte{0xB0, 0xD7, 0x05, 0xCD, 0xF8, 0x32, 0x41, 0x40}
	reader := bytes.NewReader(data[:])
	want := float64(34.398217799999998)
	got, _ := readNbtDouble(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %f but got %f.", data, want, got)
	}

	data = [8]byte{0x1F, 0xD4, 0xC9, 0x42, 0x09, 0x6D, 0xA4, 0xBF}
	reader = bytes.NewReader(data[:])
	want = float64(-0.03989438)
	got, _ = readNbtDouble(reader)

	if got != want {
		t.Fatalf("Byte: %x should result in %f but got %f.", data, want, got)
	}
}
