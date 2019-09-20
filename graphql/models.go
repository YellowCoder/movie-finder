package graphql

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type NewMovie struct {
	Title string `json:"title"`
}
