package utils

import (
	"bytes"
	"encoding/binary"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandNum(i int) int {
	if i == 0 {
		return 0
	}
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(i)
}

//float32和float64的slice互转
func Float64SliceToFloat32Slice(f []float64) []float32 {
	var o = make([]float32, len(f))
	for k := range f {
		o[k] = float32(f[k])
	}

	return o
}

func Float32SliceToFloat64Slice(f []float32) []float64 {
	var o = make([]float64, len(f))
	for k := range f {
		o[k] = float64(f[k])
	}

	return o
}

//以下是处理float64
func Float64SliceToBytes(f []float64) []byte {
	var buffer bytes.Buffer
	for _, v := range f {
		b := Float64ToByte(v)
		buffer.Write(b)
	}

	return buffer.Bytes()
}

func BytesToFloat64Slice(b []byte) (rs []float64) {
	for i := 0; i < len(b); i = i + 8 {
		rs = append(rs, ByteToFloat64(b[i:i+8]))
	}

	return
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(b []byte) float64 {
	bit := binary.LittleEndian.Uint64(b)

	return math.Float64frombits(bit)
}

func JoinFloat64BySep(a []float64, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return strconv.FormatFloat(a[0], 'f', -1, 64)
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(strconv.FormatFloat(a[i], 'f', -1, 64))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.FormatFloat(a[0], 'f', -1, 64))
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(strconv.FormatFloat(s, 'f', -1, 64))
	}
	return b.String()
}

//以下是处理float32
func Float32SliceToBytes(f []float32) []byte {
	var buffer bytes.Buffer
	for _, v := range f {
		b := Float32ToByte(v)
		buffer.Write(b)
	}

	return buffer.Bytes()
}

func BytesToFloat32Slice(b []byte) (rs []float32) {
	for i := 0; i < len(b); i = i + 4 {
		rs = append(rs, ByteToFloat32(b[i:i+4]))
	}

	return
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(b []byte) float32 {
	bit := binary.LittleEndian.Uint32(b)

	return math.Float32frombits(bit)
}

func StringToFloat32(str string, sep string) ([]float32, error) {
	string_array := strings.Split(str, sep)
	var float_array []float32
	var err error
	for _, lmk := range string_array {
		v, err := strconv.ParseFloat(lmk, 32)
		if err != nil {
			return nil, err
		}
		float_array = append(float_array, float32(v))
	}
	return float_array, err
}

func JoinFloat32BySep(f []float32, sep string) string {
	return JoinFloat64BySep(Float32SliceToFloat64Slice(f), sep)
}