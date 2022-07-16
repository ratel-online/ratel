package tile

const (
	WAN int = iota
	TIAO
	BING
	FENG
	DRAGON
	SEASON
	HUA
	JOKER
)

var TILE_DATA map[int]map[int]string = map[int]map[int]string{
	WAN: {
		1: "🀇",
		2: "🀈",
		3: "🀉",
		4: "🀊",
		5: "🀋",
		6: "🀌",
		7: "🀍",
		8: "🀎",
		9: "🀏",
	},
	TIAO: {
		1: "🀐",
		2: "🀑",
		3: "🀒",
		4: "🀓",
		5: "🀔",
		6: "🀕",
		7: "🀖",
		8: "🀗",
		9: "🀘",
	},
	BING: {
		1: "🀙",
		2: "🀚",
		3: "🀛",
		4: "🀜",
		5: "🀝",
		6: "🀞",
		7: "🀟",
		8: "🀠",
		9: "🀡",
	},
	FENG: {
		1: "🀀",
		2: "🀁",
		3: "🀂",
		4: "🀃",
	},
	DRAGON: {
		1: "🀅",
		2: "🀄︎",
		3: "🀆",
	},
	SEASON: {
		1: "🀦",
		2: "🀧",
		3: "🀨",
		4: "🀩",
	},
	HUA: {
		1: "🀢",
		2: "🀣",
		3: "🀤",
		4: "🀥",
	},
	JOKER: {
		1: "🀪",
	},
}

// var DATA map[TITLE]map[int]int = map[TITLE]map[int]int{
// 	WAN: {
// 		1: 0x1F007,
// 		2: 0x1F008,
// 		3: 0x1F009,
// 		4: 0x1F00A,
// 		5: 0x1F00B,
// 		6: 0x1F00C,
// 		7: 0x1F00D,
// 		8: 0x1F00E,
// 		9: 0x1F00F,
// 	},
// 	TWIG: {
// 		1: 0x1F010,
// 		2: 0x1F011,
// 		3: 0x1F012,
// 		4: 0x1F013,
// 		5: 0x1F014,
// 		6: 0x1F015,
// 		7: 0x1F016,
// 		8: 0x1F017,
// 		9: 0x1F018,
// 	},
// 	DOT: {
// 		1: 0x1F019,
// 		2: 0x1F01A,
// 		3: 0x1F01B,
// 		4: 0x1F01C,
// 		5: 0x1F01D,
// 		6: 0x1F01E,
// 		7: 0x1F01F,
// 		8: 0x1F020,
// 		9: 0x1F021,
// 	},
// 	WIND: {
// 		1: 0x1F000,
// 		2: 0x1F001,
// 		3: 0x1F002,
// 		4: 0x1F003,
// 	},
// 	DRAGON: {
// 		1: 0x1F004,
// 		2: 0x1F005,
// 		3: 0x1F006,
// 	},
// }
