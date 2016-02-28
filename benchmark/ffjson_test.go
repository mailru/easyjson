// +build use_ffjson

package benchmark

import (
	"testing"

	"github.com/pquerna/ffjson/ffjson"
)

func BenchmarkUnmarshalRequestFF(b *testing.B) {
	b.SetBytes(int64(len(largeStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		err := ffjson.UnmarshalFast(largeStructText, &s)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMarshalRequestFF(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := ffjson.MarshalFast(&largeStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalSmallRequestFF(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := ffjson.MarshalFast(&smallStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalRequestFFPool(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := ffjson.MarshalFast(&largeStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
		ffjson.Pool(data)
	}
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestFF(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := ffjson.MarshalFast(&xlStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestFFPool(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := ffjson.MarshalFast(&xlStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
		ffjson.Pool(data)
	}
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestFFPoolParallel(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := ffjson.MarshalFast(&xlStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
		ffjson.Pool(data)
	}
	b.SetBytes(l)
}
func BenchmarkMarshalRequestFFPoolParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := ffjson.MarshalFast(&largeStructData)
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
			ffjson.Pool(data)
		}
	})
	b.SetBytes(l)
}

func BenchmarkUnmarshalSmallRequestFF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s Entities
		err := ffjson.UnmarshalFast(smallStructText, &s)
		if err != nil {
			b.Error(err)
		}
	}
	b.SetBytes(int64(len(smallStructText)))
}

func BenchmarkMarshalSmallRequestPoolFF(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := ffjson.MarshalFast(&smallStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
		ffjson.Pool(data)
	}
	b.SetBytes(l)
}

func BenchmarkMarshalSmallRequestPoolFFParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := ffjson.MarshalFast(&smallStructData)
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
			ffjson.Pool(data)
		}
	})
	b.SetBytes(l)
}

func BenchmarkMarshalSmallRequestFFParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := ffjson.MarshalFast(&smallStructData)
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}

func BenchmarkMarshalRequestFFParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := ffjson.MarshalFast(&largeStructData)
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestFFParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := ffjson.MarshalFast(&xlStructData)
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}
