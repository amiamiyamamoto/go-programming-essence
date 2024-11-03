package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindEntries(t *testing.T) {
	//以下がなぜエラーになるかというと、
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.String())
		if r.URL.String() == "/" {
			w.Write([]byte(`
		<table summary="作家データ">
		<tr><td class="header">作家名:</td><td><font size="+2">テスト 太郎</font></td></tr>
		<tr><td class="header">作家名読み:</td><td><font size="+2">テスト 太郎</font></td></tr>
		<tr><td class="header">ローマ字表記:</td><td><font size="+2">Test, Taro</font></td></tr>
		</table>
		<ol>
		<li><a href="../cards/999999/card001.html">テスト書籍001</a></li>
		<li><a href="../cards/999999/card002.html">テスト書籍002</a></li>
		<li><a href="../cards/999999/card003.html">テスト書籍003</a></li>
		</ol>
		`))
		} else {
			pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html`)
			token := pat.FindStringSubmatch(r.URL.String())
			w.Write([]byte(fmt.Sprintf(`
			<table summary="作家データ">
			<tr><td class="header">作家名:</td><td><font size="+2">テスト 太郎</font></td></tr>
			<tr><td class="header">作家名読み:</td><td><font size="+2">テスト 太郎</font></td></tr>
			<tr><td class="header">ローマ字表記:</td><td><font size="+2">Test, Taro</font></td></tr>
			</table>
			<table border="1" summary="ダウンロードデータ" class="download">
			<tr>
			<td>
				<a href="./files/%[1]s_%[2]s.zip">%[1]s_%[2]s.zip</a></td>
			</tr>
			</table>
			`, token[1], token[2])))
		}
	}))
	defer ts.Close()

	tmp := pageURLFormat
	pageURLFormat = ts.URL + "/cards/%s/card%s.html"
	defer func() {
		pageURLFormat = tmp
	}()

	got, err := findEntries(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	want := []Entry{
		{
			AuthorID: "999999",
			Author:   "テスト 太郎",
			TitleID:  "001",
			Title:    "テスト書籍001",
			SiteURL:  ts.URL,
			ZipURL:   ts.URL + "/cards/999999/files/999999_001.zip",
		},
		{
			AuthorID: "999999",
			Author:   "テスト 太郎",
			TitleID:  "002",
			Title:    "テスト書籍002",
			SiteURL:  ts.URL,
			ZipURL:   ts.URL + "/cards/999999/files/999999_002.zip",
		},
		{
			AuthorID: "999999",
			Author:   "テスト 太郎",
			TitleID:  "003",
			Title:    "テスト書籍003",
			SiteURL:  ts.URL,
			ZipURL:   ts.URL + "/cards/999999/files/999999_003.zip",
		},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %+v, but got %+v", want, got)
	}
}

func TestExtractText(t *testing.T) {
	ts := httptest.NewServer(http.FileServer(http.Dir("./testdata")))
	defer ts.Close()
	got, err := extractText(ts.URL + "/example.zip")
	if err != nil {
		t.Error(err)
		return
	}
	want := `愛読書の印象
        芥川龍之介
        
        -------------------------------------------------------
        【テキスト中に現れる記号について】
        
        《》：ルビ
        （例）諳記《あんき》してゐた
        
        ［＃］：入力者注　主に外字の説明や、傍点の位置の指定
        （例）又ため［＃「ため」に傍点］にもなる本である
        -------------------------------------------------------
        
        　子供の時の愛読書は「西遊記」が第一である。これ等は今日でも僕の愛読書である。比喩談としてこれほどの傑作は、西洋には一つもないであらうと思ふ。名高いバンヤンの「天路歴程」なども到底この「西遊記」の敵ではない。それから「水滸伝」も愛読書の一つである。これも今以て愛読してゐる。一時は「水滸伝」の中の一百八人の豪傑の名前を悉く諳記《あんき》してゐたことがある。その時分でも押川春浪氏の冒険小説や何かよりもこの「水滸伝」だの「西遊記」だのといふ方が遥かに僕に面白かつた。
        　中学へ入学前から徳富蘆花氏の「自然と人生」や樗牛の「平家雑感」や小島烏水氏の「日本山水論」を愛読した。同時に、夏目さんの「猫」や鏡花氏の「風流線」や緑雨の「あられ酒」を愛読した。だから人の事は笑へない。僕にも「文章倶楽部」の「青年文士録」の中にあるやうな「トルストイ、坪内士行、大町桂月」時代があつた。
        　中学を卒業してから色んな本を読んだけれども、特に愛読した本といふものはないが、概して云ふと、ワイルドとかゴーチエとかいふやうな絢爛《けんらん》とした小説が好きであつた。それは僕の気質からも来てゐるであらうけれども、一つは慥《たし》かに日本の自然主義的な小説に厭きた反動であらうと思ふ。ところが、高等学校を卒業する前後から、どういふものか趣味や物の見方に大きな曲折が起つて、前に言つたワイルドとかゴーチエとかといふ作家のものがひどくいやになつた。ストリンドベルクなどに傾倒したのはこの頃である。その時分の僕の心持からいふと、ミケエロ・アンヂエロ風な力を持つてゐない芸術はすべて瓦礫のやうに感じられた。これは当時読んだ「ジヤンクリストフ」などの影響であつたらうと思ふ。
        　さういふ心持が大学を卒業する後までも続いたが、段々燃えるやうな力の崇拝もうすらいで、一年前から静かな力のある書物に最も心を惹かれるやうになつてゐる。但、静かなと言つてもたゞ静かだけでも力のないものには余り興味がない。スタンダールやメリメエや日本物で西鶴などの小説はこの点で今の僕には面白くもあり、又ため［＃「ため」に傍点］にもなる本である。
        　序ながら附け加へておくが、此間「ジヤンクリストフ」を出して読んで見たが、昔ほど感興が乗らなかつた。あの時分の本はだめなのかと思つたが、「アンナカレニナ」を出して二三章読んで見たら、これは昔のやうに有難い気がした。
        
        
        
        底本：「芥川龍之介全集　第六巻」岩波書店
        　　　1996（平成8）年4月8日発行
        初出：「文章倶楽部　第5年第8号」
        　　　1920（大正9）年8月1日発行
        ※初出誌に、顔写真と「曇天の水動かずよ芹の中」の句の筆跡写真と共に掲載された。
        入力：砂場清隆
        校正：高柳典子
        2006年2月21日作成
        2006年4月5日修正
        青空文庫作成ファイル：
        このファイルは、インターネットの図書館、青空文庫（http://www.aozora.gr.jp/）で作られました。入力、校正、制作にあたったのは、ボランティアの皆さんです。`
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}
