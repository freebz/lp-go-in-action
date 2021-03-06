package main

import "fmt"

func Trim(s string, cutset string) string {
    if s == "" || cutset == "" {
        return s
    }
    return TrimFunc(s, makeCutsetFunc(cutset))
}

func isShellSpecialVar(c uint8) bool {
    switch c {
    case '*', '#', '$', '@', '!', '?', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        return true
    }
    return false
}

type IP []byte

func (ip IP) MarshalText() ([]byte, error) {
    if len(ip) == 0 {
        return []byte(""), nil
    }
    if len(ip) != IPv4len && len(ip) != IPv6len {
        return nil, errors.New("invalid IP address")
    }
    return []byte(ip.String()), nil
}

// ipEmptyString is like ip.String except that is returns
// an empty string when ip is unset.
func ipEmptyString(ip IP) string {
    if len(ip) == 0 {
        return ""
    }
    return ip.String()
}
