package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/amay0048/mtparser/mtparser"
	_ "github.com/amay0048/mtparser/mtparser"
	_ "github.com/amay0048/mtparser/mtserializer"
)

const file = `{1:F01EBILAEADDMB00000000000}{2:I940EBILAEADXXXXN}{3:{108:EN33162102}}{4:
:20:001
:25:1234567891234
:28C:00200/001
:60F:C210428USD200000,
:61:2104280428C101,S103003071200EN11891//003071200EN11891
RANDOM REFERENCE TEXT
:86:VIRTUAL ACCOUNT SUSPENSE
89898989 A. N. Sender
MR6238900446081060061302008    OTHER TEXT
Virtual A. Name
MORE TEXT
:61:2104280428C301,06S1030040712S58968018//0040712S58968018
RANDOM REFERENCE TEXT
:86:VIRTUAL ACCOUNT SUSPENSE
7979797979 A. N. Sender
MR6238900446081060061302008    OTHER TEXT
Virtual A. Name
MORE TEXT
:62F:C210428USD200050,06
-}
`

const filetest = `{1:F01AAAAGRA0AXXX0057000289}{2:O1030919010321BBBBGRA0AXXX00570001710103210920N}{4:
:20:5387354
:23B:CRED
:23E:PHOB/20.527.19.60
:32A:000526USD1101,50
:33B:USD1121,50
:50K:FRANZ HOLZAPFEL GMBH
VIENNA
:52A:BKAUATWW
:59:723491524
C. KLEIN
BLOEMENGRACHT 15
AMSTERDAM
:71A:SHA
:71F:USD10,
:71F:USD10,
:72:/INS/CHASUS33
-}{5:{MAC:75D138E4}{CHK:DE1B0D71FA96}}`

const data = `{1:F01MIDLGB22AXXX0548034693}{2:I103BKTRUS33XBRDN3}{3:{108:MT103}}{4:
:20:8861198-0706
:23B:CRED
:32A:000612USD5443,99
:33B:USD5443,99
:50K:GIAN ANGELO IMPORTS
NAPLES
:52A:BCITITMM500
:53A:BCITUS33
:54A:IRVTUS3N
:57A:BNPAFRPPGRE
:59:/20041010050500001M02606
KILLY S.A.
GRENOBLE
:70:/RFB/INVOICE 559661
:71A:SHA
-}`

func main() {
	JsonbResearch()
	//
	//mt := &swift.MT940{}
	//err := mt.Unmarshal([]byte(file))
	//lex := gsm.NewLexer(bytes.NewReader([]byte(file)))
	//pars := gsm.NewParser(lex)
	//msg, err := pars.Parse()
	//buf, err := json.Marshal(msg)
	//fmt.Println(string(buf))
	//token, lin := pars.Scan()
	//fmt.Printf("%s %v", lin, token)
	//for _, block := range msg.Blocks {
	//	if block.ID == "4" {
	//		values, ok := block.Value.([]gsm.SwiftBlock)
	//		if !ok {
	//			fmt.Println("error")
	//		}
	//		for _, v := range values {
	//			if v.ID == "86" {
	//				fmt.Printf("%v", v.Value)
	//			}
	//		}
	//
	//	}
	//}
	//
	////parser := readerTest(strings.ReplaceAll(file, "\n", ""))
	//parser := readerTest(data)
	//err = parser.ParseBody()
	//res, err := swift.NewMessageExtractor([]byte(data)).Extract()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Printf("%v", res)
}

func readerTest(f string) mtparser.Parser {
	r := bufio.NewReader(bytes.NewReader([]byte(f)))
	psr := mtparser.New(r)
	if err := psr.Parse(); err != nil {
		fmt.Println(err)
	}

	return psr
}
