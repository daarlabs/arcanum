package hub

import "bytes"

func trimBytesNewlines(v []byte) []byte {
	return bytes.TrimSpace(bytes.Replace(v, []byte{'\n'}, []byte{' '}, -1))
}
