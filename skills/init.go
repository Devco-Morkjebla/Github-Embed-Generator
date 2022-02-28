package skills

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Skillscard struct {
	Title  string   `json:"title"`
	Skills []string `json:"skills"`
	Styles Styles   `json:"styles"`
	Body   []string `json:"body"`
}

type Styles struct {
	Title      string
	Border     string
	Background string
	Text       string
	Textfont   string
	Box        string
}

func Skills(title string, languages []string, style Styles) Skillscard {

	var languageColor = map[string]string{
		"alpinejs":                       "#83B6C4",
		"alpine abuild":                  "null",
		"swig":                           "null",
		"blitzbasic":                     "null",
		"xbase":                          "#403a40",
		"asymptote":                      "#ff0000",
		"brightscript":                   "null",
		"groovy server pages":            "null",
		"nim":                            "#ffc200",
		"cython":                         "null",
		"v":                              "#4f87c4",
		"mathematica":                    "null",
		"shen":                           "#120F14",
		"dataweave":                      "#003a52",
		"fortran":                        "#4d41b1",
		"graphql":                        "#e10098",
		"abap":                           "#E8274B",
		"asl":                            "null",
		"mercury":                        "#ff2b2b",
		"sas":                            "#B34936",
		"zimpl":                          "null",
		"pov-ray sdl":                    "null",
		"smpl":                           "#c94949",
		"awk":                            "null",
		"nu":                             "#c9df40",
		"yacc":                           "#4B6C4B",
		"erlang":                         "#B83998",
		"inform 7":                       "null",
		"turing":                         "#cf142b",
		"java":                           "#b07219",
		"mtml":                           "#b7e1f4",
		"q":                              "#0040cd",
		"r":                              "#198CE7",
		"xs":                             "null",
		"realbasic":                      "null",
		"webassembly":                    "#04133b",
		"harbour":                        "#0e60e3",
		"haskell":                        "#5e5086",
		"python":                         "#3572A5",
		"txl":                            "null",
		"x10":                            "#4B6BEF",
		"yara":                           "#220000",
		"grammatical framework":          "#ff0000",
		"hlsl":                           "null",
		"nit":                            "#009917",
		"qt script":                      "#00b841",
		"tla":                            "null",
		"gdscript":                       "#355570",
		"mirah":                          "#c7a938",
		"openscad":                       "null",
		"filebench wml":                  "null",
		"livescript":                     "#499886",
		"module management system":       "null",
		"cartocss":                       "null",
		"ncl":                            "#28431f",
		"openedge abl":                   "null",
		"rpc":                            "null",
		"vala":                           "#fbe5cd",
		"emberscript":                    "#FFF4F3",
		"glyph":                          "#c1ac7f",
		"haxe":                           "#df7900",
		"ox":                             "null",
		"cuda":                           "#3A4E3A",
		"elm":                            "#60B5CC",
		"f#":                             "#b845fc",
		"oxygene":                        "#cdd0e3",
		"tsql":                           "#cdd0e3",
		"tailwindcss":                    "#38BDF8",
		"tailwind":                       "#38BDF8",
		"emacs lisp":                     "#c065db",
		"mupad":                          "null",
		"openrc runscript":               "null",
		"webidl":                         "null",
		"component pascal":               "#B0CE4E",
		"idl":                            "#a3522f",
		"raml":                           "#77d9fb",
		"ballerina":                      "#FF5000",
		"mlir":                           "#5EC8DB",
		"xtend":                          "null",
		"fish":                           "null",
		"gnuplot":                        "#f0a9f0",
		"go":                             "#00ADD8",
		"nsis":                           "null",
		"zig":                            "#ec915c",
		"assembly":                       "#6E4C13",
		"dafny":                          "#FFEC25",
		"starlark":                       "#76d275",
		"arc":                            "#aa2afe",
		"blitzmax":                       "#cd6400",
		"coldfusion":                     "#ed2cd6",
		"hy":                             "#7790B2",
		"idris":                          "#b30000",
		"matlab":                         "#e16737",
		"clojure":                        "#db5855",
		"supercollider":                  "#46390b",
		"vcl":                            "#148AA8",
		"ampl":                           "#E6EFBB",
		"csound document":                "null",
		"parrot assembly":                "null",
		"vba":                            "#867db1",
		"papyrus":                        "#6600cc",
		"parrot":                         "#f3ca0a",
		"srecode template":               "#348a34",
		"mirc script":                    "#3d57c3",
		"pogoscript":                     "#d80074",
		"batchfile":                      "#C1F12E",
		"gams":                           "null",
		"lean":                           "null",
		"lookml":                         "#652B81",
		"mql5":                           "#4A76B8",
		"cap'n proto":                    "null",
		"mql4":                           "#62A8D6",
		"ocaml":                          "#3be133",
		"puppet":                         "#302B6D",
		"slim":                           "#2b2b2b",
		"applescript":                    "#101F1F",
		"io":                             "#a9188d",
		"javascript+erb":                 "null",
		"rebol":                          "#358a5b",
		"csound score":                   "null",
		"minid":                          "null",
		"stan":                           "#b2011d",
		"game maker language":            "#71b417",
		"mako":                           "null",
		"omgrofl":                        "#cabbff",
		"xojo":                           "null",
		"eq":                             "#a78649",
		"ooc":                            "#b0b77e",
		"picolisp":                       "null",
		"powerbuilder":                   "#8f0f8d",
		"actionscript":                   "#882B0F",
		"autoit":                         "#1C3552",
		"digital command language":       "null",
		"qml":                            "#44a51c",
		"scss":                           "#c6538c",
		"self":                           "#0579aa",
		"typescript":                     "#2b7489",
		"cycript":                        "null",
		"f*":                             "#572e30",
		"julia":                          "#a270ba",
		"m4sugar":                        "null",
		"vbscript":                       "#15dcdc",
		"m":                              "null",
		"gdb":                            "null",
		"robotframework":                 "null",
		"scheme":                         "#1e4aec",
		"futhark":                        "#5f021f",
		"prolog":                         "#74283c",
		"swift":                          "#ffac45",
		"csound":                         "null",
		"ecl":                            "#8a1267",
		"jolie":                          "#843179",
		"mask":                           "#f97732",
		"max":                            "#c4a79c",
		"purescript":                     "#1D222D",
		"roff":                           "#ecdebe",
		"xslt":                           "#EB8CEB",
		"charity":                        "null",
		"gaml":                           "#FFC766",
		"gentoo eclass":                  "null",
		"kotlin":                         "#A97BFF",
		"sass":                           "#a53b70",
		"smalltalk":                      "#596706",
		"vue":                            "#3FB27F",
		"ada":                            "#02f88c",
		"renderscript":                   "null",
		"smt":                            "null",
		"flux":                           "#88ccff",
		"logos":                          "null",
		"metal":                          "#8f14e9",
		"sieve":                          "null",
		"pascal":                         "#E3F171",
		"raku":                           "#0000fb",
		"html":                           "#e34c26",
		"jsoniq":                         "#40d47e",
		"motorola 68k assembly":          "null",
		"pep8":                           "#C76F5B",
		"uno":                            "#9933cc",
		"vhdl":                           "#adb2cb",
		"cirru":                          "#ccccff",
		"c++":                            "#f34b7d",
		"twig":                           "#c1d026",
		"bison":                          "#6A463F",
		"genie":                          "#fb855d",
		"qmake":                          "null",
		"ragel":                          "#9d5200",
		"tcl":                            "#e4cc98",
		"augeas":                         "null",
		"befunge":                        "null",
		"gentoo ebuild":                  "null",
		"igor pro":                       "#0000cc",
		"lfe":                            "#4C3023",
		"netlinx":                        "#0aa0ff",
		"newlisp":                        "#87AED7",
		"ren'py":                         "#ff7f7f",
		"smarty":                         "null",
		"wisp":                           "#7582D1",
		"open policy agent":              "null",
		"zenscript":                      "#00BCD1",
		"coldfusion cfc":                 "#ed2cd6",
		"d":                              "#ba595e",
		"dogescript":                     "#cca760",
		"jison":                          "null",
		"marko":                          "#42bff2",
		"monkey":                         "null",
		"prisma":                         "#0c344b",
		"rascal":                         "#fffaa0",
		"pike":                           "#005390",
		"faust":                          "#c37240",
		"makefile":                       "#427819",
		"blade":                          "#f7523f",
		"numpy":                          "#9C8AF9",
		"sqf":                            "#3F3F3F",
		"systemverilog":                  "#DAE1C2",
		"tcsh":                           "null",
		"ats":                            "#1ac620",
		"bitbake":                        "null",
		"common workflow language":       "#B5314C",
		"ioke":                           "#078193",
		"vim script":                     "#199f4b",
		"dockerfile":                     "#384d54",
		"docker":                         "#384d54",
		"markdown":                       "#083fa1",
		"chapel":                         "#8dc63f",
		"fantom":                         "#14253c",
		"jison lex":                      "null",
		"krl":                            "#28430A",
		"nix":                            "#7e7eff",
		"stata":                          "null",
		"gap":                            "null",
		"ring":                           "#2D54CB",
		"golo":                           "#88562A",
		"p4":                             "#7055b5",
		"quake":                          "#882233",
		"jasmin":                         "null",
		"pawn":                           "#dbb284",
		"verilog":                        "#b2b7f8",
		"antlr":                          "#9DC3FF",
		"clean":                          "#3F85AF",
		"latte":                          "#f2a542",
		"propeller spin":                 "#7fa2a7",
		"pug":                            "#a86454",
		"wdl":                            "#42f1f4",
		"hack":                           "#878787",
		"moonscript":                     "null",
		"myghty":                         "null",
		"riot":                           "#A71E49",
		"coffeescript":                   "#244776",
		"cweb":                           "null",
		"hiveql":                         "#dce200",
		"sqlpl":                          "null",
		"xquery":                         "#5232e7",
		"java server pages":              "null",
		"modula-3":                       "#223388",
		"urweb":                          "null",
		"alloy":                          "#64C800",
		"clarion":                        "#db901e",
		"cobol":                          "null",
		"sourcepawn":                     "#f69e1d",
		"brainfuck":                      "#2F2530",
		"dm":                             "#447265",
		"jupyter notebook":               "#DA5B0B",
		"netlinx+erb":                    "#747faa",
		"perl":                           "#0298c3",
		"scala":                          "#c22d40",
		"angelscript":                    "#C7D7DC",
		"codeql":                         "null",
		"m4":                             "null",
		"maxscript":                      "#00a6a6",
		"unified parallel c":             "#4e3617",
		"isabelle":                       "#FEFE00",
		"openqasm":                       "#AA70FF",
		"powershell":                     "#012456",
		"gosu":                           "#82937f",
		"grace":                          "null",
		"lua":                            "#000080",
		"rouge":                          "#cc0088",
		"standard ml":                    "#dc566d",
		"ejs":                            "#a91e50",
		"lasso":                          "#999999",
		"svg":                            "#ff9900",
		"zephir":                         "#118f9e",
		"llvm":                           "#185619",
		"muf":                            "null",
		"reason":                         "#ff5847",
		"apl":                            "#5A8164",
		"gherkin":                        "#5B2063",
		"nemerle":                        "#3d3c6e",
		"oz":                             "#fab738",
		"slice":                          "#003fa2",
		"chuck":                          "null",
		"inno setup":                     "null",
		"modula-2":                       "null",
		"racket":                         "#3c5caa",
		"svelte":                         "#ff3e00",
		"mcfunction":                     "#E22837",
		"opa":                            "null",
		"shellsession":                   "null",
		"solidity":                       "#AA6746",
		"1c enterprise":                  "#814CCC",
		"haml":                           "#ece2a9",
		"limbo":                          "null",
		"php":                            "#4F5D95",
		"wollok":                         "#a23738",
		"4d":                             "null",
		"asp.net":                        "#9400ff",
		"hyphy":                          "null",
		"ruby":                           "#701516",
		"handlebars":                     "#f7931e",
		"lsl":                            "#3d9970",
		"modelica":                       "null",
		"red":                            "#f50000",
		"classic asp":                    "#6a40fd",
		"gcc machine description":        "null",
		"meson":                          "#007800",
		"opencl":                         "null",
		"cson":                           "#244776",
		"dart":                           "#00B4AB",
		"loomscript":                     "null",
		"volt":                           "#1F1F1F",
		"frege":                          "#00cafe",
		"rust":                           "#dea584",
		"shaderlab":                      "null",
		"squirrel":                       "#800000",
		"cmake":                          "null",
		"postscript":                     "#da291c",
		"zeek":                           "null",
		"ceylon":                         "#dfa535",
		"click":                          "#E4E6F3",
		"jflex":                          "#DBCA00",
		"objective-c++":                  "#6866fb",
		"lilypond":                       "null",
		"objective-c":                    "#438eff",
		"q#":                             "#fed659",
		"sed":                            "#64b970",
		"coq":                            "null",
		"elixir":                         "#6e4a7e",
		"nextflow":                       "#3ac486",
		"plsql":                          "#dad8d8",
		"unix assembly":                  "null",
		"nasl":                           "null",
		"parrot internal representation": "null",
		"scilab":                         "null",
		"css":                            "#563d7c",
		"filterscript":                   "null",
		"slash":                          "#007eff",
		"pony":                           "null",
		"smali":                          "null",
		"tsx":                            "null",
		"literate haskell":               "null",
		"forth":                          "#341708",
		"fortran free form":              "null",
		"genshi":                         "null",
		"kaitai struct":                  "#773b37",
		"opal":                           "#f7ede0",
		"piglatin":                       "#fcd7de",
		"dylan":                          "#6c616e",
		"pan":                            "#cc0000",
		"tex":                            "#3D6117",
		"plpgsql":                        "null",
		"xc":                             "#99DA07",
		"apollo guidance computer":       "#0B3D91",
		"c2hs haskell":                   "null",
		"thrift":                         "null",
		"glsl":                           "null",
		"nesc":                           "#94B0C7",
		"al":                             "#3AA2B5",
		"common lisp":                    "#3fb68b",
		"lex":                            "#DBCA00",
		"groovy":                         "#e69f56",
		"nearley":                        "#990000",
		"redcode":                        "null",
		"sage":                           "null",
		"zap":                            "#0d665e",
		"api blueprint":                  "#2ACCA8",
		"ec":                             "#913960",
		"freemarker":                     "#0050b2",
		"literate coffeescript":          "null",
		"less":                           "#1d365d",
		"ags script":                     "#B9D9FF",
		"macaulay2":                      "#d8ffff",
		"shell":                          "#89e051",
		"agda":                           "#315665",
		"autohotkey":                     "#6594b9",
		"dtrace":                         "null",
		"odin":                           "#60AFFE",
		"processing":                     "#0096D8",
		"stylus":                         "#ff6347",
		"apex":                           "#1797c0",
		"factor":                         "#636746",
		"javascript":                     "#f1e05a",
		"labview":                        "null",
		"literate agda":                  "null",
		"lolcode":                        "#cc9900",
		"moocode":                        "null",
		"saltstack":                      "#646464",
		"ti program":                     "#A0AA87",
		"aspectj":                        "#a957b0",
		"c":                              "#555555",
		"c#":                             "#178600",
		"cool":                           "null",
		"jsx":                            "null",
		"clips":                          "null",
		"crystal":                        "#000100",
		"purebasic":                      "#5a6986",
		"runoff":                         "#665a4e",
		"zil":                            "#dc75e5",
		"bluespec":                       "null",
		"eclipse":                        "null",
		"netlogo":                        "#ff6375",
		"g-code":                         "#D08CF2",
		"isabelle root":                  "null",
		"rexx":                           "null",
		"e":                              "#ccce35",
		"python console":                 "null",
		"terra":                          "#00004c",
		"yaml":                           "#cb171e",
		"boo":                            "#d4bec1",
		"hcl":                            "null",
		"jsonnet":                        "#0064bd",
		"logtalk":                        "null",
		"objectscript":                   "#424893",
		"xproc":                          "null",
		"eiffel":                         "#4d6977",
		"objective-j":                    "#ff0c5a",
		"unrealscript":                   "#a54c4d",
		"yasnippet":                      "#32AB90",
		"dhall":                          "#dfafff",
		"j":                              "#9EEDFF",
		"fancy":                          "#7b9db4",
		"holyc":                          "#ffefaf",
		"visual basic .net":              "#945db7",
		"wordpress":                      "#1F6F93",
	}

	height := 700
	width := 600
	titleboxheight := 50
	padding := 10
	strokewidth := 3
	boxwidth := 120
	boxheight := 60
	body := []string{
		`<style>`,
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.text { font: 20px sans-serif; fill: #` + style.Text + `; font-family: ` + style.Textfont + `; text-decoration: underline;}`,
		`.textwhite { font: 20px sans-serif; fill: #ffffff; font-family: ` + style.Textfont + `; text-decoration: underline;}`,
		`.large {
			font: 25px sans-serif; 
			fill: black
		}`,
		`.title { font: 25px sans-serif; fill: #` + style.Title + `}`,
		`.repobox { 
			fill: #` + style.Box + `;
			border: ` + strconv.Itoa(strokewidth) + `px solid #` + style.Border + `;
		}`,
		`.box {
			fill: #` + style.Background + `;
			border: 3px solid #` + style.Border + `;
			stroke: #` + style.Border + `;
			stroke-width: ` + strconv.Itoa(strokewidth) + `px;
		}`,
		`</style>`,
		fmt.Sprintf(`<text x="20" y="35" class="title">%s</text>`, ToTitleCase(title)),
	}
	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}

	// Algoritm for checking if color is too dark
	colorToDark := func(color string) bool {
		var c = strings.Replace(color, "#", "", -1) // strip #
		rgb, err := strconv.ParseInt(c, 16, 32)     // convert rrggbb to decimal
		if err != nil {
			panic(err.Error())
		}
		r := (rgb >> 16) & 0xff // extract red
		g := (rgb >> 8) & 0xff  // extract green
		b := (rgb >> 0) & 0xff  // extract blue

		rFloat := 0.2126
		gFloat := 0.7152
		bFloat := 0.0722
		r2Float := float64(r)
		g2Float := float64(g)
		b2Float := float64(b)
		luma := math.Sqrt(rFloat*(r2Float*r2Float) +
			gFloat*(g2Float*g2Float) +
			bFloat*(b2Float*b2Float))

		return luma < 80
	}

	// Calculate where repositoryboxes should begin
	posY := titleboxheight + padding

	posX := 0

	// imgsize := boxwidth - (padding * 2)
	originalpos := posX
	newwidth := width
	newheight := height

	row := func(content []string, lang string) {

		if languageColor[lang] == "" {
			bodyAdd(fmt.Sprintf(`<g class="repobox" title="%v" transform="translate(%v,%v) rotate(0)">`, lang, posX+padding, posY))
		} else {
			bodyAdd(fmt.Sprintf(`<g fill="%v" title="%v" transform="translate(%v,%v) rotate(0)">`, languageColor[lang], lang, posX+padding, posY))
		}

		for _, v := range content {
			bodyAdd(v)
		}
		bodyAdd(`</g>`)

		newheight = posY + boxheight + padding
		// check if next box will fit into card
		if posX+boxwidth+(boxwidth+padding) >= width {
			posY += boxheight + padding
			newwidth = posX + boxwidth + (padding * 2)
			posX = originalpos - (boxwidth + padding)
		}

	}
	for _, v := range languages {

		// icon := fmt.Sprintf(`icons/%v/%v-original.svg`, v, v)
		// if v == "tailwindcss" {
		// 	// icon = fmt.Sprintf(`https://raw.githubusercontent.com/devicons/devicon/master/icons/%v/%v-plain.svg`, v, v)
		// 	icon = fmt.Sprintf(`icons/%v/%v-plain.svg`, v, v)
		// }
		// img := fmt.Sprintf(`<image x="%v" y="%v" href="%v" height="%v" width="%v"/>`, boxwidth-imgsize-padding, boxheight-imgsize-padding, icon, imgsize, imgsize)

		str := languageColor[v]
		if languageColor[v] != "" && colorToDark(str) {
			row([]string{
				fmt.Sprintf(`<rect x="0" y="0" rx="5" class="" width="%v" height="%v" />`, boxwidth, boxheight),
				fmt.Sprintf(`<text x="%v" y="%v" class="textwhite">%v</text>`, padding, (boxheight / 2), ToTitleCase(v)),
			}, v)
		} else {
			row([]string{
				fmt.Sprintf(`<rect x="0" y="0" rx="5" class="" width="%v" height="%v" />`, boxwidth, boxheight),
				fmt.Sprintf(`<text x="%v" y="%v" class="text">%v</text>`, padding, (boxheight / 2), ToTitleCase(v)),
			}, v)
		}

		posX += boxwidth + padding
	}

	// adjust the svg size to the content
	if newwidth != width {
		width = newwidth
	}
	if newheight != height {
		height = newheight
	}

	// Line on top
	body = append([]string{fmt.Sprintf(`<rect x="0" y="%v" width="%v" height="%v" fill="#%v"/>`, titleboxheight, width, strokewidth, style.Border)}, body...)
	body = append([]string{fmt.Sprintf(`<rect x="%v" y="%v" class="box" width="%v" height="%v" rx="15"  />`, strokewidth/2, strokewidth/2, width, height)}, body...)
	svgTag := fmt.Sprintf(`<svg width="%v" height="%v" fill="none" viewBox="0 0 %v %v" xmlns="http://www.w3.org/2000/svg">`, width+strokewidth, height+strokewidth, width+strokewidth, height+strokewidth)
	body = append([]string{svgTag}, body...)
	bodyAdd(`</svg>`)
	newcard := Skillscard{title, languages, style, body}
	return newcard

}
func ToTitleCase(str string) string {
	return strings.Title(str)
}
