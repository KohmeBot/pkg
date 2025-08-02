package version

import (
	"fmt"
	"strconv"
	"strings"
)

type Version uint64

func NewVersion(x, y uint16, z uint32) Version {
	return X(x).Y(y).Z(z)
}

func (v Version) X(x uint16) Version {
	return (v &^ Version(0xFFFF000000000000)) | Version(x)<<48
}

func (v Version) Y(y uint16) Version {
	return (v &^ Version(0x0000FFFF00000000)) | Version(y)<<32
}

func (v Version) Z(z uint32) Version {
	return (v &^ Version(0x00000000FFFFFFFF)) | Version(z)
}

func (v Version) GetX() uint16 {
	return uint16((v & Version(0xFFFF000000000000)) >> 48)
}

func (v Version) GetY() uint16 {
	return uint16((v & Version(0x0000FFFF00000000)) >> 32)
}

func (v Version) GetZ() uint32 {
	return uint32(v & Version(0x00000000FFFFFFFF))
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.GetX(), v.GetY(), v.GetZ())
}

func X(x uint16) Version {
	var v Version
	return v.X(x)
}

func Y(y uint16) Version {
	var v Version
	return v.Y(y)
}

func Z(z uint32) Version {
	var v Version
	return v.Z(z)
}

func Parse(s string) (Version, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.TrimPrefix(s, "v")

	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid version string: %s", s)
	}

	x, err := strconv.ParseUint(parts[0], 10, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid X part: %v", err)
	}

	y, err := strconv.ParseUint(parts[1], 10, 16)
	if err != nil {
		return 0, fmt.Errorf("invalid Y part: %v", err)
	}

	z, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid Z part: %v", err)
	}

	return NewVersion(uint16(x), uint16(y), uint32(z)), nil
}
