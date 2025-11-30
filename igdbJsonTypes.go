package main

type IgdbToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type IgdbGames []struct {
	ID                    int     `json:"id"`
	AgeRatings            []int   `json:"age_ratings"`
	AggregatedRating      float64 `json:"aggregated_rating,omitempty"`
	AggregatedRatingCount int     `json:"aggregated_rating_count,omitempty"`
	AlternativeNames      []int   `json:"alternative_names,omitempty"`
	Bundles               []int   `json:"bundles,omitempty"`
	Cover                 int     `json:"cover"`
	CreatedAt             int     `json:"created_at"`
	Expansions            []int   `json:"expansions,omitempty"`
	ExternalGames         []int   `json:"external_games,omitempty"`
	FirstReleaseDate      int     `json:"first_release_date"`
	Franchise             int     `json:"franchise,omitempty"`
	Franchises            []int   `json:"franchises"`
	GameEngines           []int   `json:"game_engines,omitempty"`
	GameModes             []int   `json:"game_modes"`
	Genres                []int   `json:"genres"`
	InvolvedCompanies     []int   `json:"involved_companies"`
	Keywords              []int   `json:"keywords,omitempty"`
	MultiplayerModes      []int   `json:"multiplayer_modes,omitempty"`
	Name                  string  `json:"name"`
	Platforms             []int   `json:"platforms"`
	PlayerPerspectives    []int   `json:"player_perspectives"`
	Rating                float64 `json:"rating,omitempty"`
	RatingCount           int     `json:"rating_count,omitempty"`
	ReleaseDates          []int   `json:"release_dates"`
	Screenshots           []int   `json:"screenshots"`
	SimilarGames          []int   `json:"similar_games"`
	Slug                  string  `json:"slug"`
	Storyline             string  `json:"storyline,omitempty"`
	Summary               string  `json:"summary"`
	Tags                  []int   `json:"tags"`
	Themes                []int   `json:"themes"`
	TotalRating           float64 `json:"total_rating,omitempty"`
	TotalRatingCount      int     `json:"total_rating_count,omitempty"`
	UpdatedAt             int     `json:"updated_at"`
	URL                   string  `json:"url"`
	Videos                []int   `json:"videos,omitempty"`
	Websites              []int   `json:"websites,omitempty"`
	Checksum              string  `json:"checksum"`
	Remakes               []int   `json:"remakes,omitempty"`
	Remasters             []int   `json:"remasters,omitempty"`
	GameLocalizations     []int   `json:"game_localizations,omitempty"`
	Collections           []int   `json:"collections,omitempty"`
	GameType              int     `json:"game_type"`
	Artworks              []int   `json:"artworks,omitempty"`
	ParentGame            int     `json:"parent_game,omitempty"`
	LanguageSupports      []int   `json:"language_supports,omitempty"`
	VersionParent         int     `json:"version_parent,omitempty"`
	VersionTitle          string  `json:"version_title,omitempty"`
	Hypes                 int     `json:"hypes,omitempty"`
	Dlcs                  []int   `json:"dlcs,omitempty"`
}
