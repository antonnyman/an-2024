package layout

import (
	"encoding/hex"
	"hash/fnv"
	"os"
)

func latestRevisionUrl(path string) string {
	content, err := os.ReadFile("." + path)
	if err != nil {
		return path // Fallback to original path if can't read
	}

	h := fnv.New64a()
	_, err = h.Write(content)
	if err != nil {
		return path
	}

	hash := hex.EncodeToString(h.Sum(nil))[:8]
	return path + "?v=" + hash
}
