// https://leetcode.com/problems/defanging-an-ip-address
package easy

import "strings"

func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}
