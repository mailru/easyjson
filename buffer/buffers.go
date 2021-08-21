package buffer

import "io"

// This code is copied from golang/net, which is not supported by TinyGo

// Buffers contains zero or more runs of bytes to write.
//
// On certain machines, for certain types of connections, this is
// optimized into an OS-specific batch write operation (such as
// "writev").
type Buffers [][]byte

var (
	_ io.WriterTo = (*Buffers)(nil)
	_ io.Reader   = (*Buffers)(nil)
)

func (v *Buffers) WriteTo(w io.Writer) (n int64, err error) {
	// Remove this as TinyGo cannot handle such reflection
	// if wv, ok := w.(buffersWriter); ok {
	// 	return wv.writeBuffers(v)
	// }
	for _, b := range *v {
		nb, err := w.Write(b)
		n += int64(nb)
		if err != nil {
			v.consume(n)
			return n, err
		}
	}
	v.consume(n)
	return n, nil
}

func (v *Buffers) Read(p []byte) (n int, err error) {
	for len(p) > 0 && len(*v) > 0 {
		n0 := copy(p, (*v)[0])
		v.consume(int64(n0))
		p = p[n0:]
		n += n0
	}
	if len(*v) == 0 {
		err = io.EOF
	}
	return
}

func (v *Buffers) consume(n int64) {
	for len(*v) > 0 {
		ln0 := int64(len((*v)[0]))
		if ln0 > n {
			(*v)[0] = (*v)[0][n:]
			return
		}
		n -= ln0
		(*v)[0] = nil
		*v = (*v)[1:]
	}
}
