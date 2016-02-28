// +build use_easyjson

package benchmark

import (
	"testing"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jwriter"
)

func BenchmarkUnmarshalRequestEJ(b *testing.B) {
	b.SetBytes(int64(len(largeStructText)))
	for i := 0; i < b.N; i++ {
		var s LargeStruct
		err := s.UnmarshalJSON(largeStructText)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMarshalRequestEJ(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := easyjson.Marshal(&largeStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestEJ(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := easyjson.Marshal(&xlStructData)
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestToWriterEJ(b *testing.B) {
	var l int64
	out := &DummyWriter{}
	for i := 0; i < b.N; i++ {
		w := jwriter.Writer{}
		xlStructData.MarshalEasyJSON(&w)
		if w.Error != nil {
			b.Error(w.Error)
		}

		l = int64(w.Size())
		w.DumpTo(out)
	}
	b.SetBytes(l)

}
func BenchmarkMarshalRequestEJParallel(b *testing.B) {
	b.SetBytes(int64(len(largeStructText)))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := largeStructData.MarshalJSON()
			if err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkMarshalToWriterEJ(b *testing.B) {
	var l int64
	out := &DummyWriter{}
	for i := 0; i < b.N; i++ {
		w := jwriter.Writer{}
		largeStructData.MarshalEasyJSON(&w)
		if w.Error != nil {
			b.Error(w.Error)
		}

		l = int64(w.Size())
		w.DumpTo(out)
	}
	b.SetBytes(l)

}
func BenchmarkMarshalRequestToWriterEJParallel(b *testing.B) {
	out := &DummyWriter{}

	b.RunParallel(func(pb *testing.PB) {
		var l int64
		for pb.Next() {
			w := jwriter.Writer{}
			largeStructData.MarshalEasyJSON(&w)
			if w.Error != nil {
				b.Error(w.Error)
			}

			l = int64(w.Size())
			w.DumpTo(out)
		}
		if l > 0 {
			b.SetBytes(l)
		}
	})

}

func BenchmarkMarshalXLRequestEJParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := xlStructData.MarshalJSON()
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}

func BenchmarkMarshalXLRequestToWriterEJParallel(b *testing.B) {
	out := &DummyWriter{}
	b.RunParallel(func(pb *testing.PB) {
		var l int64
		for pb.Next() {
			w := jwriter.Writer{}

			xlStructData.MarshalEasyJSON(&w)
			if w.Error != nil {
				b.Error(w.Error)
			}
			l = int64(w.Size())
			w.DumpTo(out)
		}
		if l > 0 {
			b.SetBytes(l)
		}
	})
}

func BenchmarkUnmarshalSmallRequestEJ(b *testing.B) {
	b.SetBytes(int64(len(smallStructText)))

	for i := 0; i < b.N; i++ {
		var s Entities
		err := s.UnmarshalJSON(smallStructText)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMarshalSmallRequestEJ(b *testing.B) {
	var l int64
	for i := 0; i < b.N; i++ {
		data, err := smallStructData.MarshalJSON()
		if err != nil {
			b.Error(err)
		}
		l = int64(len(data))
	}
	b.SetBytes(l)
}

func BenchmarkMarshalSmallRequestEJParallel(b *testing.B) {
	var l int64
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data, err := smallStructData.MarshalJSON()
			if err != nil {
				b.Error(err)
			}
			l = int64(len(data))
		}
	})
	b.SetBytes(l)
}
