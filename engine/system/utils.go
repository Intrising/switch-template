package system

import (
	"net"
	"strconv"
)

func ConvertIPMaskToCIDR(addr, netmask string) (string, error) {
	mask := net.IPMask(net.ParseIP(netmask).To4())
	prefixSize, _ := mask.Size()
	_, subnet, err := net.ParseCIDR(addr + "/" + strconv.Itoa(prefixSize))
	if err != nil {
		return "", err
	}
	return subnet.String(), nil
}
