package category

type Category struct {
	ID        int
	Name      string
	Counter   int
	Color     string
	ColorName string
}

var ColorMap = map[string]string{
	"#FFFFFF": "White",
	"#B5F1CC": "Green",
	"#AAE3E2": "Cyan",
}
