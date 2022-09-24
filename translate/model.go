package translate

import "fmt"

const baseTwitterSearchURI = "https://twitter.com/search?q=%s&src=typed_query&f=live"

type Translation struct {
	Text         string
	Translation  string
	Language     string
	Alternatives []string
}

func (t *Translation) getTwitterLink() string {
	return fmt.Sprintf(baseTwitterSearchURI, t.Translation)
}
