package builder

import (
	"encoding/base64"
	"strings"
)

var Base64KeywordEncoding = base64KeywordEncoding{
	PlusSign:  "Σ",
	SlashSign: "Δ",
	EqualSign: "Ξ",
}

type base64KeywordEncoding struct {
	PlusSign  string
	SlashSign string
	EqualSign string
}

func (e base64KeywordEncoding) EncodeToString(bs []byte) string {
	str := base64.StdEncoding.EncodeToString(bs)
	str = strings.ReplaceAll(str, "+", e.PlusSign)
	str = strings.ReplaceAll(str, "/", e.SlashSign)
	return strings.ReplaceAll(str, "=", e.EqualSign)
}
