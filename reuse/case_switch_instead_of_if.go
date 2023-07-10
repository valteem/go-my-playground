package reuse

import (
	"strings"
)

func TellStory(param1, param2, param3 string,) string {

	story := ""
	switch {
	case param1 == "good":
		story += "good "
// if first clause if fulfilled all next are just ignored		
	case param2 == "long":
		story += "long "
	case param3 == "boring":
		story += "boring"
	}

	switch {
	case story == "":
		story = "empty story"
	case story[len(story) - 1] == 32:
		story = strings.TrimSuffix(story, " ") 
	}

	return story
}