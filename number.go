package gofakeit

import (
	"errors"
	"math"
	"math/rand"
)

// Number will generate a random number between given min And max
func Number(min int, max int) int {
	return randIntRange(min, max)
}

// Uint8 will generate a random uint8 value
func Uint8() uint8 {
	return uint8(randIntRange(0, math.MaxUint8))
}

// Uint16 will generate a random uint16 value
func Uint16() uint16 {
	return uint16(randIntRange(0, math.MaxUint16))
}

// Uint32 will generate a random uint32 value
func Uint32() uint32 {
	return uint32(randIntRange(0, math.MaxInt32))
}

// Uint64 will generate a random uint64 value
func Uint64() uint64 {
	return uint64(rand.Int63n(math.MaxInt64))
}

// Int8 will generate a random Int8 value
func Int8() int8 {
	return int8(randIntRange(math.MinInt8, math.MaxInt8))
}

// Int16 will generate a random int16 value
func Int16() int16 {
	return int16(randIntRange(math.MinInt16, math.MaxInt16))
}

// Int32 will generate a random int32 value
func Int32() int32 {
	return int32(randIntRange(math.MinInt32, math.MaxInt32))
}

// Int64 will generate a random int64 value
func Int64() int64 {
	return rand.Int63n(math.MaxInt64) + math.MinInt64
}

// Float32 will generate a random float32 value
func Float32() float32 {
	return randFloat32Range(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Float32Range will generate a random float32 value between min and max
func Float32Range(min, max float32) float32 {
	return randFloat32Range(min, max)
}

// Float64 will generate a random float64 value
func Float64() float64 {
	return randFloat64Range(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// Float64Range will generate a random float64 value between min and max
func Float64Range(min, max float64) float64 {
	return randFloat64Range(min, max)
}

// ShuffleInts will randomize a slice of ints
func ShuffleInts(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func addNumberLookup() {
	AddLookupData("number", Info{
		Category:    "number",
		Description: "Random number between given range",
		Example:     "14866",
		Params: []Param{
			{Field: "min", Type: "int", Default: "-2147483648", Description: "Minimum integer value"},
			{Field: "max", Type: "int", Default: "2147483647", Description: "Maximum integer value"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			min, err := info.GetInt(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetInt(m, "max")
			if err != nil {
				return nil, err
			}

			if min > max {
				return nil, errors.New("Max integer must be larger than Min")
			}

			return Number(min, max), nil
		},
	})

	AddLookupData("uint8", Info{
		Category:    "number",
		Description: "Random uint8 value",
		Example:     "152",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Uint8(), nil
		},
	})

	AddLookupData("uint16", Info{
		Category:    "number",
		Description: "Random uint16 value",
		Example:     "34968",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Uint16(), nil
		},
	})

	AddLookupData("uint32", Info{
		Category:    "number",
		Description: "Random uint32 value",
		Example:     "1075055705",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Uint32(), nil
		},
	})

	AddLookupData("uint64", Info{
		Category:    "number",
		Description: "Random uint64 value",
		Example:     "843730692693298265",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Uint64(), nil
		},
	})

	AddLookupData("int8", Info{
		Category:    "number",
		Description: "Random int8 value",
		Example:     "24",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Int8(), nil
		},
	})

	AddLookupData("int16", Info{
		Category:    "number",
		Description: "Random int16 value",
		Example:     "2200",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Int16(), nil
		},
	})

	AddLookupData("int32", Info{
		Category:    "number",
		Description: "Random int32 value",
		Example:     "-1072427943",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Int32(), nil
		},
	})

	AddLookupData("int64", Info{
		Category:    "number",
		Description: "Random int64 value",
		Example:     "-8379641344161477543",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Int64(), nil
		},
	})

	AddLookupData("float32", Info{
		Category:    "number",
		Description: "Random float32 value",
		Example:     "3.1128167e+37",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Float32(), nil
		},
	})

	AddLookupData("float32range", Info{
		Category:    "number",
		Description: "Random float32 between given range",
		Example:     "914774.6",
		Params: []Param{
			{Field: "min", Type: "int", Description: "Minimum float32 value"},
			{Field: "max", Type: "int", Description: "Maximum float32 value"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			min, err := info.GetFloat32(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat32(m, "max")
			if err != nil {
				return nil, err
			}

			return Float32Range(min, max), nil
		},
	})

	AddLookupData("float64", Info{
		Category:    "number",
		Description: "Random float64 value",
		Example:     "1.644484108270445e+307",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Float64(), nil
		},
	})

	AddLookupData("float64range", Info{
		Category:    "number",
		Description: "Random float64 between given range",
		Example:     "914774.5585333086",
		Params: []Param{
			{Field: "min", Type: "int", Description: "Minimum float64 value"},
			{Field: "max", Type: "int", Description: "Maximum float64 value"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			min, err := info.GetFloat64(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat64(m, "max")
			if err != nil {
				return nil, err
			}

			return Float64Range(min, max), nil
		},
	})

	AddLookupData("shuffleints", Info{
		Category:    "number",
		Description: "Shuffle an array of ints",
		Example:     "1,2,3,4 => 3,1,4,2",
		Params: []Param{
			{Field: "ints", Type: "intarray", Description: "Delimited seperated ints"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			ints, err := info.GetIntArray(m, "ints")
			if err != nil {
				return nil, err
			}

			ShuffleInts(ints)

			return ints, nil
		},
	})
}
