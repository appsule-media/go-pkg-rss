package feeder

type Media struct {
	Thumbnails []*Thumbnail
	Contents   []*MediaContent
}

type Thumbnail struct {
	Url    string
	Width  int
	Height int
	Time   string
	Text   string
}

type MediaContent struct {
	Url    string
	Medium string
}
