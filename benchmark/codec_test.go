// +build use_codec

package benchmark

import (
	"testing"

	"github.com/ugorji/go/codec"
)

func BenchmarkCodec_Unmarshal_M(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)
	dec := codec.NewDecoderBytes(nil, h)

	b.SetBytes(int64(len(largeStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		dec.ResetBytes(largeStructText)
		if err := dec.Decode(&s); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCodec_Unmarshal_S(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)
	dec := codec.NewDecoderBytes(nil, h)

	b.SetBytes(int64(len(smallStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		dec.ResetBytes(smallStructText)
		if err := dec.Decode(&s); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkCodec_Marshal_S(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&smallStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = nil
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_M(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&largeStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = nil
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_L(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&xlStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = nil
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_S_Reuse(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&smallStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = out[:0]
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_M_Reuse(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&largeStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = out[:0]
	}

	b.SetBytes(l)
}

func BenchmarkCodec_Marshal_L_Reuse(b *testing.B) {
	var h codec.Handle = new(codec.JsonHandle)

	var out []byte
	enc := codec.NewEncoderBytes(&out, h)

	var l int64
	for i := 0; i < b.N; i++ {
		enc.ResetBytes(&out)
		if err := enc.Encode(&xlStructData); err != nil {
			b.Error(err)
		}
		l = int64(len(out))
		out = out[:0]
	}

	b.SetBytes(l)
}

/*

func BenchmarkCodec_Marshal_RequestFF(b *testing.B) {
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

func BenchmarkCodec_Marshal_SmallRequestFF(b *testing.B) {
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

func BenchmarkCodec_Marshal_RequestFFPool(b *testing.B) {
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

func BenchmarkCodec_Marshal_XLRequestFF(b *testing.B) {
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

func BenchmarkCodec_Marshal_XLRequestFFPool(b *testing.B) {
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

func BenchmarkCodec_Marshal_XLRequestFFPoolParallel(b *testing.B) {
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
func BenchmarkCodec_Marshal_RequestFFPoolParallel(b *testing.B) {
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

func BenchmarkCodec_Marshal_SmallRequestPoolFF(b *testing.B) {
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

func BenchmarkCodec_Marshal_SmallRequestPoolFFParallel(b *testing.B) {
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

func BenchmarkCodec_Marshal_SmallRequestFFParallel(b *testing.B) {
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

func BenchmarkCodec_Marshal_RequestFFParallel(b *testing.B) {
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

func BenchmarkCodec_Marshal_XLRequestFFParallel(b *testing.B) {
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
*/
