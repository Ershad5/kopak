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
