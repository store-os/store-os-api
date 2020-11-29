package api

var size int = 20

type Autocomplete struct {
	Title string `json:"title"`
	Image string `json:"image"`
	ID    string `json:"id"`
}

type Autocompletes []Autocomplete

type Feature struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}
type Spec struct {
	Measure string `json:"measure"`
	Spec    string `json:"spec"`
	Value   string `json:"value"`
}
type Stock struct {
	Color string   `json:"color"` //Facetable
	Sizes []string `json:"sizes"` //Facetable
}
type Metadata struct {
	Equipment []string  `json:"equipment"`
	Stocks    []Stock   `json:"stocks"`
	Features  []Feature `json:"features"`
	Specs     []Spec    `json:"specs"`
}
type Comment struct {
	Name        string `json:"name"`
	Rating      int    `json:"int"`
	Title       string `json:"title"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Response    string `json:"response"`
}
type Levels struct {
	Category       []string `json:"category"`
	SubCategory    []string `json:"subcategory"`
	SubSubCategory []string `json:"subsubcategory"`
}
type Product struct {
	Title           string    `json:"title"`            //Required Sortable | Relevance
	Date            string    `json:"date"`             //Optional
	ID              string    `json:"id"`               //Required
	Description     string    `json:"description"`      //Required
	MiniDescription string    `json:"mini_description"` //Optional
	Images          []string  `json:"images"`           //Required
	Available       bool      `json:"available"`        //Required Facetable
	Price           int       `json:"price"`            //Optional Facetable range | Sortable
	ShipPrice       int       `json:"ship_price"`       //Optional, by default 0
	Discount        int       `json:"discount"`         //Optional, by default 0
	FinalPrice      int       `json:"final_price"`      //Optional
	Brand           string    `json:"brand"`            //Optional, by default "" Facetable
	Gender          string    `json:"gender"`           //Optional, by default "" Facetable
	Rating          []int     `json:"rating"`           //Optional, by default null
	Comments        []Comment `json:"comments"`         //Optional, by default null
	Levels          Levels    `json:"levels"`           //Optional, by default "" Facetable
	Metadata        Metadata  `json:"metadata"`
	Url             string    `json:"url"`
}

type Products []Product

type SearchResponse struct {
	Products     Products               `json:"products"`
	Aggregations map[string]interface{} `json:"aggregations"`
	Hits         int                    `json:"hits"`
}

type AutocompleteResponse struct {
	Autocomplete Autocompletes `json:"autocomplete"`
}

type Paragraph struct {
	Header   string   `json:"header"`   //Optional
	Content  string   `json:"content"`  //Required
	Images   []string `json:"images"`   //Optional
	Position string   `json:"position"` //Image position regarding the text. (Top, Left, Right, Bottom)
}
type Social struct {
	Instagram string `json:"instagram"` //Optional
	Facebook  string `json:"facebook"`  //Optional
	Linkedin  string `json:"linkedin"`  //Optional
}

type Post struct {
	Title       string      `json:"title"`       //Required
	ID          string      `json:"id"`          //Required
	Description string      `json:"description"` //Required
	Images      []string    `json:"images"`      //Required
	Author      string      `json:"author"`      //Optional
	Available   bool        `json:"available"`   //Required (whether it currently is shown or not)
	Date        string      `json:"date"`        //Required
	Label       []string    `json:"label"`       //Required
	Content     []Paragraph `json:"paragraph"`   //Optional, by default null
	SocialMed   Social      `json:"social"`      //Optional
	Rating      []int       `json:"rating"`      //Optional, by default null
	Comments    []Comment   `json:"comments"`    //Optional, by default null
}

type Posts []Post
