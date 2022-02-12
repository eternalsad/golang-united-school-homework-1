package solution

import (
	"github.com/kyokomi/emoji"
)

// GetMessage returns hello world string
func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
