package api

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
	RelatedProducts []string  `json:"related_products"` //Optional, by default null
	Brand           string    `json:"brand"`            //Optional, by default "" Facetable
	Gender          string    `json:"gender"`           //Optional, by default "" Facetable
	Rating          []int     `json:"rating"`           //Optional, by default null
	Comments        []Comment `json:"comments"`         //Optional, by default null
	Levels          Levels    `json:"levels"`           //Optional, by default "" Facetable
	Metadata        Metadata  `json:"metadata"`
	Url             string    `json:"url"`
}

type UpdateProduct struct {
	Doc Product `json:"doc"`
}

type Products []Product

type RelatedProducts struct {
	Hits     int      `json:"hits"`
	Products Products `json:"products"`
}

type OneProductResponse struct {
	Product         Product         `json:"product"`
	RelatedProducts RelatedProducts `json:"relatedProducts"`
}

type SearchResponse struct {
	Products     Products               `json:"products"`
	Aggregations map[string]interface{} `json:"aggregations"`
	Hits         int                    `json:"hits"`
}

type AggsResponse struct {
	Aggregations map[string]interface{} `json:"aggregations"`
}

type AutocompleteResponse struct {
	Autocomplete Autocompletes `json:"autocomplete"`
}

type Social struct {
	Instagram string `json:"instagram"` //Optional
	Facebook  string `json:"facebook"`  //Optional
	Linkedin  string `json:"linkedin"`  //Optional
}

type Author struct {
	Avatar string `json:"avatar"` //Optional
	Name   string `json:"name"`   //Optional
	Role   string `json:"role"`   //Optional
}

type Post struct {
	Title       string    `json:"title"`       //Required
	ID          string    `json:"id"`          //Required
	Description string    `json:"description"` //Required
	Images      []string  `json:"images"`      //Required
	Author      Author    `json:"author"`      //Optional
	Available   bool      `json:"available"`   //Required (whether it currently is shown or not)
	Date        string    `json:"date"`        //Required
	Label       []string  `json:"label"`       //Required
	Content     string    `json:"content"`     //Optional, by default null
	SocialMed   Social    `json:"social"`      //Optional
	Rating      []int     `json:"rating"`      //Optional, by default null
	Comments    []Comment `json:"comments"`    //Optional, by default null
}

type Posts []Post
