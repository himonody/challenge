package lang

import (
	"strings"
	"sync"
)

const (
	English      = "en"
	BahasaMelayu = "ms"
	Tamil        = "ta"
	Indonesia    = "id"
	Chinese      = "zh"
)

var LangMap sync.Map

func init() {
	LangMap.Store(English, en)
	LangMap.Store(Chinese, zh)
	LangMap.Store(BahasaMelayu, ms)
	LangMap.Store(Tamil, ta)
	LangMap.Store(Indonesia, id)
}

// Normalize turns variant codes (e.g. zh-CN, en-US) into canonical codes.
func Normalize(langCode string) string {
	if langCode == "" {
		return English
	}
	langCode = strings.ToLower(strings.ReplaceAll(langCode, "_", "-"))
	base := langCode
	if idx := strings.Index(langCode, "-"); idx > 0 {
		base = langCode[:idx]
	}
	switch base {
	case English:
		return English
	case Chinese:
		return Chinese
	case BahasaMelayu, "my", "msa":
		return BahasaMelayu
	case Tamil:
		return Tamil
	case Indonesia, "in":
		return Indonesia
	default:
		return English
	}
}

// Get returns message by language code and error code via sync.Map.
func Get(langCode string, code int) string {
	langCode = Normalize(langCode)
	if v, ok := LangMap.Load(langCode); ok {
		if m, ok := v.(map[int]string); ok && m != nil {
			msg, ok := m[code]
			if ok {
				return msg
			}
			return ""
		}
	}
	return ""
}

// Set replaces language map safely via sync.Map.
func Set(langCode string, m map[int]string) {
	if langCode == "" || m == nil {
		return
	}
	LangMap.Store(langCode, m)
}
