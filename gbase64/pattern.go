package gbase64

import "regexp"

var Pattern = regexp.MustCompile(`[-A-Za-z0-9+/=.]{16,}`)
