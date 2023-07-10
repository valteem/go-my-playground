package reuse_test

import(
	"fmt"
	"testing"

	"github.com/valteem/reuse"
)

func TestTellStory(t *testing.T) {

	fmt.Println(reuse.TellStory("good", "long", "stupid"))
}