package main

import (
	"fmt"
	"strconv"
	"strings"
)

type colorMajor struct {
	name string
	code string
}

func (c *colorMajor) String() string {
	return fmt.Sprintf("[%s]%s", c.code, c.name)
}

// expecting h is #000000 type.
func toRGB(h string) []int {
	// drop first letter #
	h = strings.Replace(h, "#", "", -1)
	if len(h) != 6 {
		return []int{0, 0, 0}
	}
	r := h[0:2]
	g := h[2:4]
	b := h[4:6]
	ri, _ := strconv.ParseInt("0x"+r, 0, 9)
	gi, _ := strconv.ParseInt("0x"+g, 0, 9)
	bi, _ := strconv.ParseInt("0x"+b, 0, 9)
	return []int{int(ri), int(gi), int(bi)}
}

func toTextRGB(h string) string {
	rgb := toRGB(h)
	return fmt.Sprintf("%d, %d, %d", rgb[0], rgb[1], rgb[2])
}

var colorMajors = []*colorMajor{
	{"red", "#ef5350"},
	{"pink", "#ec407a"},
	{"purple", "#ab47bc"},
	{"deep purple", "#7e57c2"},
	{"indigo", "#5c6bc0"},
	{"blue", "#42a5f5"},
	{"light blue", "#29b6f6"},
	{"cyan", "#26c6da"},
	{"teal", "#26a69a"},
	{"green", "#66bb6a"},
	{"light green", "#9ccc65"},
	{"lime", "#d4e157"},
	{"yellow", "#ffee58"},
	{"amber", "#ffca28"},
	{"orange", "#ffa726"},
	{"deep orange", "#ff7043"},
	{"brown", "#8d6e63"},
	{"grey", "#bdbdbd"},
	{"blue grey", "#78909c"},
}

var colorMap = map[string][]string{
	"red":         cmred,
	"pink":        cmpink,
	"purple":      cmpurple,
	"deep purple": cmdeeppurple,
	"indigo":      cmindigo,
	"blue":        cmblue,
	"light blue":  cmlightblue,
	"cyan":        cmcyan,
	"teal":        cmteal,
	"green":       cmgreen,
	"light green": cmlightgreen,
	"lime":        cmlime,
	"yellow":      cmyellow,
	"amber":       cmamber,
	"orange":      cmorange,
	"deep orange": cmdeeporange,
	"brown":       cmbrown,
	"grey":        cmgrey,
	"blue grey":   cmbluegrey,
}

var cmred = []string{
	"[#ffebee]50",
	"[#ffcdd2]100",
	"[#ef9a9a]200",
	"[#e57373]300",
	"[#ef5350]400",
	"[#f44336]500",
	"[#e53935]600",
	"[#d32f2f]700",
	"[#c62828]800",
	"[#b71c1c]900",
	"[#ff8a80]a100",
	"[#ff5252]a200",
	"[#ff1744]a400",
	"[#d50000]a700",
}
var cmpink = []string{
	"[#fce4ec]50",
	"[#f8bbd0]100",
	"[#f48fb1]200",
	"[#f06292]300",
	"[#ec407a]400",
	"[#e91e63]500",
	"[#d81b60]600",
	"[#c2185b]700",
	"[#ad1457]800",
	"[#880e4f]900",
	"[#ff80ab]a100",
	"[#ff4081]a200",
	"[#f50057]a400",
	"[#c51162]a700",
}
var cmpurple = []string{
	"[#f3e5f5]50",
	"[#e1bee7]100",
	"[#ce93d8]200",
	"[#ba68c8]300",
	"[#ab47bc]400",
	"[#9c27b0]500",
	"[#8e24aa]600",
	"[#7b1fa2]700",
	"[#6a1b9a]800",
	"[#4a148c]900",
	"[#ea80fc]a100",
	"[#e040fb]a200",
	"[#d500f9]a400",
	"[#aa00ff]a700",
}
var cmdeeppurple = []string{
	"[#ede7f6]50",
	"[#d1c4e9]100",
	"[#b39ddb]200",
	"[#9575cd]300",
	"[#7e57c2]400",
	"[#673ab7]500",
	"[#5e35b1]600",
	"[#512da8]700",
	"[#4527a0]800",
	"[#311b92]900",
	"[#b388ff]a100",
	"[#7c4dff]a200",
	"[#651fff]a400",
	"[#6200ea]a700",
}
var cmindigo = []string{
	"[#e8eaf6]50",
	"[#c5cae9]100",
	"[#9fa8da]200",
	"[#7986cb]300",
	"[#5c6bc0]400",
	"[#3f51b5]500",
	"[#3949ab]600",
	"[#303f9f]700",
	"[#283593]800",
	"[#1a237e]900",
	"[#8c9eff]a100",
	"[#536dfe]a200",
	"[#3d5afe]a400",
	"[#304ffe]a700",
}
var cmblue = []string{
	"[#e3f2fd]50",
	"[#bbdefb]100",
	"[#90caf9]200",
	"[#64b5f6]300",
	"[#42a5f5]400",
	"[#2196f3]500",
	"[#1e88e5]600",
	"[#1976d2]700",
	"[#1565c0]800",
	"[#0d47a1]900",
	"[#82b1ff]a100",
	"[#448aff]a200",
	"[#2979ff]a400",
	"[#2962ff]a700",
}
var cmlightblue = []string{
	"[#e1f5fe]50",
	"[#b3e5fc]100",
	"[#81d4fa]200",
	"[#4fc3f7]300",
	"[#29b6f6]400",
	"[#03a9f4]500",
	"[#039be5]600",
	"[#0288d1]700",
	"[#0277bd]800",
	"[#01579b]900",
	"[#80d8ff]a100",
	"[#40c4ff]a200",
	"[#00b0ff]a400",
	"[#0091ea]a700",
}
var cmcyan = []string{
	"[#e0f7fa]50",
	"[#b2ebf2]100",
	"[#80deea]200",
	"[#4dd0e1]300",
	"[#26c6da]400",
	"[#00bcd4]500",
	"[#00acc1]600",
	"[#0097a7]700",
	"[#00838f]800",
	"[#006064]900",
	"[#84ffff]a100",
	"[#18ffff]a200",
	"[#00e5ff]a400",
	"[#00b8d4]a700",
}
var cmteal = []string{
	"[#e0f2f1]50",
	"[#b2dfdb]100",
	"[#80cbc4]200",
	"[#4db6ac]300",
	"[#26a69a]400",
	"[#009688]500",
	"[#00897b]600",
	"[#00796b]700",
	"[#00695c]800",
	"[#004d40]900",
	"[#a7ffeb]a100",
	"[#64ffda]a200",
	"[#1de9b6]a400",
	"[#00bfa5]a700",
}
var cmgreen = []string{
	"[#e8f5e9]50",
	"[#c8e6c9]100",
	"[#a5d6a7]200",
	"[#81c784]300",
	"[#66bb6a]400",
	"[#4caf50]500",
	"[#43a047]600",
	"[#388e3c]700",
	"[#2e7d32]800",
	"[#1b5e20]900",
	"[#b9f6ca]a100",
	"[#69f0ae]a200",
	"[#00e676]a400",
	"[#00c853]a700",
}
var cmlightgreen = []string{
	"[#f1f8e9]50",
	"[#dcedc8]100",
	"[#c5e1a5]200",
	"[#aed581]300",
	"[#9ccc65]400",
	"[#8bc34a]500",
	"[#7cb342]600",
	"[#689f38]700",
	"[#558b2f]800",
	"[#33691e]900",
	"[#ccff90]a100",
	"[#b2ff59]a200",
	"[#76ff03]a400",
	"[#64dd17]a700",
}
var cmlime = []string{
	"[#f9fbe7]50",
	"[#f0f4c3]100",
	"[#e6ee9c]200",
	"[#dce775]300",
	"[#d4e157]400",
	"[#cddc39]500",
	"[#c0ca33]600",
	"[#afb42b]700",
	"[#9e9d24]800",
	"[#827717]900",
	"[#f4ff81]a100",
	"[#eeff41]a200",
	"[#c6ff00]a400",
	"[#aeea00]a700",
}
var cmyellow = []string{
	"[#fffde7]50",
	"[#fff9c4]100",
	"[#fff59d]200",
	"[#fff176]300",
	"[#ffee58]400",
	"[#ffeb3b]500",
	"[#fdd835]600",
	"[#fbc02d]700",
	"[#f9a825]800",
	"[#f57f17]900",
	"[#ffff8d]a100",
	"[#ffff00]a200",
	"[#ffea00]a400",
	"[#ffd600]a700",
}
var cmamber = []string{
	"[#fff8e1]50",
	"[#ffecb3]100",
	"[#ffe082]200",
	"[#ffd54f]300",
	"[#ffca28]400",
	"[#ffc107]500",
	"[#ffb300]600",
	"[#ffa000]700",
	"[#ff8f00]800",
	"[#ff6f00]900",
	"[#ffe57f]a100",
	"[#ffd740]a200",
	"[#ffc400]a400",
	"[#ffab00]a700",
}
var cmorange = []string{
	"[#fff3e0]50",
	"[#ffe0b2]100",
	"[#ffcc80]200",
	"[#ffb74d]300",
	"[#ffa726]400",
	"[#ff9800]500",
	"[#fb8c00]600",
	"[#f57c00]700",
	"[#ef6c00]800",
	"[#e65100]900",
	"[#ffd180]a100",
	"[#ffab40]a200",
	"[#ff9100]a400",
	"[#ff6d00]a700",
}
var cmdeeporange = []string{
	"[#fbe9e7]50",
	"[#ffccbc]100",
	"[#ffab91]200",
	"[#ff8a65]300",
	"[#ff7043]400",
	"[#ff5722]500",
	"[#f4511e]600",
	"[#e64a19]700",
	"[#d84315]800",
	"[#bf360c]900",
	"[#ff9e80]a100",
	"[#ff6e40]a200",
	"[#ff3d00]a400",
	"[#dd2c00]a700",
}
var cmbrown = []string{
	"[#efebe9]50",
	"[#d7ccc8]100",
	"[#bcaaa4]200",
	"[#a1887f]300",
	"[#8d6e63]400",
	"[#795548]500",
	"[#6d4c41]600",
	"[#5d4037]700",
	"[#4e342e]800",
	"[#3e2723]900",
}
var cmgrey = []string{
	"[#fafafa]50",
	"[#f5f5f5]100",
	"[#eeeeee]200",
	"[#e0e0e0]300",
	"[#bdbdbd]400",
	"[#9e9e9e]500",
	"[#757575]600",
	"[#616161]700",
	"[#424242]800",
	"[#212121]900",
}
var cmbluegrey = []string{
	"[#eceff1]50",
	"[#cfd8dc]100",
	"[#b0bec5]200",
	"[#90a4ae]300",
	"[#78909c]400",
	"[#607d8b]500",
	"[#546e7a]600",
	"[#455a64]700",
	"[#37474f]800",
	"[#263238]900",
}
