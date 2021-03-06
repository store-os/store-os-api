// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/{client}/autocomplete": {
            "get": {
                "description": "get autocomplete",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "autocomplete"
                ],
                "summary": "List autocomplete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client",
                        "name": "client",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "q",
                        "description": "name search by q",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "category=",
                        "description": "category filter",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "subcategory=",
                        "description": "subcategory filter",
                        "name": "subcategory",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "subsubcategory=",
                        "description": "subsubcategory filter",
                        "name": "subsubcategory",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "from",
                        "description": "from price",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "to",
                        "description": "to price",
                        "name": "to",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Autocomplete"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/{client}/blog": {
            "get": {
                "description": "List posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Blog endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client",
                        "name": "client",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "page=1",
                        "description": "paging number",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/{client}/blog/{id}": {
            "get": {
                "description": "get post by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "Blog endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client",
                        "name": "client",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/{client}/products": {
            "get": {
                "description": "List products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Products endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client",
                        "name": "client",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "q",
                        "description": "name search by q",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "page=1",
                        "description": "paging number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "category=",
                        "description": "category filter",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "subcategory=",
                        "description": "subcategory filter",
                        "name": "subcategory",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "subsubcategory=",
                        "description": "subsubcategory filter",
                        "name": "subsubcategory",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "from",
                        "description": "from price",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "to",
                        "description": "to price",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "page",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "fieldsort",
                        "description": "fieldsort final_price or title.keyword",
                        "name": "fieldsort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "order",
                        "description": "order (asc or desc)",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "size",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/{client}/products/{id}": {
            "get": {
                "description": "get product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "OneProduct endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client",
                        "name": "client",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.OneProductResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "post product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Products endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "client",
                        "name": "client",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Product"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Author": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "Optional",
                    "type": "string"
                },
                "name": {
                    "description": "Optional",
                    "type": "string"
                },
                "role": {
                    "description": "Optional",
                    "type": "string"
                }
            }
        },
        "api.Autocomplete": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.Comment": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "int": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "response": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.Feature": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "api.Levels": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "subcategory": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "subsubcategory": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.Metadata": {
            "type": "object",
            "properties": {
                "equipment": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "features": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Feature"
                    }
                },
                "specs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Spec"
                    }
                },
                "stocks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Stock"
                    }
                }
            }
        },
        "api.OneProductResponse": {
            "type": "object",
            "properties": {
                "product": {
                    "$ref": "#/definitions/api.Product"
                },
                "relatedProducts": {
                    "$ref": "#/definitions/api.RelatedProducts"
                }
            }
        },
        "api.Post": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "Optional",
                    "$ref": "#/definitions/api.Author"
                },
                "available": {
                    "description": "Required (whether it currently is shown or not)",
                    "type": "boolean"
                },
                "comments": {
                    "description": "Optional, by default null",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Comment"
                    }
                },
                "content": {
                    "description": "Optional, by default null",
                    "type": "string"
                },
                "date": {
                    "description": "Required",
                    "type": "string"
                },
                "description": {
                    "description": "Required",
                    "type": "string"
                },
                "id": {
                    "description": "Required",
                    "type": "string"
                },
                "images": {
                    "description": "Required",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "label": {
                    "description": "Required",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "rating": {
                    "description": "Optional, by default null",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "social": {
                    "description": "Optional",
                    "$ref": "#/definitions/api.Social"
                },
                "title": {
                    "description": "Required",
                    "type": "string"
                }
            }
        },
        "api.Product": {
            "type": "object",
            "properties": {
                "available": {
                    "description": "Required Facetable",
                    "type": "boolean"
                },
                "brand": {
                    "description": "Optional, by default \"\" Facetable",
                    "type": "string"
                },
                "comments": {
                    "description": "Optional, by default null",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Comment"
                    }
                },
                "date": {
                    "description": "Optional",
                    "type": "string"
                },
                "description": {
                    "description": "Required",
                    "type": "string"
                },
                "discount": {
                    "description": "Optional, by default 0",
                    "type": "integer"
                },
                "final_price": {
                    "description": "Optional",
                    "type": "integer"
                },
                "gender": {
                    "description": "Optional, by default \"\" Facetable",
                    "type": "string"
                },
                "id": {
                    "description": "Required",
                    "type": "string"
                },
                "images": {
                    "description": "Required",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "levels": {
                    "description": "Optional, by default \"\" Facetable",
                    "$ref": "#/definitions/api.Levels"
                },
                "metadata": {
                    "$ref": "#/definitions/api.Metadata"
                },
                "mini_description": {
                    "description": "Optional",
                    "type": "string"
                },
                "price": {
                    "description": "Optional Facetable range | Sortable",
                    "type": "integer"
                },
                "rating": {
                    "description": "Optional, by default null",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "related_products": {
                    "description": "Optional, by default null",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "ship_price": {
                    "description": "Optional, by default 0",
                    "type": "integer"
                },
                "title": {
                    "description": "Required Sortable | Relevance",
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "api.RelatedProducts": {
            "type": "object",
            "properties": {
                "hits": {
                    "type": "integer"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Product"
                    }
                }
            }
        },
        "api.Social": {
            "type": "object",
            "properties": {
                "facebook": {
                    "description": "Optional",
                    "type": "string"
                },
                "instagram": {
                    "description": "Optional",
                    "type": "string"
                },
                "linkedin": {
                    "description": "Optional",
                    "type": "string"
                }
            }
        },
        "api.Spec": {
            "type": "object",
            "properties": {
                "measure": {
                    "type": "string"
                },
                "spec": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "api.Stock": {
            "type": "object",
            "properties": {
                "color": {
                    "description": "Facetable",
                    "type": "string"
                },
                "sizes": {
                    "description": "Facetable",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "store-api.alchersan.com",
	BasePath:    "/api/v1",
	Schemes:     []string{"https"},
	Title:       "Swagger Store OS API",
	Description: "This is a sample server celler server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
