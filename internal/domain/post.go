package domain

const TYPE_TEXT = "text"
const TYPE_PIC = "pic"
const TYPE_VIDEO = "video"

type Post struct {
	Title   string
	Source  string
	Content string
	Type    string
}
