package model

const TYPE_TEXT = "text"
const TYPE_PIC = "pic"
const TYPE_VIDEO = "video"
const TYPE_GALLERY = "gallery"

type Post struct {
	Title   string
	Source  string
	Content string
	Type    string
	Media   []string
}
