package feeder

type Media struct {
	Thumbnails []*Thumbnail
}

type Thumbnail struct {
	Url    string
	Width  int
	Height int
	Time   string
}
