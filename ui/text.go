/*
 * Copyright (c) 2018, João Lucas Nunes e Silva
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in the
 *       documentation and/or other materials provided with the distribution.
 *     * Neither the name of the <organization> nor the
 *       names of its contributors may be used to endorse or promote products
 *       derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL JOÃO LUCAS NUNES E SILVA BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package ui

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func highlight(s string) string {
	runes := []rune(s)
	runesCount := len(runes)
	out := bytes.Buffer{}
	opposite := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for i := 0; i < runesCount; i++ {
		c := runes[i]

		switch {
		case strings.ContainsRune("{[(", c):
			out.WriteString("[#c83737]")
			out.WriteRune(c)
			out.WriteString("[white:#162d50]")

			i++
			balance := 1
			o := opposite[c]
			for ; i < runesCount && balance > 0; i++ {
				x := runes[i]
				if x == o {
					balance--
					if balance > 0 {
						out.WriteRune(x)
					} else {
						out.WriteString("[#c83737:black]")
						out.WriteRune(x)
						out.WriteString("[white]")
					}
				} else {
					out.WriteRune(x)
					if x == c {
						balance++
					}
				}
			}

			if runes[i] == o && balance == 0 {
				i++
			}
		case strings.ContainsRune("\"'`", c):
			out.WriteString("[#217844]")
			out.WriteRune(c)

			i++
			for ; i < runesCount; i++ {
				x := runes[i]
				out.WriteRune(x)

				if c == x {
					out.WriteString("[white]")
					break
				}
			}
		case strings.ContainsRune("/\\:-=.", c):
			out.WriteString("[#c83737]")
			out.WriteRune(c)
			out.WriteString("[white]")
		default:
			out.WriteRune(c)
		}
	}

	out.WriteString("[white:black]")
	r := regexp.MustCompile(`(line\s\d+)|(L\d+)|(col\s\d+)|(C\d+)|(column\s\d+)`)
	return r.ReplaceAllStringFunc(out.String(), func(x string) string {
		return fmt.Sprintf("[#c83737]%v[white]", x)
	})
}
