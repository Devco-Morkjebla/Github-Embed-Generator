package style

import (
	"fmt"
	"regexp"
	"strings"
)

type Styles struct {
	Title,
	Border,
	Background,
	Text,
	Textfont,
	Box string
}

func CheckHex(str map[string]string) Styles {
	var style Styles
	r, _ := regexp.Compile("^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")
	if !r.MatchString(str[style.Border]) {
		style.Border = "000000"
	} else {
		style.Border = str["Border"]
	}
	if !r.MatchString(str[style.Title]) {
		style.Title = "000000"
	} else {
		style.Title = str["Border"]
	}
	if !r.MatchString(str["Background"]) {
		style.Background = "ffffff"
	} else {
		style.Background = str["Border"]
	}
	if !r.MatchString(str["Text"]) {
		style.Text = "000000"
	} else {
		style.Text = str["Text"]
	}
	if !r.MatchString(str["Box"]) {
		style.Box = "dddddd"
	} else {
		style.Box = str["Box"]
	}
	return style
}

func RadialGradient(id string, colors []string) string {
	gradient := []string{
		fmt.Sprintf(`<radialGradient id="%v" gradientUnits="userSpaceOnUse">`, id),
	}

	var offset float64 = 1.0 / float64(cap(colors)-1)
	for i, v := range colors {
		gradient = append(gradient, fmt.Sprintf(`<stop offset="%v" stop-color="%v"/>`, offset*float64(i), v))
	}

	gradient = append(gradient, `</radialGradient>`)
	return strings.Join(gradient, "\n")
}
func LinearGradient(id string, colors []string) string {
	gradient := []string{
		fmt.Sprintf(`<linearGradient x="0" y="0" x2="100" id="%v" gradientUnits="userSpaceOnUse">`, id),
	}
	if cap(colors) < 2 {
		panic(`Gradient must have 2 colors`)
	}
	var offset float64 = 1.0 / float64(cap(colors)-1)
	for i, v := range colors {
		gradient = append(gradient, fmt.Sprintf(`<stop offset="%v" stop-color="%v"/>`, offset*float64(i), v))
	}

	gradient = append(gradient, `</linearGradient>`)
	return strings.Join(gradient, "\n")
}
func HexagonPattern() string {

	return `<pattern id="pattern-hex" x="0" y="0" width="112" height="190" patternUnits="userSpaceOnUse" viewBox="56 -254 112 190">
	<g id="hexagon">
	<path d="M168-127.1c0.5,0,1,0.1,1.3,0.3l53.4,30.5c0.7,0.4,1.3,1.4,1.3,2.2v61c0,0.8-0.6,1.8-1.3,2.2L169.3-0.3 c-0.7,0.4-1.9,0.4-2.6,0l-53.4-30.5c-0.7-0.4-1.3-1.4-1.3-2.2v-61c0-0.8,0.6-1.8,1.3-2.2l53.4-30.5C167-127,167.5-127.1,168-127.1 L168-127.1z"></path>
	<path d="M112-222.5c0.5,0,1,0.1,1.3,0.3l53.4,30.5c0.7,0.4,1.3,1.4,1.3,2.2v61c0,0.8-0.6,1.8-1.3,2.2l-53.4,30.5 c-0.7,0.4-1.9,0.4-2.6,0l-53.4-30.5c-0.7-0.4-1.3-1.4-1.3-2.2v-61c0-0.8,0.6-1.8,1.3-2.2l53.4-30.5 C111-222.4,111.5-222.5,112-222.5L112-222.5z"></path>
	<path d="M168-317.8c0.5,0,1,0.1,1.3,0.3l53.4,30.5c0.7,0.4,1.3,1.4,1.3,2.2v61c0,0.8-0.6,1.8-1.3,2.2L169.3-191 c-0.7,0.4-1.9,0.4-2.6,0l-53.4-30.5c-0.7-0.4-1.3-1.4-1.3-2.2v-61c0-0.8,0.6-1.8,1.3-2.2l53.4-30.5 C167-317.7,167.5-317.8,168-317.8L168-317.8z"></path>
	</g>

	</pattern>`
}
func CubePattern() string {
	return `<pattern id="pattern-cubes" x="0" y="126" patternUnits="userSpaceOnUse" width="126" height="200" viewBox="0 0 10 16"> 
     
	<g id="cube">
		<path class="left-shade" d="M0 0l5 3v5l-5 -3z"></path>
		<path class="right-shade" d="M10 0l-5 3v5l5 -3"></path>
	</g>
   
	  <use x="5" y="8" href="#cube"></use>
	  <use x="-5" y="8" href="#cube"></use>

	</pattern>`
}
func StarPattern() string {
	return `<pattern id="star" viewBox="0,0,10,10" width="10%" height="10%">
	<polygon points="0,0 2,5 0,10 5,8 10,10 8,5 10,0 5,2"/>
	  </pattern>`
}
func StarsFilter() string {
	// feColorMatrix
	//------------------
	//	   R G B A M
	//--------------
	// R | 1 0 0 0 0
	// G | 0 1 0 0 0
	// B | 0 0 1 0 0
	// A | 0 0 0 1 0
	return `<filter id="stars">
	<feTurbulence baseFrequency="0.2"/>
	
	<feColorMatrix values="0 0 0 9 -4
						   0 0 0 9 -4
						   0 0 0 9 -4
						   0 0 0 0 1"/>
		</filter>`
}
func DropShadow() string {
	return `<filter id="filter2_d_0_1" x="406" y="71" width="155" height="154" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
	<feFlood flood-opacity="0" result="BackgroundImageFix"/>
	<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
	<feOffset dy="4"/>
	<feGaussianBlur stdDeviation="2"/>
	<feComposite in2="hardAlpha" operator="out"/>
	<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
	<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_0_1"/>
	<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_0_1" result="shape"/>
	</filter>`
}
func Filter2() string {
	return `<filter id="filter1_d_0_1" x="0" y="0" width="100%" height="100%" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
	<feFlood flood-opacity="0" result="BackgroundImageFix"/>
	<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
	<feOffset dy="4"/>
	<feGaussianBlur stdDeviation="2"/>
	<feComposite in2="hardAlpha" operator="out"/>
	<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
	<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_0_1"/>
	<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_0_1" result="shape"/>
	</filter>`
}
func DropShadowRing1() string {
	return `<filter id="filter0_d_0_1" x="391" y="55" width="185" height="186" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
	<feFlood flood-opacity="0" result="BackgroundImageFix"/>
	<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
	<feOffset dy="4"/>
	<feGaussianBlur stdDeviation="2"/>
	<feComposite in2="hardAlpha" operator="out"/>
	<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
	<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_0_1"/>
	<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_0_1" result="shape"/>
	</filter>`
}
func Blur(amount int) string {
	return fmt.Sprintf(`<filter id="blur%v" x="0" y="0">
	<feGaussianBlur in="SourceGraphic" stdDeviation="%v" />
  </filter>`, amount, amount)
}
func Border() string {
	return `<filter id="border" filterUnits="userSpaceOnUse" x="0" y="0" width="200" height="120">
	<feGaussianBlur in="SourceAlpha" stdDeviation="4" result="blur" />
	<feOffset in="blur" dx="4" dy="4" result="offsetBlur" />
	<feSpecularLighting in="blur" surfaceScale="5" specularConstant=".75" specularExponent="20" lighting-color="#bbbbbb" result="specOut">
	  <fePointLight x="-5000" y="-10000" z="20000" />
	</feSpecularLighting>
	<feComposite in="specOut" in2="SourceAlpha" operator="in" result="specOut" />
	<feComposite in="SourceGraphic" in2="specOut" operator="arithmetic" k1="0" k2="1" k3="1" k4="0" result="litPaint" />
	<feMerge>
	  <feMergeNode in="offsetBlur" />
	  <feMergeNode in="litPaint" />
	</feMerge>
  </filter>`
}
