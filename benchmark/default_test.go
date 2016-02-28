// +build !use_easyjson,!use_ffjson

package benchmark

import (
	"encoding/json"
	"testing"
)

func BenchmarkUnmarshalRequestStd(b *testing.B) {
	b.SetBytes(int64(len(largeStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		err := json.Unmarshal(largeStructText, &s)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMarshalRequestStd(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(&largeStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestStd(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(&xlStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalRequestStdParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := json.Marshal(&largeStructData)
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestStdParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := json.Marshal(&xlStructData)
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}

func BenchmarkUnmarshalSmallRequestStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s Entities
		err := json.Unmarshal(smallStructText, &s)
		if err != nil {
			b.Error(err)
		}
	}
	b.SetBytes(int64(len(smallStructText)))
}

func BenchmarkMarshalSmallRequestStd(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := json.Marshal(&smallStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalSmallRequestStdParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := json.Marshal(&smallStructData)
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}

func BenchmarkMarshalToWriterStd(b *testing.B) {
	enc := json.NewEncoder(&DummyWriter{})
	for i := 0; i < b.N; i++ {
		err := enc.Encode(&largeStructData)
		if err != nil {
			b.Error(err)
		}
	}
}
