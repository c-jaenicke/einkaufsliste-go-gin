package category

type Category struct {
	ID         int
	Name       string
	Counter    int
	Color      string
	Color_name string
}

var ColorMap = map[string]string{
	"#FFFFFF": "Weiss",
	"#D3F8E2": "Grün",
	"#E4C1F9": "Violet",
	"#F694C1": "Rot",
	"#EDE7B1": "Gelb",
	"#A9DEF9": "Blau",
	"#BAF2D8": "Türkis",
	"#FFC4D6": "Pink",
	"#FEC5BB": "Orange",
}
