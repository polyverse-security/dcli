package toggles

import "strings"

// Toggles are parsed at app start
type Toggle struct {
	name   string
	active bool
}

var mToggle map[string]*Toggle

func NewToggle(name string) *Toggle {
	if mToggle == nil {
		mToggle = make(map[string]*Toggle)
	}

	if _, ok := mToggle[name]; ok {
		return mToggle[name]
	}

	t := &Toggle{name: name, active: false}
	mToggle[name] = t
	return t
}

func ParseToggles(args []string) []string {
	var notToggles []string
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			cleaned := strings.TrimPrefix(args[i], "--")
			// Check if toggle is registered
			if _, ok := mToggle[cleaned]; ok {
				mToggle[cleaned].active = true // activate the toggle
			} else {
				notToggles = append(notToggles, args[i]) // not a toggle, add it to notToggles
			}
			continue
		} else {
			notToggles = append(notToggles, args[i]) // not a toggle, add it to notToggles
		}
	}
	return notToggles
}

func GetToggle(name string) Toggle {
	if _, ok := mToggle[name]; ok {
		return *mToggle[name]
	}
	return Toggle{name: name}
}

func (t Toggle) Active() bool {
	return t.active
}
