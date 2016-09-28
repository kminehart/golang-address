package address

type Address struct {
	House           int
	StreetDirection string
	StreetName      string
	StreetType      string
	SuiteType       string
	SuiteNumber     int
}

// http://pe.usps.gov/text/pub28/28apc_002.htm
// Include a key that matches the value (ex; "avn": "avn") to make validating valid strings easy.
// If these change, pull requests are welcome!
var StreetTypeAbbreviations map[string]string = map[string]string{
	"alley": "aly",
	"allee": "aly",
	"ally":  "aly",
	"aly":   "aly",

	"anex":  "anx",
	"annex": "anx",
	"annx":  "anx",
	"anx":   "anx",

	"arcade": "arc",
	"arc":    "arc",

	"avenue": "ave",
	"av":     "ave",
	"aven":   "ave",
	"avn":    "ave",
	"anvue":  "ave",
	"ave":    "ave",

	"bayou": "byu",
	"bayoo": "byu",
	"byu":   "byu",

	"beach": "bch",
	"bch":   "bch",

	"bend": "bnd",
	"bnd":  "bnd",

	"bluff": "blf",
	"bluf":  "blf",
	"blf":   "blf",

	"bottom": "btm",
	"bot":    "btm",
	"bottm":  "btm",
	"btm":    "btm",

	"boulevard": "blvd",
	"boulv":     "blvd",
	"boul":      "blvd",
	"blvd":      "blvd",

	"branch": "br",
	"brnch":  "br",
	"br":     "br",

	"bridge": "brg",
	"brdge":  "brg",
	"brg":    "brg",

	"brook": "brk",
	"brk":   "brk",

	"brooks": "brks",
	"brks":   "brks",

	"burg": "bg",
	"bg":   "bg",

	"burgs": "bgs",
	"bgs":   "bgs",

	"bypass": "byp",
	"bypa":   "byp",
	"bypas":  "byp",
	"byps":   "byp",
	"byp":    "byp",

	"camp": "cp",
	"cmp":  "cp",
	"cp":   "cp",

	"canyon": "cyn",
	"canyn":  "cyn",
	"cnyn":   "cyn",
	"cyn":    "cyn",

	"cape": "cpe",
	"cpe":  "cpe",

	"causeway": "cswy",
	"causwa":   "cswy",
	"cswy":     "cswy",

	"center": "ctr",
	"cen":    "ctr",
	"cent":   "ctr",
	"centr":  "ctr",
	"centre": "ctr",
	"cnter":  "ctr",
	"ctr":    "ctr",

	"centers": "ctrs",
	"ctrs":    "ctrs",

	"circle": "cir",
	"circ":   "cir",
	"circl":  "cir",
	"crcl":   "cir",
	"crcle":  "cir",
	"cir":    "cir",

	"cliff": "clf",
	"clf":   "clf",

	"cliffs": "clfs",
	"clfs":   "clfs",

	"club": "clb",
	"clb":  "clb",

	"common": "cmn",
	"cmn":    "cmn",

	"commons": "cmns",
	"cmns":    "cmns",

	"corner": "cor",
	"cor":    "cor",

	"corners": "cors",
	"cors":    "cors",

	"course": "crse",
	"crse":   "crse",

	"court": "ct",
	"ct":    "ct",

	"courts": "cts",
	"cts":    "cts",

	"cove": "cv",
	"cv":   "cv",

	"coves": "cvs",
	"cvs":   "cvs",

	"creek": "crk",
	"crk":   "crk",

	"crescent": "cres",
	"crsent":   "cres",
	"crsnt":    "cres",
	"cres":     "cres",

	"crest": "crst",
	"crst":  "crst",

	"crossing": "xing",
	"crssing":  "xing",
	"xing":     "xing",

	"crossroad": "xrd",
	"xrd":       "xrd",

	"crossroads": "xrds",
	"xrds":       "xrds",

	"curve": "curv",
	"curv":  "curv",

	"dale": "dl",
	"dl":   "dl",

	"dam": "dm",
	"dm":  "dm",

	"divide": "dv",
	"div":    "dv",
	"dvd":    "dv",
	"dv":     "dv",

	"drive": "dr",
	"driv":  "dr",
	"drv":   "dr",
	"dr":    "dr",

	"drives": "drs",
	"drs":    "drs",

	"estate": "est",
	"est":    "est",

	"estates": "ests",
	"ests":    "ests",

	"expressway": "expy",
	"expr":       "expy",
	"express":    "expy",
	"expw":       "expy",
	"expy":       "expy",

	"extension": "ext",
	"extn":      "ext",
	"extnsn":    "ext",
	"ext":       "ext",

	"fall": "fall",

	"falls": "fls",
	"fls":   "fls",

	"ferry": "fry",
	"frry":  "fry",
	"fry":   "fry",

	"field": "fld",
	"fld":   "fld",

	"fields": "flds",
	"flds":   "flds",

	"flat": "flt",
	"flt":  "flt",

	"flats": "flts",
	"flts":  "flts",

	"ford": "frd",
	"frd":  "frd",

	"fords": "frds",
	"frds":  "frds",

	"forest": "frst",
	"frst":   "frst",

	"forge": "frg",
	"forg":  "frg",
	"frg":   "frg",

	"forges": "frgs",
	"frgs":   "frgs",

	"fork": "frk",
	"frk":  "frk",

	"fort": "ft",
	"frt":  "ft",
	"ft":   "ft",

	"freeway": "fwy",
	"freewy":  "fwy",
	"frway":   "fwy",
	"frwy":    "fwy",
	"fwy":     "fwy",

	"garden": "gdn",
	"gardn":  "gdn",
	"grden":  "gdn",
	"grdn":   "gdn",
	"gdn":    "gdn",

	"gardens": "gdns",
	"grdns":   "gdns",
	"gdns":    "gdns",

	"gateway": "gtwy",
	"gatewy":  "gtwy",
	"gatway":  "gtwy",
	"gtway":   "gtwy",
	"gtwy":    "gtwy",

	"glen": "gln",
	"gln":  "gln",

	"glens": "glns",
	"glns":  "glns",

	"green": "grn",
	"grn":   "grn",

	"greens": "grns",
	"grns":   "grns",

	"grove": "grv",
	"grov":  "grv",
	"grv":   "grv",

	"harbor": "hbr",
	"harb":   "hbr",
	"harbr":  "hbr",
	"hrbr":   "hbr",
	"hbr":    "hbr",

	"harbors": "hbrs",
	"hbrs":    "hbrs",

	"haven": "hvn",
	"hvn":   "hvn",

	"heights": "hts",
	"ht":      "hts",
	"hts":     "hts",

	"highway": "hwy",
	"highwy":  "hwy",
	"hiway":   "hwy",
	"hiwy":    "hwy",
	"hway":    "hwy",
	"hwy":     "hwy",

	"hill": "hl",
	"hl":   "hl",

	"hills": "hls",
	"hls":   "hls",

	"hollow":  "holw",
	"hllw":    "holw",
	"hollows": "holw",
	"holws":   "holw",
	"holw":    "holw",

	"hodor": "hodor",

	"inlet": "inlt",
	"inlt":  "inlt",

	"island": "is",
	"islnd":  "is",
	"is":     "is",

	"isle":  "isle",
	"isles": "isle",

	"junction": "jct",
	"jction":   "jct",
	"jctn":     "jct",
	"junctn":   "jct",
	"juncton":  "jct",
	"jct":      "jct",

	"junctions": "jcts",
	"jctns":     "jcts",
	"jcts":      "jcts",

	"key": "ky",
	"ky":  "ky",

	"keys": "kys",
	"kys":  "kys",

	"knoll": "knl",
	"knol":  "knl",
	"knl":   "knl",

	"knolls": "knls",
	"knls":   "knls",

	"lake": "lk",
	"lk":   "lk",

	"lakes": "lks",
	"lks":   "lks",

	"land": "land",

	"landing": "lndg",
	"lndng":   "lndg",
	"lndg":    "lndg",

	"lane": "ln",
	"ln":   "ln",

	"light": "lgt",
	"lgt":   "lgt",

	"lights": "lgts",
	"lgts":   "lgts",

	"loaf": "lf",
	"lf":   "lf",

	"lock": "lck",
	"lck":  "lck",

	"locks": "lcks",
	"lcks":  "lcks",

	"lodge": "ldg",
	"ldge":  "ldg",
	"lodg":  "ldg",
	"ldg":   "ldg",

	"loop":  "loop",
	"loops": "loop",

	"mall": "mall",

	"manor": "mnr",
	"mnr":   "mnr",

	"manors": "mnrs",
	"mnrs":   "mnrs",

	// Accoding to that table...:
	// "meadow": "mdw",
	// "mdw": "mdw",

	// But on the following line...
	"meadows": "mdws",
	"mdw":     "mdws",
	"meadow":  "mdws",
	"medows":  "mdws",
	"mdws":    "mdws",

	"mews": "mews",

	"mill": "ml",
	"ml":   "ml",

	"mission": "msn",
	"misn":    "msn",
	"msn":     "msn",

	"motorway": "mtwy",
	"mtwy":     "mtwy",

	"mountain": "mtn",
	"mntain":   "mtn",
	"mt":       "mtn",
	"mountin":  "mtn",
	"mtin":     "mtn",
	"mtn":      "mtn",

	"mountains": "mtns",
	"mtns":      "mtns",

	"neck": "nck",
	"nck":  "nck",

	"orchard": "orch",
	"orchrd":  "orch",
	"orch":    "orch",

	"oval": "ovl",
	"ovl":  "ovl",

	"overpass": "opas",
	"opas":     "opas",

	"park":  "park",
	"parks": "park",
	"prk":   "park",

	"parkway":  "pkwy",
	"parkwy":   "pkwy",
	"pkway":    "pkwy",
	"pky":      "pkwy",
	"parkways": "pkwy",
	"pkwys":    "pkwy",
	"pkwy":     "pkwy",

	"pass": "pass",

	"passage": "psge",
	"psge":    "psge",

	"path":  "path",
	"paths": "paths",

	"pike":  "pike",
	"pikes": "pike",

	"pine": "pne",
	"pne":  "pne",

	"pines": "pnes",
	"pnes":  "pnes",

	"place": "pl",
	"pl":    "pl",

	"plain": "pln",
	"pln":   "pln",

	"plains": "plns",
	"plns":   "plns",

	"plaza": "plz",
	"plza":  "plz",
	"plz":   "plz",

	"point": "pt",
	"pt":    "pt",

	"points": "pts",
	"pts":    "pts",

	"port": "prt",
	"prt":  "prt",

	"ports": "prts",
	"prts":  "prts",

	"prarie": "pr",
	"prr":    "pr",
	"pr":     "pr",

	"radial": "radl",
	"rad":    "radl",
	"radiel": "radl",
	"radl":   "radl",

	"ramp": "ramp",

	"ranch":   "rnch",
	"ranches": "rnch",
	"rnchs":   "rnch",
	"rnch":    "rnch",

	"rapid": "rpd",
	"rpd":   "rpd",

	"ridge": "rdg",
	"rdge":  "rdg",
	"rdg":   "rdg",

	"rest": "rst",
	"rst":  "rst",

	"ridges": "rdgs",
	"rdgs":   "rdgs",

	"route": "rte",
	"rte":   "rte",

	"row": "row",

	"rue": "rue",

	"run": "run",

	"shoal": "shl",
	"shl":   "shl",

	"shoals": "shls",
	"shls":   "shls",

	"shore":  "shr",
	"shores": "shr",
	"shrs":   "shr",
	"shr":    "shr",

	"skyway": "skwy",

	"spring": "spg",
	"spng":   "spg",
	"sprng":  "spg",
	"spg":    "spg",

	"springs": "spgs",
	"spngs":   "spgs",
	"sprngs":  "spgs",
	"spgs":    "spgs",

	"spur":  "spur",
	"spurs": "spur",

	"square": "sq",
	"sqr":    "sq",
	"sqre":   "sq",
	"squ":    "sq",
	"sq":     "sq",

	"squares": "sqs",
	"sqrs":    "sqs",
	"sqs":     "sqs",

	"station": "sta",
	"statn":   "sta",
	"stn":     "sta",
	"sta":     "sta",

	"stravenue": "stra",
	"strav":     "stra",
	"straven":   "stra",
	"stravn":    "stra",
	"strvn":     "stra",
	"strvnue":   "stra",
	"stra":      "stra",

	"stream": "strm",
	"streme": "strm",
	"strm":   "strm",

	"street": "st",
	"strt":   "st",
	"str":    "st",
	"st":     "st",

	"streets": "sts",

	"summit": "smt",
	"sumit":  "smt",
	"sumitt": "smt",
	"smt":    "smt",

	"terrace": "ter",
	"terr":    "ter",
	"ter":     "ter",

	"throughway": "trwy",
	"trwy":       "trwy",

	"trace":  "trce",
	"traces": "trce",
	"trce":   "trce",

	"track":  "trak",
	"tracks": "trak",
	"trk":    "trak",
	"trks":   "trak",
	"trak":   "trak",

	"trailer": "trlr",
	"trlrs":   "trlr",
	"trlr":    "trlr",

	"tunnel":  "tunl",
	"tunls":   "tunl",
	"tunnels": "tunl",
	"tunnl":   "tunl",
	"tunl":    "tunl",

	"turnpike": "tpke",
	"trnp":     "tpke",
	"turnpk":   "tpke",
	"tpke":     "tpke",

	"underpass": "upas",
	"upas":      "upas",

	"union": "un",
	"un":    "un",

	"valley": "vly",
	"vally":  "vly",
	"vlly":   "vly",
	"vly":    "vly",

	"valleys": "vlys",
	"vlys":    "vlys",

	"viaduct": "via",
	"vdct":    "via",
	"viadct":  "via",
	"via":     "via",

	"view": "vw",
	"vw":   "vw",

	"views": "vws",
	"vws":   "vws",

	"village":  "vlg",
	"villag":   "vlg",
	"vill":     "vlg",
	"villg":    "vlg",
	"villiage": "vlg",
	"vlg":      "vlg",

	"villages": "vlgs",
	"vlgs":     "vlgs",

	"ville": "vl",
	"vl":    "vl",

	"vista": "vis",
	"vist":  "vis",
	"vst":   "vis",
	"vsta":  "vis",
	"vis":   "vis",

	"walk":  "walk",
	"walks": "walk",

	"wall": "wall",

	"way": "way",
	"wy":  "way",

	"ways": "ways",

	"well": "wl",
	"wl":   "wl",

	"wells": "wls",
	"wls":   "wls",
}

var SuiteTypeAbbreviations map[string]string = map[string]string{
	"apartment": "apt",
	"#":         "apt",
	"apt":       "apt",

	"building": "bldg",
	"bldg":     "bldg",

	"floor": "fl",
	"fl":    "fl",

	"suite": "ste",
	"ste":   "ste",

	"unit": "unit",

	"room": "rm",
	"rm":   "rm",

	"department": "dept",
	"dept":       "dept",
}

var CardinalDirectionAbbreviations map[string]string = map[string]string{
	"north": "n",
	"n":     "n",

	"northwest": "nw",
	"nw":        "nw",

	"northeast": "ne",
	"ne":        "ne",

	"south": "s",
	"s":     "s",

	"southwest": "sw",
	"sw":        "sw",

	"southeast": "se",
	"se":        "se",

	"east": "e",
	"e":    "e",

	"west": "w",
	"w":    "w",
}
