package kopak

import (
	"github.com/Ershad5/kopak2"
)

func Bark() string {
	return "neyniakh ada!"
}

func Barks() string {

	return "ada dodan neyniakh!"
}

func Bigbark() string {
	return kopak2.WhenGrownUp(Bark())
}
func Bigbarks() string {
	return kopak2.WhenGrownUp(Barks())
}

// i added the tags to test it on git hub V1.1.0
