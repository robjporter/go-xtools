package xicons

import "fmt"

type Icons struct{}
type iconstyle [1]string

func (i *Icons) GetIconStyles() []string {
	var tmp []string
	for k := range iconstyles {
		tmp = append(tmp, k)
	}
	return tmp
}

func (i *Icons) PrintIconStyles() {
	var tmp []string
	for k := range iconstyles {
		tmp = append(tmp, k)
	}
	fmt.Print(tmp)
}

func New() *Icons {
	return &Icons{}
}

func (i *Icons) GetIcon(name string) string {
	t := iconstyle{}
	if iconstyles[name] != t {
		return iconstyles[name][0]
	}
	return ""
}

func (i *Icons) PrintIcon(name string) {
	t := iconstyle{}
	if iconstyles[name] != t {
		fmt.Print(iconstyles[name][0])
	}
	return
}

var iconstyles = map[string]iconstyle{
	"tick": {
		"✔",
	},
	"cross": {
		"✖",
	},
	"star": {
		"★",
	},
	"square": {
		"▇",
	},
	"smallsquare": {
		"◻",
	},
	"smallsquarefilled": {
		"◼",
	},
	"circle": {
		"◯",
	},
	"circlefilled": {
		"◉",
	},
	"circuledotted": {
		"◌",
	},
	"circledouble": {
		"◎",
	},
	"circlecircle": {
		"ⓞ",
	},
	"circlecrossed": {
		"ⓧ",
	},
	"circlepipe": {
		"ⓘ",
	},
	"bullet": {
		"●",
	},
	"checkoff": {
		"☐",
	},
	"checkboxon": {
		"☒",
	},
	"checkboxcircleon": {
		"ⓧ",
	},
	"checkboxcircleoff": {
		"ⓘ",
	},
	"fancyquestionmark": {
		"❓",
	},
	"neq": {
		"≠",
	},
	"geq": {
		"≥",
	},
	"leq": {
		"≤",
	},
	"times": {
		"×",
	},
	"upperblock1": {
		"▔",
	},
	"upperblock4": {
		"▀",
	},
	"lowerblock1": {
		"▁",
	},
	"lowerblock2": {
		"▂",
	},
	"lowerblock3": {
		"▃",
	},
	"lowerblock4": {
		"▄",
	},
	"lowerblock5": {
		"▅",
	},
	"lowerblock6": {
		"▆",
	},
	"lowerblock7": {
		"▇",
	},
	"lowerblock8": {
		"█",
	},
	"fullblock": {
		"█",
	},
	"line": {
		"─",
	},
	"doubleline": {
		"═",
	},
	"ellipsis": {
		"…",
	},
	"continue": {
		"…",
	},
	"pointer": {
		"❯",
	},
	"info": {
		"ℹ",
	},
	"warning": {
		"⚠",
	},
	"menu": {
		"☰",
	},
	"smiley": {
		"☺",
	},
	"mustache": {
		"෴",
	},
	"heart": {
		"♥",
	},
	"arrowup": {
		"↑",
	},
	"arrowdown": {
		"↓",
	},
	"arrowleft": {
		"←",
	},
	"arrowright": {
		"→",
	},
	"radioon": {
		"◉",
	},
	"radiooff": {
		"◯",
	},
	"sup0": {
		"⁰",
	},
	"sup1": {
		"¹",
	},
	"sup2": {
		"²",
	},
	"sup3": {
		"³",
	},
	"sup4": {
		"⁴",
	},
	"sup5": {
		"⁵",
	},
	"sup6": {
		"⁶",
	},
	"sup7": {
		"⁷",
	},
	"sup8": {
		"⁸",
	},
	"sup9": {
		"⁹",
	},
	"supminus": {
		"⁻",
	},
	"supplus": {
		"⁺",
	},
	"play": {
		"▶",
	},
	"stop": {
		"■",
	},
	"record": {
		"●",
	},
	"copyright": {
		"©",
	},
	"registredtm": {
		"®",
	},
	"trademark": {
		"™",
	},
	"celsius": {
		"℃",
	},
	"fahrenheit": {
		"℉",
	},
	"sunshine": {
		"☀",
	},
	"cloudy": {
		"☁",
	},
	"rain": {
		"☂",
	},
	"snow": {
		"☃",
	},
	"starblack": {
		"★",
	},
	"starwhite": {
		"☆",
	},
	"death": {
		"☠",
	},
	"heartoutline": {
		"♡",
	},
	"diamond": {
		"♢",
	},
	"spade": {
		"♤",
	},
	"club": {
		"♧",
	},
	"arrowbackward": {
		"◀️",
	},
	"arrowdownwards": {
		"🔽",
	},
	"arrowforward": {
		"▶️",
	},
	"arrowupwards": {
		"🔼",
	},
	"arrowdown2": {
		"⬇️",
	},
	"arrowleft2": {
		"⬅️",
	},
	"arrowright2": {
		"➡️",
	},
	"arrowup2": {
		"⬆️",
	},
	"asterisk": {
		"*⃣",
	},
	"ball8": {
		"🎱",
	},
	"alien": {
		"👽",
	},
	"apple": {
		"🍏",
	},
	"anchor": {
		"⚓",
	},
	"angel": {
		"👼",
	},
	"angry": {
		"😠",
	},
	"amazed": {
		"😲",
	},
	"baby": {
		"👶",
	},
	"babychick": {
		"🐤",
	},
	"balloon": {
		"🎈",
	},
	"banana": {
		"🍌",
	},
	"bank": {
		"🏦",
	},
	"bat": {
		"🦇",
	},
	"battery": {
		"🔋",
	},
	"beer": {
		"🍺",
	},
	"beers": {
		"🍻",
	},
	"bell": {
		"🔔",
	},
	"bike": {
		"🚲",
	},
	"blackflag": {
		"🏴",
	},
	"blackheart": {
		"🖤",
	},
	"blossom": {
		"🌼",
	},
	"blossom2": {
		"🌸",
	},
	"blush": {
		"😊",
	},
	"bomb": {
		"💣",
	},
	"briefcase": {
		"💼",
	},
	"bulb": {
		"💡",
	},
	"bus": {
		"🚌",
	},
	"butterfly": {
		"🦋",
	},
	"cake": {
		"🍰",
	},
	"calendar1": {
		"📆",
	},
	"calendar2": {
		"📅",
	},
	"callme": {
		"🤙",
	},
	"camel": {
		"🐫",
	},
	"camera": {
		"📷",
	},
	"camping": {
		"🏕",
	},
	"candy": {
		"🍬",
	},
	"canoe": {
		"🛶",
	},
	"car1": {
		"🚙",
	},
	"cat": {
		"🐱",
	},
	"cd": {
		"💿",
	},
	"chains": {
		"⛓",
	},
	"champagne": {
		"🍾",
	},
	"check": {
		"✅",
	},
	"chicken": {
		"🐔",
	},
	"christmastree": {
		"🎄",
	},
	"clap": {
		"👏",
	},
	"clapper": {
		"🎬",
	},
	"clipboard": {
		"📋",
	},
	"clock": {
		"⏱",
	},
	"clownface": {
		"🤡",
	},
	"cloud": {
		"☁️",
	},
	"cloudlighting": {
		"🌩",
	},
	"cloudlightingrain": {
		"⛈",
	},
	"cloudrain": {
		"🌧",
	},
	"cloudsnow": {
		"🌨",
	},
	"clubs": {
		"♣️",
	},
	"cocktail": {
		"🍸",
	},
	"confused": {
		"😕",
	},
	"construction": {
		"🏗",
	},
	"couple": {
		"👫",
	},
	"cow": {
		"🐮",
	},
	"crab": {
		"🦀",
	},
	"creditcard": {
		"💳",
	},
	"crocodile": {
		"🐊",
	},
	"cross2": {
		"❌",
	},
	"crossedfingers": {
		"🤞",
	},
}

/*
crown":{
"👑",
},
cry":{
"😢",
},
sunglasses":{
"🕶",
},
dart":{
"🎯",
},
desktop":{
"🖥",
},
diamonds":{
"♦️",
},
disappointed":{
"😞",
},
dissapointedrelived ":{
"😥",
},
dizzy":{
"😵",
},
dog":{
"🐶",
},
dolphin":{
"🐬",
},
door":{
"🚪",
},
doughnut":{
"🍩",
},
dove":{
"🕊",
},
dragon":{
"🐉",
},
dress":{
"👗",
},
drop":{
"💧",
},
drooling":{
"🤤",
},
duck":{
"🦆",
},
eagle":{
"🦅",
},
ear":{
"👂",
},
earth":{
"🌍",
},
egg":{
"🥚",
},
elephant":{
"🐘",
},
email":{
"📧",
},
expressionless":{
"😑",
},
eye":{
"👁",
},
eyes":{
"👀",
},
fire":{
"🔥",
},
glasses":{
"👓",
},
fear":{
"😨",
},
filecabinet":{
"🗄",
},
finishflag":{
"🏁",
},
fireengine":{
"🚒",
},
fish":{
"🐟",
},
fist":{
"✊",
},
foggy":{
"🌁",
},
football":{
"⚽",
},
footprints":{
"👣",
},
frog":{
"🐸",
},
frowing":{
"😦",
},
gem":{
"💎",
},
ghost":{
"👻",
},
gift":{
"🎁",
},
goat":{
"🐐",
},
golf":{
"⛳",
},
gorilla":{
"🦍",
},
grin":{
"😁",
},
guard":{
"💂",
},
gun":{
"🔫",
},
hamburger":{
"🍔",
},
handshake":{
"🤝",
},
head":{
"👤",
},
headphones":{
"🎧",
},
heads":{
"👥",
},
hearts":{
"♥️",
},
hearteyes":{
"😍",
},
hourglass":{
"⌛",
},
house":{
"🏠",
},
hugs":{
"🤗",
},
hushed":{
"😯",
},
innocent":{
"😇",
},
joy":{
"😂",
},
joycat":{
"😹",
},
key1":{
"🔑",
},
key2":{
"🗝",
},
kiss":{
"💋",
},
kissingheart":{
"😘",
},
laughing":{
"😆",
},
lion":{
"🦁",
},
lock1":{
"🔐",
},
lock2":{
"🔒",
},
mask":{
"😷",
},
medal1":{
"🎖",
},
medal2":{
"🏅",
},
metal":{
"🤘",
},
microphone":{
"🎤",
},
monkey":{
"🐒",
},
moon":{
"🌙",
},
motar":{
"🎓",
},
mouse":{
"🐭",
},
muscle":{
"💪",
},
mushroom":{
"🍄",
},
neutral":{
"😐",
},
nomouth":{
"😶",
},
ocean":{
"🌊",
},
ok":{
"👌",
},
openhand":{
"👐",
},
owl":{
"🦉",
},
panda":{
"🐼",
},
pear":{
"🍐",
},
pepper":{
"🌶",
},
phone1":{
"☎️",
},
phone2":{
"📞",
},
piano":{
"🎹",
},
pig":{
"🐷",
},
pizza":{
"🍕",
},
plane":{
"✈️",
},
pointdown2":{
"👇",
},
pointleft2":{
"👈",
},
pointright2":{
"👉",
},
pointup2":{
"👆",
},
police":{
"👮",
},
policecar":{
"🚓",
},
pray":{
"🙏",
},
printer":{
"🖨",
},
punch":{
"👊",
},
rabbit":{
"🐰",
},
racecar":{
"🏎",
},
rage":{
"😡",
},
rainbow":{
"🌈",
},
raisedbackhand":{
"🤚",
},
raisedhand":{
"✋",
},
raisedhandsprayed":{
"🖐",
},
raisedhands":{
"🙌",
},
relaxed":{
"☺️",
},
relived":{
"😌",
},
ribbon":{
"🎗",
},
robot":{
"🤖",
},
rocket":{
"🚀",
},
rofl":{
"🤣",
},
rolleyes":{
"🙄",
},
rose":{
"🌹",
},
rugby":{
"🏈",
},
rugby2":{
"🏉",
},
santa":{
"🎅",
},
scream":{
"😱",
},
seenoevil":{
"🙈",
},
shamrock":{
"☘",
},
sheep":{
"🐑",
},
shell":{
"🐚",
},
shield":{
"🛡",
},
ship":{
"🚢",
},
skull":{
"💀",
},
sleeping":{
"😴",
},
sleepy":{
"😪",
},
smile":{
"😄",
},
smiley2":{
"😃",
},
smirk":{
"😏",
},
snail":{
"🐌",
},
sneeze":{
"🤧",
},
snowflake":{
"❄️",
},
snowman":{
"⛄",
},
spades":{
"♠️",
},
speach":{
"💬",
},
speaknoevil":{
"🙊",
},
sob":{
"😭",
},
star":{
"⭐",
},
star2":{
"🌟",
},
stop":{
"⏹",
},
strawberry":{
"🍓",
},
sunbehindcloud":{
"🌤",
},
sunbehindlargecloud ":{
"🌥",
},
sunbehindraincloud":{
"🌦",
},
sunflower":{
"🌻",
},
sunglasses2":{
"😎",
},
sunny":{
"☀️",
},
swimming":{
"🏊",
},
tangerine":{
"🍊",
},
taxi":{
"🚕",
},
tea":{
"🍵",
},
tennis":{
"🎾",
},
tent":{
"⛺",
},
ticket1":{
"🎫",
},
ticket2":{
"🎟",
},
thinking":{
"🤔",
},
thumbsup":{
"👍",
},
thumbsdown":{
"👎",
},
tiger":{
"🐯",
},
tired":{
"😫",
},
tomato":{
"🍅",
},
tongueout":{
"😛",
},
train":{
"🚋",
},
tree":{
"🌳",
},
trolley":{
"🛒",
},
trophy":{
"🏆",
},
truck":{
"🚚",
},
upsideface":{
"🙃",
},
vulcan":{
"🖖",
},
wave":{
"👋",
},
weary":{
"😩",
},
whale":{
"🐳",
},
wine":{
"🍷",
},
wink":{
"😉",
},
zap":{
"⚡",
},
zzzz":{
"💤",
},
}
*/
