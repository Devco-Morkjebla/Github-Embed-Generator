package project

import (
	"encoding/json"
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ProjectActivity []struct {
	Total int `json:"total"`
	Weeks []struct {
		Week      int `json:"w"`
		Additions int `json:"a"`
		Deletions int `json:"d"`
		Commits   int `json:"c"`
	} `json:"weeks"`
	Author struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
}

func recoverFromError() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func Project(project string, cardstyle style.Styles) string {
	goal := 1000

	apiurl := "https://api.github.com/repos/devco-morkjebla/" + project + "/stats/contributors"

	reqAPI, err := http.NewRequest("GET", apiurl, nil)

	if err != nil {
		panic(err.Error())
	}
	clientAPI := &http.Client{}

	responseAPI, err := clientAPI.Do(reqAPI)
	defer recoverFromError()
	if err != nil {
		panic(err.Error())
	}
	defer responseAPI.Body.Close()

	responseDataAPI, err := ioutil.ReadAll(responseAPI.Body)

	if err != nil {
		panic(err)
	}

	var resObjectAPI ProjectActivity
	json.Unmarshal(responseDataAPI, &resObjectAPI)

	i, err := strconv.ParseInt(strconv.Itoa(resObjectAPI[0].Weeks[len(resObjectAPI[0].Weeks)-1].Week), 10, 64)
	defer recoverFromError()
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)

	additions := resObjectAPI[0].Weeks[len(resObjectAPI[0].Weeks)-1].Additions
	deletions := resObjectAPI[0].Weeks[len(resObjectAPI[0].Weeks)-1].Deletions
	commits := resObjectAPI[0].Weeks[len(resObjectAPI[0].Weeks)-1].Commits
	calculatePercent := func(number, total int) int {
		return int((float64(number) / float64(total)) * float64(100))
	}
	fmt.Println(additions)
	fmt.Println(deletions)
	fmt.Println(commits)
	fmt.Println(tm)

	customstyles := []string{
		`.circle {
		transform: rotate(-90deg);
		}`,
		`.rank-circle-rim {
			stroke: #70a5fd;
			fill: none;
			opacity: 0.2;	
		}`,
	}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
		// style.StarsFilter(),
		style.DropShadowRing1(),
		// style.CubePattern(),
	}
	paddingX := 30
	paddingY := 30

	prog1, style1 := card.CircleProgressbar(calculatePercent(additions, goal), 80, 10, 0, 0, "url(#gradient-fill)", "circle")
	prog2, style2 := card.CircleProgressbar(calculatePercent(deletions, goal), 70, 10, 0, 0, "red", "circle")
	prog3, style3 := card.CircleProgressbar(calculatePercent(commits, goal), 60, 10, 0, 0, "url(#paint0_angular_0_1)", "circle")
	customstyles = append(customstyles, style1)
	customstyles = append(customstyles, style2)
	customstyles = append(customstyles, style3)

	body := []string{
		`<g id="Box">`,
		`    <mask id="path-1-inside-1_36_15" fill="white">`,
		`        <path d="M539.343 287.881C545.213 287.881 549.972 283.123 549.972 277.252V248.844C549.972 245.794 547.498 243.338 544.453 243.171C542.626 243.271 540.786 243.322 538.934 243.322C538.892 243.322 538.85 243.322 538.808 243.322C533.308 243.315 527.799 243.315 522.3 243.322C522.258 243.322 522.216 243.322 522.173 243.322C520.804 243.322 519.441 243.294 518.085 243.239C516.73 243.294 515.367 243.322 513.997 243.322C513.956 243.322 513.914 243.322 513.873 243.322C508.644 243.315 503.408 243.315 498.179 243.322C498.137 243.322 498.096 243.322 498.054 243.322C498.012 243.322 497.97 243.322 497.927 243.322C492.292 243.315 486.647 243.315 481.012 243.322C480.969 243.322 480.927 243.322 480.885 243.322C426.021 243.322 381.546 198.847 381.546 143.983C381.546 89.1204 426.021 44.645 480.885 44.645C480.927 44.645 480.97 44.645 481.012 44.6451C486.647 44.6522 492.292 44.6522 497.927 44.6451C497.969 44.645 498.012 44.645 498.054 44.645C498.096 44.645 498.138 44.645 498.179 44.6451C503.408 44.6515 508.644 44.6515 513.872 44.6451C513.914 44.645 513.956 44.645 513.997 44.645C515.367 44.645 516.73 44.6727 518.085 44.7276C519.441 44.6727 520.804 44.645 522.173 44.645C522.216 44.645 522.258 44.645 522.3 44.6451C527.799 44.6519 533.308 44.6519 538.807 44.6451C538.85 44.645 538.892 44.645 538.934 44.645C540.786 44.645 542.626 44.6956 544.453 44.7957C547.498 44.6289 549.972 42.1729 549.972 39.1234V10.7146C549.972 4.84445 545.213 0.0857539 539.343 0.0857544L11.173 0.0858006C5.30285 0.0858011 0.544158 4.84449 0.544159 10.7146L0.544182 277.252C0.544183 283.123 5.30287 287.881 11.173 287.881L539.343 287.881Z"/>`,
		`    </mask>`,
		fmt.Sprintf(`<path fill="#%v" stroke="black" stroke-width="6" mask="url(#path-1-inside-1_36_15)" d="M539.343 287.881C545.213 287.881 549.972 283.123 549.972 277.252V248.844C549.972 245.794 547.498 243.338 544.453 243.171C542.626 243.271 540.786 243.322 538.934 243.322C538.892 243.322 538.85 243.322 538.808 243.322C533.308 243.315 527.799 243.315 522.3 243.322C522.258 243.322 522.216 243.322 522.173 243.322C520.804 243.322 519.441 243.294 518.085 243.239C516.73 243.294 515.367 243.322 513.997 243.322C513.956 243.322 513.914 243.322 513.873 243.322C508.644 243.315 503.408 243.315 498.179 243.322C498.137 243.322 498.096 243.322 498.054 243.322C498.012 243.322 497.97 243.322 497.927 243.322C492.292 243.315 486.647 243.315 481.012 243.322C480.969 243.322 480.927 243.322 480.885 243.322C426.021 243.322 381.546 198.847 381.546 143.983C381.546 89.1204 426.021 44.645 480.885 44.645C480.927 44.645 480.97 44.645 481.012 44.6451C486.647 44.6522 492.292 44.6522 497.927 44.6451C497.969 44.645 498.012 44.645 498.054 44.645C498.096 44.645 498.138 44.645 498.179 44.6451C503.408 44.6515 508.644 44.6515 513.872 44.6451C513.914 44.645 513.956 44.645 513.997 44.645C515.367 44.645 516.73 44.6727 518.085 44.7276C519.441 44.6727 520.804 44.645 522.173 44.645C522.216 44.645 522.258 44.645 522.3 44.6451C527.799 44.6519 533.308 44.6519 538.807 44.6451C538.85 44.645 538.892 44.645 538.934 44.645C540.786 44.645 542.626 44.6956 544.453 44.7957C547.498 44.6289 549.972 42.1729 549.972 39.1234V10.7146C549.972 4.84445 545.213 0.0857539 539.343 0.0857544L11.173 0.0858006C5.30285 0.0858011 0.544158 4.84449 0.544159 10.7146L0.544182 277.252C0.544183 283.123 5.30287 287.881 11.173 287.881L539.343 287.881Z"/>`, cardstyle.Background),
		`</g>`,
		`<g id="Stat" transform="translate(480,145)">`,
		prog1,
		prog2,
		prog3,
		`</g>`,
		`<g data-testid="card-text">`,
		fmt.Sprintf(`<text x="%v" y="%v" id="Stats" class="title">%v Stats</text>`, paddingX, paddingY, card.ToTitleCase(project)),
		fmt.Sprintf(`<text x="%v" y="130" id="Goal" class="text">Goal: %v</text>`, paddingX, goal),
		fmt.Sprintf(`<text x="%v" y="150" id="Additions" class="text">Additions: %v%v</text>`, paddingX, calculatePercent(additions, goal), "%"),
		fmt.Sprintf(`<text x="%v" y="170" id="Deletions" class="text">Deletions: %v%v</text>`, paddingX, calculatePercent(deletions, goal), "%"),
		fmt.Sprintf(`<text x="%v" y="190" id="Commits" class="text">Commits: %v</text>`, paddingX, commits),
		fmt.Sprintf(`<text x="440" y="130" id="Additions" class="text">Add: %v%v</text>`, calculatePercent(additions, goal), "%"),
		fmt.Sprintf(`<text x="440" y="150" id="Deletions" class="text">Del: %v%v</text>`, calculatePercent(deletions, goal), "%"),
		fmt.Sprintf(`<text x="440" y="170" id="Deletions" class="text">Com: %v%v</text>`, calculatePercent(commits, goal), "%"),
		`</g>`,
		`<line id="Line 1" x1="19" y1="108.5" x2="342" y2="108.5" stroke="url(#paint0_angular_0_1)"/>
		<path id="Progress" d="M1064.75 630.448C1064.75 634.16 1063.47 637.253 1060.91 639.728C1058.39 642.16 1054.53 643.376 1049.33 643.376H1040.75V662H1034.93V617.392H1049.33C1054.36 617.392 1058.18 618.608 1060.78 621.04C1063.43 623.472 1064.75 626.608 1064.75 630.448ZM1049.33 638.576C1052.57 638.576 1054.96 637.872 1056.5 636.464C1058.03 635.056 1058.8 633.051 1058.8 630.448C1058.8 624.944 1055.64 622.192 1049.33 622.192H1040.75V638.576H1049.33ZM1077.81 632.624C1078.84 630.619 1080.29 629.061 1082.17 627.952C1084.09 626.843 1086.41 626.288 1089.14 626.288V632.304H1087.61C1081.08 632.304 1077.81 635.845 1077.81 642.928V662H1071.99V626.928H1077.81V632.624ZM1111.16 662.576C1107.88 662.576 1104.89 661.829 1102.2 660.336C1099.56 658.843 1097.47 656.731 1095.93 654C1094.44 651.227 1093.69 648.027 1093.69 644.4C1093.69 640.816 1094.46 637.659 1095.99 634.928C1097.57 632.155 1099.71 630.043 1102.39 628.592C1105.08 627.099 1108.09 626.352 1111.42 626.352C1114.75 626.352 1117.75 627.099 1120.44 628.592C1123.13 630.043 1125.24 632.133 1126.78 634.864C1128.36 637.595 1129.15 640.773 1129.15 644.4C1129.15 648.027 1128.33 651.227 1126.71 654C1125.13 656.731 1122.98 658.843 1120.25 660.336C1117.52 661.829 1114.49 662.576 1111.16 662.576ZM1111.16 657.456C1113.25 657.456 1115.21 656.965 1117.05 655.984C1118.88 655.003 1120.36 653.531 1121.47 651.568C1122.62 649.605 1123.19 647.216 1123.19 644.4C1123.19 641.584 1122.64 639.195 1121.53 637.232C1120.42 635.269 1118.97 633.819 1117.18 632.88C1115.39 631.899 1113.44 631.408 1111.35 631.408C1109.22 631.408 1107.26 631.899 1105.47 632.88C1103.72 633.819 1102.31 635.269 1101.24 637.232C1100.17 639.195 1099.64 641.584 1099.64 644.4C1099.64 647.259 1100.15 649.669 1101.18 651.632C1102.24 653.595 1103.65 655.067 1105.4 656.048C1107.15 656.987 1109.07 657.456 1111.16 657.456ZM1151.27 626.352C1154.3 626.352 1156.94 627.013 1159.2 628.336C1161.51 629.659 1163.21 631.323 1164.32 633.328V626.928H1170.21V662.768C1170.21 665.968 1169.53 668.805 1168.16 671.28C1166.8 673.797 1164.84 675.76 1162.28 677.168C1159.76 678.576 1156.81 679.28 1153.44 679.28C1148.84 679.28 1145 678.192 1141.92 676.016C1138.85 673.84 1137.04 670.875 1136.48 667.12H1142.24C1142.88 669.253 1144.21 670.96 1146.21 672.24C1148.22 673.563 1150.63 674.224 1153.44 674.224C1156.64 674.224 1159.25 673.221 1161.25 671.216C1163.3 669.211 1164.32 666.395 1164.32 662.768V655.408C1163.17 657.456 1161.46 659.163 1159.2 660.528C1156.94 661.893 1154.3 662.576 1151.27 662.576C1148.15 662.576 1145.32 661.808 1142.76 660.272C1140.24 658.736 1138.25 656.581 1136.8 653.808C1135.35 651.035 1134.63 647.877 1134.63 644.336C1134.63 640.752 1135.35 637.616 1136.8 634.928C1138.25 632.197 1140.24 630.085 1142.76 628.592C1145.32 627.099 1148.15 626.352 1151.27 626.352ZM1164.32 644.4C1164.32 641.755 1163.79 639.451 1162.72 637.488C1161.66 635.525 1160.21 634.032 1158.37 633.008C1156.58 631.941 1154.6 631.408 1152.42 631.408C1150.24 631.408 1148.26 631.92 1146.47 632.944C1144.68 633.968 1143.25 635.461 1142.18 637.424C1141.11 639.387 1140.58 641.691 1140.58 644.336C1140.58 647.024 1141.11 649.371 1142.18 651.376C1143.25 653.339 1144.68 654.853 1146.47 655.92C1148.26 656.944 1150.24 657.456 1152.42 657.456C1154.6 657.456 1156.58 656.944 1158.37 655.92C1160.21 654.853 1161.66 653.339 1162.72 651.376C1163.79 649.371 1164.32 647.045 1164.32 644.4ZM1185.88 632.624C1186.9 630.619 1188.35 629.061 1190.23 627.952C1192.15 626.843 1194.47 626.288 1197.21 626.288V632.304H1195.67C1189.14 632.304 1185.88 635.845 1185.88 642.928V662H1180.05V626.928H1185.88V632.624ZM1235.93 643.12C1235.93 644.229 1235.86 645.403 1235.74 646.64H1207.7C1207.92 650.096 1209.09 652.805 1211.22 654.768C1213.4 656.688 1216.02 657.648 1219.1 657.648C1221.61 657.648 1223.7 657.072 1225.37 655.92C1227.07 654.725 1228.27 653.147 1228.95 651.184H1235.22C1234.29 654.555 1232.41 657.307 1229.59 659.44C1226.78 661.531 1223.28 662.576 1219.1 662.576C1215.77 662.576 1212.78 661.829 1210.14 660.336C1207.53 658.843 1205.49 656.731 1203.99 654C1202.5 651.227 1201.75 648.027 1201.75 644.4C1201.75 640.773 1202.48 637.595 1203.93 634.864C1205.38 632.133 1207.41 630.043 1210.01 628.592C1212.65 627.099 1215.68 626.352 1219.1 626.352C1222.42 626.352 1225.37 627.077 1227.93 628.528C1230.49 629.979 1232.45 631.984 1233.82 634.544C1235.22 637.061 1235.93 639.92 1235.93 643.12ZM1229.91 641.904C1229.91 639.685 1229.42 637.787 1228.44 636.208C1227.46 634.587 1226.11 633.371 1224.41 632.56C1222.74 631.707 1220.89 631.28 1218.84 631.28C1215.9 631.28 1213.38 632.219 1211.29 634.096C1209.24 635.973 1208.07 638.576 1207.77 641.904H1229.91ZM1256.03 662.576C1253.34 662.576 1250.93 662.128 1248.8 661.232C1246.67 660.293 1244.98 659.013 1243.74 657.392C1242.51 655.728 1241.82 653.829 1241.7 651.696H1247.71C1247.88 653.445 1248.69 654.875 1250.14 655.984C1251.64 657.093 1253.58 657.648 1255.97 657.648C1258.19 657.648 1259.94 657.157 1261.22 656.176C1262.5 655.195 1263.14 653.957 1263.14 652.464C1263.14 650.928 1262.45 649.797 1261.09 649.072C1259.72 648.304 1257.61 647.557 1254.75 646.832C1252.15 646.149 1250.02 645.467 1248.35 644.784C1246.73 644.059 1245.32 643.013 1244.13 641.648C1242.98 640.24 1242.4 638.405 1242.4 636.144C1242.4 634.352 1242.93 632.709 1244 631.216C1245.07 629.723 1246.58 628.549 1248.54 627.696C1250.51 626.8 1252.75 626.352 1255.26 626.352C1259.15 626.352 1262.28 627.333 1264.67 629.296C1267.06 631.259 1268.34 633.947 1268.51 637.36H1262.69C1262.56 635.525 1261.81 634.053 1260.45 632.944C1259.12 631.835 1257.33 631.28 1255.07 631.28C1252.98 631.28 1251.32 631.728 1250.08 632.624C1248.84 633.52 1248.22 634.693 1248.22 636.144C1248.22 637.296 1248.59 638.256 1249.31 639.024C1250.08 639.749 1251.02 640.347 1252.13 640.816C1253.28 641.243 1254.86 641.733 1256.86 642.288C1259.38 642.971 1261.43 643.653 1263.01 644.336C1264.59 644.976 1265.93 645.957 1267.04 647.28C1268.19 648.603 1268.79 650.331 1268.83 652.464C1268.83 654.384 1268.3 656.112 1267.23 657.648C1266.16 659.184 1264.65 660.4 1262.69 661.296C1260.77 662.149 1258.55 662.576 1256.03 662.576ZM1289.47 662.576C1286.78 662.576 1284.37 662.128 1282.24 661.232C1280.1 660.293 1278.42 659.013 1277.18 657.392C1275.94 655.728 1275.26 653.829 1275.13 651.696H1281.15C1281.32 653.445 1282.13 654.875 1283.58 655.984C1285.07 657.093 1287.02 657.648 1289.41 657.648C1291.62 657.648 1293.37 657.157 1294.65 656.176C1295.93 655.195 1296.57 653.957 1296.57 652.464C1296.57 650.928 1295.89 649.797 1294.53 649.072C1293.16 648.304 1291.05 647.557 1288.19 646.832C1285.59 646.149 1283.45 645.467 1281.79 644.784C1280.17 644.059 1278.76 643.013 1277.57 641.648C1276.41 640.24 1275.84 638.405 1275.84 636.144C1275.84 634.352 1276.37 632.709 1277.44 631.216C1278.5 629.723 1280.02 628.549 1281.98 627.696C1283.94 626.8 1286.18 626.352 1288.7 626.352C1292.58 626.352 1295.72 627.333 1298.11 629.296C1300.5 631.259 1301.78 633.947 1301.95 637.36H1296.13C1296 635.525 1295.25 634.053 1293.89 632.944C1292.56 631.835 1290.77 631.28 1288.51 631.28C1286.42 631.28 1284.75 631.728 1283.52 632.624C1282.28 633.52 1281.66 634.693 1281.66 636.144C1281.66 637.296 1282.02 638.256 1282.75 639.024C1283.52 639.749 1284.46 640.347 1285.57 640.816C1286.72 641.243 1288.3 641.733 1290.3 642.288C1292.82 642.971 1294.87 643.653 1296.45 644.336C1298.02 644.976 1299.37 645.957 1300.48 647.28C1301.63 648.603 1302.23 650.331 1302.27 652.464C1302.27 654.384 1301.74 656.112 1300.67 657.648C1299.6 659.184 1298.09 660.4 1296.13 661.296C1294.21 662.149 1291.99 662.576 1289.47 662.576Z" fill="black"/>
		<path id="Rectangle 1" d="M237.382 570V703H111.757H25C11.7452 703 1 692.255 1 679V594C1 580.745 11.7452 570 25 570H111.757H237.382ZM239.382 703V570H356.419H466.279V703H356.419H239.382ZM468.279 703V570H578.14H664C677.255 570 688 580.745 688 594V679C688 692.255 677.255 703 664 703H578.14H468.279Z" stroke="black" stroke-width="2"/>
		<path id="69" d="M116.476 612.164C115.948 609.5 114.304 608.168 111.544 608.168C109.408 608.168 107.812 608.996 106.756 610.652C105.7 612.284 105.184 614.984 105.208 618.752C105.76 617.504 106.672 616.532 107.944 615.836C109.24 615.116 110.68 614.756 112.264 614.756C114.736 614.756 116.704 615.524 118.168 617.06C119.656 618.596 120.4 620.72 120.4 623.432C120.4 625.064 120.076 626.528 119.428 627.824C118.804 629.12 117.844 630.152 116.548 630.92C115.276 631.688 113.728 632.072 111.904 632.072C109.432 632.072 107.5 631.52 106.108 630.416C104.716 629.312 103.744 627.788 103.192 625.844C102.64 623.9 102.364 621.5 102.364 618.644C102.364 609.836 105.436 605.432 111.58 605.432C113.932 605.432 115.78 606.068 117.124 607.34C118.468 608.612 119.26 610.22 119.5 612.164H116.476ZM111.58 617.528C110.548 617.528 109.576 617.744 108.664 618.176C107.752 618.584 107.008 619.22 106.432 620.084C105.88 620.924 105.604 621.956 105.604 623.18C105.604 625.004 106.132 626.492 107.188 627.644C108.244 628.772 109.756 629.336 111.724 629.336C113.404 629.336 114.736 628.82 115.72 627.788C116.728 626.732 117.232 625.316 117.232 623.54C117.232 621.668 116.752 620.204 115.792 619.148C114.832 618.068 113.428 617.528 111.58 617.528ZM129.032 625.016C129.272 626.384 129.824 627.44 130.688 628.184C131.576 628.928 132.764 629.3 134.252 629.3C136.244 629.3 137.708 628.52 138.644 626.96C139.604 625.4 140.06 622.772 140.012 619.076C139.508 620.156 138.668 621.008 137.492 621.632C136.316 622.232 135.008 622.532 133.568 622.532C131.96 622.532 130.52 622.208 129.248 621.56C128 620.888 127.016 619.916 126.296 618.644C125.576 617.372 125.216 615.836 125.216 614.036C125.216 611.468 125.96 609.404 127.448 607.844C128.936 606.26 131.048 605.468 133.784 605.468C137.144 605.468 139.496 606.56 140.84 608.744C142.208 610.928 142.892 614.18 142.892 618.5C142.892 621.524 142.616 624.02 142.064 625.988C141.536 627.956 140.624 629.456 139.328 630.488C138.056 631.52 136.304 632.036 134.072 632.036C131.624 632.036 129.716 631.376 128.348 630.056C126.98 628.736 126.2 627.056 126.008 625.016H129.032ZM134.108 619.76C135.716 619.76 137.036 619.268 138.068 618.284C139.1 617.276 139.616 615.92 139.616 614.216C139.616 612.416 139.112 610.964 138.104 609.86C137.096 608.756 135.68 608.204 133.856 608.204C132.176 608.204 130.832 608.732 129.824 609.788C128.84 610.844 128.348 612.236 128.348 613.964C128.348 615.716 128.84 617.12 129.824 618.176C130.808 619.232 132.236 619.76 134.108 619.76Z" fill="black"/>
		<path id="50" d="M338.723 608.816H326.087V617.024C326.639 616.256 327.455 615.632 328.535 615.152C329.615 614.648 330.779 614.396 332.027 614.396C334.019 614.396 335.639 614.816 336.887 615.656C338.135 616.472 339.023 617.54 339.551 618.86C340.103 620.156 340.379 621.536 340.379 623C340.379 624.728 340.055 626.276 339.407 627.644C338.759 629.012 337.763 630.092 336.419 630.884C335.099 631.676 333.455 632.072 331.487 632.072C328.967 632.072 326.927 631.424 325.367 630.128C323.807 628.832 322.859 627.104 322.523 624.944H325.727C326.039 626.312 326.699 627.38 327.707 628.148C328.715 628.916 329.987 629.3 331.523 629.3C333.419 629.3 334.847 628.736 335.807 627.608C336.767 626.456 337.247 624.944 337.247 623.072C337.247 621.2 336.767 619.76 335.807 618.752C334.847 617.72 333.431 617.204 331.559 617.204C330.287 617.204 329.171 617.516 328.211 618.14C327.275 618.74 326.591 619.568 326.159 620.624H323.063V605.936H338.723V608.816ZM344.732 618.716C344.732 614.588 345.404 611.372 346.748 609.068C348.092 606.74 350.444 605.576 353.804 605.576C357.14 605.576 359.48 606.74 360.824 609.068C362.168 611.372 362.84 614.588 362.84 618.716C362.84 622.916 362.168 626.18 360.824 628.508C359.48 630.836 357.14 632 353.804 632C350.444 632 348.092 630.836 346.748 628.508C345.404 626.18 344.732 622.916 344.732 618.716ZM359.6 618.716C359.6 616.628 359.456 614.864 359.168 613.424C358.904 611.96 358.34 610.784 357.476 609.896C356.636 609.008 355.412 608.564 353.804 608.564C352.172 608.564 350.924 609.008 350.06 609.896C349.22 610.784 348.656 611.96 348.368 613.424C348.104 614.864 347.972 616.628 347.972 618.716C347.972 620.876 348.104 622.688 348.368 624.152C348.656 625.616 349.22 626.792 350.06 627.68C350.924 628.568 352.172 629.012 353.804 629.012C355.412 629.012 356.636 628.568 357.476 627.68C358.34 626.792 358.904 625.616 359.168 624.152C359.456 622.688 359.6 620.876 359.6 618.716Z" fill="black"/>
		<path id="1000" d="M530.194 608.96V605.972H536.962V632H533.65V608.96H530.194ZM542.589 618.716C542.589 614.588 543.261 611.372 544.605 609.068C545.949 606.74 548.301 605.576 551.661 605.576C554.997 605.576 557.337 606.74 558.681 609.068C560.025 611.372 560.697 614.588 560.697 618.716C560.697 622.916 560.025 626.18 558.681 628.508C557.337 630.836 554.997 632 551.661 632C548.301 632 545.949 630.836 544.605 628.508C543.261 626.18 542.589 622.916 542.589 618.716ZM557.457 618.716C557.457 616.628 557.313 614.864 557.025 613.424C556.761 611.96 556.197 610.784 555.333 609.896C554.493 609.008 553.269 608.564 551.661 608.564C550.029 608.564 548.781 609.008 547.917 609.896C547.077 610.784 546.513 611.96 546.225 613.424C545.961 614.864 545.829 616.628 545.829 618.716C545.829 620.876 545.961 622.688 546.225 624.152C546.513 625.616 547.077 626.792 547.917 627.68C548.781 628.568 550.029 629.012 551.661 629.012C553.269 629.012 554.493 628.568 555.333 627.68C556.197 626.792 556.761 625.616 557.025 624.152C557.313 622.688 557.457 620.876 557.457 618.716ZM565.195 618.716C565.195 614.588 565.867 611.372 567.211 609.068C568.555 606.74 570.907 605.576 574.267 605.576C577.603 605.576 579.943 606.74 581.287 609.068C582.631 611.372 583.303 614.588 583.303 618.716C583.303 622.916 582.631 626.18 581.287 628.508C579.943 630.836 577.603 632 574.267 632C570.907 632 568.555 630.836 567.211 628.508C565.867 626.18 565.195 622.916 565.195 618.716ZM580.063 618.716C580.063 616.628 579.919 614.864 579.631 613.424C579.367 611.96 578.803 610.784 577.939 609.896C577.099 609.008 575.875 608.564 574.267 608.564C572.635 608.564 571.387 609.008 570.523 609.896C569.683 610.784 569.119 611.96 568.831 613.424C568.567 614.864 568.435 616.628 568.435 618.716C568.435 620.876 568.567 622.688 568.831 624.152C569.119 625.616 569.683 626.792 570.523 627.68C571.387 628.568 572.635 629.012 574.267 629.012C575.875 629.012 577.099 628.568 577.939 627.68C578.803 626.792 579.367 625.616 579.631 624.152C579.919 622.688 580.063 620.876 580.063 618.716ZM587.8 618.716C587.8 614.588 588.472 611.372 589.816 609.068C591.16 606.74 593.512 605.576 596.872 605.576C600.208 605.576 602.548 606.74 603.892 609.068C605.236 611.372 605.908 614.588 605.908 618.716C605.908 622.916 605.236 626.18 603.892 628.508C602.548 630.836 600.208 632 596.872 632C593.512 632 591.16 630.836 589.816 628.508C588.472 626.18 587.8 622.916 587.8 618.716ZM602.668 618.716C602.668 616.628 602.524 614.864 602.236 613.424C601.972 611.96 601.408 610.784 600.544 609.896C599.704 609.008 598.48 608.564 596.872 608.564C595.24 608.564 593.992 609.008 593.128 609.896C592.288 610.784 591.724 611.96 591.436 613.424C591.172 614.864 591.04 616.628 591.04 618.716C591.04 620.876 591.172 622.688 591.436 624.152C591.724 625.616 592.288 626.792 593.128 627.68C593.992 628.568 595.24 629.012 596.872 629.012C598.48 629.012 599.704 628.568 600.544 627.68C601.408 626.792 601.972 625.616 602.236 624.152C602.524 622.688 602.668 620.876 602.668 618.716Z" fill="black"/>`,
	}

	return strings.Join(card.GenerateCard(cardstyle, defs, body, 600, 300, customstyles...), "\n")
}
