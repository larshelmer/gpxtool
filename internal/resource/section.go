package resource

type Section struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Trail      string  `json:"trail"`
	Length     float64 `json:"length"`
	GPSLength  float64 `json:"gps_length"`
	MiddleLat  float64 `json:"middle_lat"`
	MiddleLon  float64 `json:"middle_lon"`
	MaxLat     float64 `json:"max_lat"`
	MaxLon     float64 `json:"max_lon"`
	MinLat     float64 `json:"min_lat"`
	MinLon     float64 `json:"min_lon"`
	Difficulty int64   `json:"difficulty"`
	File       string  `json:"-"`
}

var Sections = []Section{
	{
		Code:       "SL1-1",
		Name:       "Sölvesborg - Grundsjön",
		Trail:      "SL1",
		Length:     14.0,
		Difficulty: 1,
		File:       "SL1_1.gpx",
	},
	{
		Code:   "SL1-2-ansl",
		Name:   "Näsum-Näsums bokskogar",
		Trail:  "SL1",
		Length: 3.0,
		File:   "SL1-2-ansl-N%25C3%25A4sum.gpx",
	},
	{
		Code:       "SL1-2",
		Name:       "Grundsjön - Östafors",
		Trail:      "SL1",
		Length:     16.0,
		Difficulty: 1,
		File:       "SL1_2.gpx",
	},
	{
		Code:   "SL1-3",
		Name:   "Östafors - Bökestad",
		Trail:  "SL1",
		Length: 9.0,
		File:   "SL1_3.gpx",
	},
	{
		Code:       "SL1-3A",
		Name:       "Boafall - Olofström",
		Trail:      "SL1",
		Length:     12.0,
		Difficulty: 2,
		File:       "SL1-B%25C3%25B6kestad-Olofstr%25C3%25B6m.gpx",
	},
	{
		Code:   "SL1-3A-ansl",
		Name:   "Olofström - Halens camping",
		Trail:  "SL1",
		Length: 1.0,
		File:   "SL1_3Aansl.gpx",
	},
	{
		Code:       "SL1-4",
		Name:       "Bökestad - Brotorpet",
		Trail:      "SL1",
		Length:     13.0,
		Difficulty: 1,
		File:       "SL1_4.gpx",
	},
	{
		Code:       "SL1-5",
		Name:       "Brotorpet - Vesslarp",
		Trail:      "SL1",
		Length:     13.0,
		Difficulty: 1,
		File:       "SL1_5.gpx",
	},
	{
		Code:       "SL1-6",
		Name:       "Vesslarp - Glimåkra",
		Trail:      "SL1",
		Length:     19.0,
		Difficulty: 1,
		File:       "SL1_6.gpx",
	},
	{
		Code:       "SL1-6A",
		Name:       "Breanäsleden",
		Trail:      "SL1",
		Length:     19.0,
		Difficulty: 1,
		File:       "SL1-6A.gpx",
	},
	{
		Code:       "SL1-6A-ansl",
		Name:       "Breanäsleden - Sibbhult",
		Trail:      "SL1",
		Length:     5.0,
		Difficulty: 1,
		File:       "SL1-6Aansl.gpx",
	},
	{
		Code:   "SL1-7",
		Name:   "Glimåkra - Osby",
		Trail:  "SL1",
		Length: 18.0,
		File:   "SL1_7.gpx",
	},
	{
		Code:   "SL1-8",
		Name:   "Osby - Verum",
		Trail:  "SL1",
		Length: 24.0,
		File:   "SL1_8.gpx",
	},
	{
		Code:       "SL1-9",
		Name:       "Verum - Vittsjö",
		Trail:      "SL1",
		Length:     12.0,
		Difficulty: 1,
		File:       "SL1_9.gpx",
	},
	{
		Code:   "SL1-10",
		Name:   "Vittsjö - Hårsjö",
		Trail:  "SL1",
		Length: 8.0,
		File:   "SL1_10.gpx",
	},
	{
		Code:   "SL1-11",
		Name:   "Hårsjö - Lärkesholm",
		Trail:  "SL1",
		Length: 20.0,
		File:   "SL1_11.gpx",
	},
	{
		Code:       "SL1-12",
		Name:       "Lärkesholm - Bjärabygget",
		Trail:      "SL1",
		Length:     13.0,
		Difficulty: 1,
		File:       "SL1_12.gpx",
	},
	{
		Code:       "SL1-12A",
		Name:       "Lärkesholm - Grytåsa",
		Trail:      "SL1",
		Length:     17.0,
		Difficulty: 1,
		File:       "SL1_12A.gpx",
	},
	{
		Code:       "SL1-13",
		Name:       "Bjärabygget - Koarp",
		Trail:      "SL1",
		Length:     17.0,
		Difficulty: 1,
		File:       "SL1_13.gpx",
	},
	{
		Code:       "SL1-14",
		Name:       "Koarp - Brammarp",
		Trail:      "SL1",
		Length:     17.0,
		Difficulty: 1,
		File:       "SL1_14.gpx",
	},
	{
		Code:   "SL1-15",
		Name:   "Brammarp - Båstad",
		Trail:  "SL1",
		Length: 12.0,
		File:   "SL1_15.gpx",
	},
	{
		Code:       "SL1-16",
		Name:       "Båstad - Knösen",
		Trail:      "SL1",
		Length:     16.0,
		Difficulty: 2,
		File:       "SL1_16.gpx",
	},
	{
		Code:       "SL1-17",
		Name:       "Knösen - Torekov",
		Trail:      "SL1",
		Length:     8.0,
		Difficulty: 1,
		File:       "SL1_17.gpx",
	},
	{
		Code:       "SL1-18",
		Name:       "Torekov - Vejbystrand",
		Trail:      "SL1",
		Length:     20.0,
		Difficulty: 1,
		File:       "SL1_18.gpx",
	},
	{
		Code:   "SL1-19",
		Name:   "Vejbystrand - Hålehall",
		Trail:  "SL1",
		Length: 11.0,
		File:   "SL1_19.gpx",
	},
	{
		Code:   "SL1-20",
		Name:   "Hålehall - Örlid",
		Trail:  "SL1",
		Length: 9.0,
		File:   "SL1_20.gpx",
	},
	{
		Code:   "SL1-21",
		Name:   "Ängelsbäcksstrand - Båstad",
		Trail:  "SL1",
		Length: 14.0,
		File:   "SL1_21.gpx",
	},
	{
		Code:   "SL1-22",
		Name:   "Vejbystrand - Ängelholm",
		Trail:  "SL1",
		Length: 13.0,
		File:   "SL1_22.gpx",
	},
}
