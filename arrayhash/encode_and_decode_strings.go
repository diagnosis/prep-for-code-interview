package arrayhash

import (
	"strconv"
	"strings"
)

type Solution struct {
}

func (s *Solution) Encode(input []string) string {
	var b strings.Builder
	for _, str := range input {
		b.WriteString(strconv.Itoa(len(str)))
		b.WriteByte('#')
		b.WriteString(str)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	for i := 0; i < len(encoded); i++ {
		for j := i; j < len(encoded); j++ {
			if encoded[j] == '#' {
				n, err := strconv.Atoi(encoded[i:j])
				if err == nil {
					res = append(res, encoded[j+1:j+n+1])
					i = j + n
					break
				}
			}

		}
	}
	return res
}
