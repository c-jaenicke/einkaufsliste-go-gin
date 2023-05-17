// Package category
// Contains struct wich embodies a category and a map for looking up color names
package category

type Category struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Counter    int    `json:"counter"`
	Color      string `json:"color"`
	Color_name string `json:"color_Name"`
}

// ColorMap is a map for looking up which name references which color
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
