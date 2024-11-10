package layout

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"os"
)

// func cssUrl() string {
// 	path := "/assets/stylesheets/styles.css"

// 	fileInfo, err := os.Stat("." + path)

// 	if err != nil {
// 		return ""
// 	}

// 	modTime := fileInfo.ModTime().Unix()

// 	h := fnv.New64a()
// 	_, err = h.Write([]byte(fmt.Sprintf("%d", modTime)))
// 	if err != nil {
// 		return ""
// 	}

// 	hash := h.Sum(nil)
// 	shortHash := hex.EncodeToString(hash)[:8]

// 	return path + "?v=" + shortHash

// }

func latestRevisionUrl(path string) string {
	
	fileInfo, err := os.Stat("." + path)
	
	if err != nil {
		return ""
	}

	modTime := fileInfo.ModTime().Unix()

	h := fnv.New64a()
	_, err = h.Write([]byte(fmt.Sprintf("%d", modTime)))
	if err != nil {
		return ""
	}

	hash := h.Sum(nil)
	shortHash := hex.EncodeToString(hash)[:8]
	
	return path + "?v=" + shortHash

}