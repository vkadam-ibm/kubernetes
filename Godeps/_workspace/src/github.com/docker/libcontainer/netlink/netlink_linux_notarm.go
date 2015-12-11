// +build !arm,!s390x,!ppc64,!ppc64le

package netlink

func ifrDataByte(b byte) int8 {
	return int8(b)
}
