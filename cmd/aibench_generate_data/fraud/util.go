package fraud

import (
	"encoding/binary"
	"math"
	"math/rand"
	"strconv"
)

func randomByteStringSliceChoice(s [][]byte) []byte {
	return s[rand.Intn(len(s))]
}

func randomInt64SliceChoice(s []int64) int64 {
	return s[rand.Intn(len(s))]
}

func ConvertSliceStringToFloat(transactionDataString []string) []float32 {
	res := make([]float32, len(transactionDataString))
	for i := range transactionDataString {
		value, _ := strconv.ParseFloat(transactionDataString[i], 64)
		res[i] = float32(value)
	}
	return res
}

func randReferenceData(n int) []float32 {
	res := make([]float32, n)
	for i := range res {
		res[i] = rand.Float32()
	}
	return res
}

func Uint64frombytes(bytes []byte) uint64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return bits
}

func Float32bytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}