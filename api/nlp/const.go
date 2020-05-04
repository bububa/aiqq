package nlp

type ComType = int

const (
	UNKNOWN_COM ComType = 0
	WEATHER_COM ComType = 1
	MUSIC_COM   ComType = 2
	STOCK_COMP  ComType = 3
	NEWS_COMP   ComType = 4
)

type IntentType = int

const (
	UNKNOWN_INTENT          = 0
	LYRICS_INTENT           = 1
	DOWNLOAD_INTENT         = 2
	INSTRUMENT_INTENT       = 3
	MUSIC_INTENT            = 4
	NAME_INTENT             = 5
	TIME_INTENT             = 6
	LOCATION_INTENT         = 7
	STYLE_INTENT            = 8
	NUMBER_INTENT           = 9
	VIDEO_INTENT            = 10
	NATION_INTENT           = 11
	ALBUM_INTENT            = 12
	ORDER_INTENT            = 13
	VARIETY_INTENT          = 14
	BAND_INTENT             = 15
	SCENIC_INTENT           = 16
	MOVIE_INTENT            = 17
	TV_INTENT               = 18
	WIKI_INTENT             = 19
	STOCK_NAME_INTENT       = 34
	STOCK_CODE_INTENT       = 35
	INDEX_INDENT            = 36
	PRICE_INTENT            = 37
	INSIGHT_INTENT          = 38
	MOUNTAIN_INTENT         = 40
	LAKE_INTENT             = 41
	WETHER_INTENT           = 42
	RESTAURANT_INTENT       = 43
	DISH_NAME_INTENT        = 44
	RHYMES_INTENT           = 45
	STORY_INTENT            = 46
	CROSS_TALK_INTENT       = 47
	STORYTELLING_INTENT     = 48
	AUDIO_BOOK_INTENT       = 49
	CATEGORY_WORD_INTENT    = 128
	RELATION_WORD_INTENT    = 129
	CONTRACTION_WORD_INTENT = 130
)

type PolarType = int

const (
	NEGTIVE_POLAR  PolarType = -1
	NEUTRAL_POLAR  PolarType = 0
	POSITIVE_POLAR PolarType = 1
)
