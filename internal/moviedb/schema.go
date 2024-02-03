package moviedb

type multiResponse struct {
	Page       int            `json:"page"`
	Results    []searchResult `json:"results"`
	TotalPages int            `json:"total_pages"`
}

type searchResult struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	ID               int      `json:"id"`
	Title            string   `json:"title,omitempty"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	MediaType        string   `json:"media_type"`
	GenreIds         []int    `json:"genre_ids"`
	Popularity       float64  `json:"popularity"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	Video            bool     `json:"video,omitempty"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Name             string   `json:"name,omitempty"`
	OriginalName     string   `json:"original_name,omitempty"`
	FirstAirDate     string   `json:"first_air_date,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
}
