package api

var size int = 20

type Suggestion struct {
	Title string `json:"title"`
}

type Suggestions []Suggestion

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
	DiscountPrice   int       `json:"discount_price"`   //Optional, by default 0
	Brand           string    `json:"brand"`            //Optional, by default "" Facetable
	Gender          string    `json:"gender"`           //Optional, by default "" Facetable
	Rating          []int     `json:"rating"`           //Optional, by default null
	Comments        []Comment `json:"comments"`         //Optional, by default null
	Category        []string  `json:"category"`         //Optional, by default null Facetable
	Subcategory     []string  `json:"subcategory"`      //Optional, by default null Facetable
	Metadata        Metadata  `json:"metadata"`
}

type Products []Product

type SearchResponse struct {
	Products     Products               `json:"products"`
	Aggregations map[string]interface{} `json:"aggregations"`
	Hits         int                    `json:"hits"`
}
