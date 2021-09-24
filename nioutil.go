package nioutil

import "io"

type DoubleReader struct {
	reader  int
	Readers [2]io.Reader
}

func (mr *DoubleReader) Read(p []byte) (n int, err error) {
	for mr.reader < 2 {
		n, err = mr.Readers[mr.reader].Read(p)
		if err == io.EOF {
			mr.reader += 1
		}
		if n > 0 || err != io.EOF {
			if err == io.EOF && mr.reader < 2 {
				err = nil
			}
			return
		}
	}
	return 0, io.EOF
}

func (d *DoubleReader) Reset(first, second io.Reader) {
	*d = DoubleReader{
		Readers: [2]io.Reader{first, second},
	}
}
