package structs

import "strconv"
import "strings"

func toIPv4(value string) string {

	if strings.Contains(value, ".") {

		var tmp1 []string = strings.Split(value, ".")
		var tmp2 []uint8
		var tmp3 []string

		for t := 0; t < len(tmp1); t++ {

			num, err := strconv.ParseUint(tmp1[t], 10, 8)

			if err == nil {

				if num >= 0 && num <= 255 {
					tmp2 = append(tmp2, uint8(num))
				}

			}

		}

		for t := 0; t < len(tmp2); t++ {

			var str = strconv.FormatUint(uint64(tmp2[t]), 10)

			if str != "" {
				tmp3 = append(tmp3, str)
			}

		}

		if len(tmp3) == 4 {
			return strings.Join(tmp3, ".")
		}

	}

	return ""

}

func toIPv6(value string) string {

	if strings.Contains(value, "::") {

		var tmp = strings.Split(value, "::")
		var prefix = strings.Split(tmp[0], ":")
		var suffix = strings.Split(tmp[1], ":")
		var complete = []string{"0000", "0000", "0000", "0000", "0000", "0000", "0000", "0000"}

		for p := 0; p < len(prefix); p++ {

			var chunk = prefix[p]
			var c = p

			if len(chunk) == 1 {
				complete[c] = "000" + chunk
			} else if len(chunk) == 2 {
				complete[c] = "00" + chunk
			} else if len(chunk) == 3 {
				complete[c] = "0" + chunk
			} else if len(chunk) == 4 {
				complete[c] = chunk
			}

		}

		for s := 0; s < len(suffix); s++ {

			var c = 8 - len(suffix) + s
			var chunk = suffix[s]

			if len(chunk) == 1 {
				complete[c] = "000" + chunk
			} else if len(chunk) == 2 {
				complete[c] = "00" + chunk
			} else if len(chunk) == 3 {
				complete[c] = "0" + chunk
			} else if len(chunk) == 4 {
				complete[c] = chunk
			}

		}

		return strings.Join(complete, ":")

	} else if strings.Contains(value, ":") {

		var chunks = strings.Split(value, ":")
		var complete = []string{"0000", "0000", "0000", "0000", "0000", "0000", "0000", "0000"}

		for c := 0; c < len(chunks); c++ {

			var chunk = chunks[c]

			if len(chunk) == 1 {
				complete[c] = "000" + chunk
			} else if len(chunk) == 2 {
				complete[c] = "00" + chunk
			} else if len(chunk) == 3 {
				complete[c] = "0" + chunk
			} else if len(chunk) == 4 {
				complete[c] = chunk
			}

		}

		return strings.Join(complete, ":")

	}

	return ""

}

func toMaskv4(value string) string {

	var mask = []string{"0", "0", "0", "0", "0", "0", "0", "0"}
	var valid bool = true

	if len(value) == 8 {

		for v := 0; v < len(value); v++ {

			_, err := strconv.ParseUint(string(value[v]), 16, 64)

			if err != nil {
				valid = false
				break
			} else {
				mask[v] = string(value[v])
			}

		}

	}

	if valid == true {
		return strings.Join(mask, "")
	}

	return strings.Join(mask, "")

}

func toMaskv6(value string) string {

	var mask = []string{
		"0", "0", "0", "0", "0", "0", "0", "0",
		"0", "0", "0", "0", "0", "0", "0", "0",
		"0", "0", "0", "0", "0", "0", "0", "0",
		"0", "0", "0", "0", "0", "0", "0", "0",
	}
	var valid bool = true

	if len(value) == 32 {

		for v := 0; v < len(value); v++ {

			_, err := strconv.ParseUint(string(value[v]), 16, 64)

			if err != nil {
				valid = false
				break
			} else {
				mask[v] = string(value[v])
			}

		}

	}

	if valid == true {
		return strings.Join(mask, "")
	}

	return strings.Join(mask, "")

}

func toScopev4(value string) string {

	var scope string = "public"
	var PRIVATE_V4 = []string{

		// RFC1122
		"0.0.0.0",
		"127.",

		// RFC1918
		"10.",
		"172.16.", "172.17.", "172.18.", "172.19.",
		"172.20.", "172.21.", "172.22.", "172.23.",
		"172.24.", "172.25.", "172.26.", "172.27.",
		"172.28.", "172.29.", "172.30.", "172.31.",
		"192.168.",

		// RFC2544
		"198.0.", "198.2.", "198.4.", "198.6.",
		"198.8.", "198.10.", "198.12.", "198.14.",
		"198.16.", "198.18.", "198.20.", "198.22.",
		"198.24.", "198.26.", "198.28.", "198.30.",
		"198.32.", "198.34.", "198.36.", "198.38.",
		"198.40.", "198.42.", "198.44.", "198.46.",
		"198.48.", "198.50.", "198.52.", "198.54.",
		"198.56.", "198.58.", "198.60.", "198.62.",
		"198.64.", "198.66.", "198.68.", "198.70.",
		"198.72.", "198.74.", "198.76.", "198.78.",
		"198.80.", "198.82.", "198.84.", "198.86.",
		"198.88.", "198.90.", "198.92.", "198.94.",
		"198.96.", "198.98.", "198.100.", "198.102.",
		"198.104.", "198.106.", "198.108.", "198.110.",
		"198.112.", "198.114.", "198.116.", "198.118.",
		"198.120.", "198.122.", "198.124.", "198.126.",
		"198.128.", "198.130.", "198.132.", "198.134.",
		"198.136.", "198.138.", "198.140.", "198.142.",
		"198.144.", "198.146.", "198.148.", "198.150.",
		"198.152.", "198.154.", "198.156.", "198.158.",
		"198.160.", "198.162.", "198.164.", "198.166.",
		"198.168.", "198.170.", "198.172.", "198.174.",
		"198.176.", "198.178.", "198.180.", "198.182.",
		"198.184.", "198.186.", "198.188.", "198.190.",
		"198.192.", "198.194.", "198.196.", "198.198.",
		"198.200.", "198.202.", "198.204.", "198.206.",
		"198.208.", "198.210.", "198.212.", "198.214.",
		"198.216.", "198.218.", "198.220.", "198.222.",
		"198.224.", "198.226.", "198.228.", "198.230.",
		"198.232.", "198.234.", "198.236.", "198.238.",
		"198.240.", "198.242.", "198.244.", "198.246.",
		"198.248.", "198.250.", "198.252.", "198.254.",

		// RFC3068
		"192.88.99.",

		// RFC3927
		"169.254.",

		// RFC6598
		"100.64.",
		"100.128.",
		"100.192.",

		// RFC6890
		"192.0.2.",
		"198.51.100.",
		"203.0.113.",
	}

	for p := 0; p < len(PRIVATE_V4); p++ {

		if strings.HasPrefix(value, PRIVATE_V4[p]) {
			scope = "private"
			break
		}

	}

	return scope

}

func toScopev6(value string) string {

	var scope string = "public"
	var PRIVATE_V6 = []string{

		// RFC3513
		"0000:0000:0000:0000:0000:0000:0000:0000",
		"0000:0000:0000:0000:0000:0000:0000:0001",
		"fe80:0000:0000:0000",
	}

	for p := 0; p < len(PRIVATE_V6); p++ {

		if strings.HasPrefix(value, PRIVATE_V6[p]) {
			scope = "private"
			break
		}

	}

	return scope

}

type Address struct {
	IP    string `json:"ip"`
	Mask  string `json:"mask"`
	Scope string `json:"scope"`
	Type  string `json:"type"`
}

func NewAddress(ip_value string, mask_value string) Address {

	var address Address

	if strings.Contains(ip_value, ":") {

		var ip = toIPv6(ip_value)
		var mask = toMaskv6(mask_value)
		var scope = toScopev6(ip)

		if ip != "" && scope != "" {
			address.IP = ip
			address.Mask = mask
			address.Scope = scope
			address.Type = "v6"
		}

	} else if strings.Contains(ip_value, ".") {

		var ip = toIPv4(ip_value)
		var mask = toMaskv4(mask_value)
		var scope = toScopev4(ip)

		if ip != "" && mask != "" && scope != "" {
			address.IP = ip
			address.Mask = mask
			address.Scope = scope
			address.Type = "v4"
		}

	}

	return address

}

func (address *Address) IsValid() bool {

	if address.IP != "" {

		if address.Type == "v4" {

			var result bool = true

			if address.IP != toIPv4(address.IP) {
				result = false
			}

			if address.Mask != toMaskv4(address.Mask) {
				result = false
			}

			if address.Scope != toScopev4(address.IP) {
				result = false
			}

			return result

		} else if address.Type == "v6" {

			var result bool = true

			if address.IP != toIPv6(address.IP) {
				result = false
			}

			if address.Mask != toMaskv6(address.Mask) {
				result = false
			}

			if address.Scope != toScopev6(address.IP) {
				result = false
			}

			return result

		}

	}

	return false

}

func (address *Address) SetIP(value string) {

	if strings.Contains(value, ":") {

		var ip = toIPv6(value)
		var mask = toMaskv6("")
		var scope = toScopev6(value)

		if ip != "" {
			address.IP = ip
			address.Mask = mask
			address.Scope = scope
			address.Type = "v6"
		}

	} else if strings.Contains(value, ".") {

		var ip = toIPv4(value)
		var mask = toMaskv4("")
		var scope = toScopev4(value)

		if ip != "" {
			address.IP = ip
			address.Mask = mask
			address.Scope = scope
			address.Type = "v4"
		}

	}

}

func (address *Address) SetMask(value string) {

	if address.Type == "v6" {

		var mask = toMaskv6(value)

		if mask != "" {
			address.Mask = mask
		}

	} else if address.Type == "v4" {

		var mask = toMaskv4(value)

		if mask != "" {
			address.Mask = mask
		}

	}

}

func (address *Address) SetScope(value string) {

	if value == "private" || value == "public" {
		address.Scope = value
	}

}
