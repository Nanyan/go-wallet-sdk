package bitcoin

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInscribe(t *testing.T) {
	network := &chaincfg.TestNet3Params

	commitTxPrevOutputList := make([]*PrevOutput, 0)
	commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
		TxId:       "3c6f205ec2995696d5bc852709d234a63aad82131b5b7615504e2e3e9ff88987",
		VOut:       0,
		Amount:     546,
		Address:    "mouQtmBWDS7JnT65Grj2tPzdSmGKJgRMhE",
		PrivateKey: "cPnvkvUYyHcSSS26iD1dkrJdV7k1RoUqJLhn3CYxpo398PdLVE22",
	})
	commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
		TxId:       "453aa6dd39f31f06cd50b72a8683b8c0402ab36f889d96696317503a025a21b5",
		VOut:       0,
		Amount:     546,
		Address:    "2NF33rckfiQTiE5Guk5ufUdwms8PgmtnEdc",
		PrivateKey: "cPnvkvUYyHcSSS26iD1dkrJdV7k1RoUqJLhn3CYxpo398PdLVE22",
	})
	commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
		TxId:       "22c8a4869f2aa9ee5994959c0978106130290cda53f6e933a8dda2dcb82508d4",
		VOut:       0,
		Amount:     546,
		Address:    "tb1qtsq9c4fje6qsmheql8gajwtrrdrs38kdzeersc",
		PrivateKey: "cPnvkvUYyHcSSS26iD1dkrJdV7k1RoUqJLhn3CYxpo398PdLVE22",
	})
	commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
		TxId:       "aa09fa48dda0e2b7de1843c3db8d3f2d7f2cbe0f83331a125b06516a348abd26",
		VOut:       4,
		Amount:     1142196,
		Address:    "tb1pklh8lqax5l7m2ycypptv2emc4gata2dy28svnwcp9u32wlkenvsspcvhsr",
		PrivateKey: "cPnvkvUYyHcSSS26iD1dkrJdV7k1RoUqJLhn3CYxpo398PdLVE22",
	})

	inscriptionDataList := make([]InscriptionData, 0)
	inscriptionDataList = append(inscriptionDataList, InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"xcvb","amt":"100"}`),
		RevealAddr:  "tb1pklh8lqax5l7m2ycypptv2emc4gata2dy28svnwcp9u32wlkenvsspcvhsr",
	})
	inscriptionDataList = append(inscriptionDataList, InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"xcvb","amt":"10"}`),
		RevealAddr:  "mouQtmBWDS7JnT65Grj2tPzdSmGKJgRMhE",
	})
	inscriptionDataList = append(inscriptionDataList, InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"xcvb","amt":"10000"}`),
		RevealAddr:  "tb1qtsq9c4fje6qsmheql8gajwtrrdrs38kdzeersc",
	})
	inscriptionDataList = append(inscriptionDataList, InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"xcvb","amt":"1"}`),
		RevealAddr:  "2NF33rckfiQTiE5Guk5ufUdwms8PgmtnEdc",
	})

	request := &InscriptionRequest{
		CommitTxPrevOutputList: commitTxPrevOutputList,
		CommitFeeRate:          2,
		RevealFeeRate:          2,
		RevealOutValue:         546,
		InscriptionDataList:    inscriptionDataList,
		ChangeAddress:          "tb1pklh8lqax5l7m2ycypptv2emc4gata2dy28svnwcp9u32wlkenvsspcvhsr",
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(requestBytes))

	txs, err := Inscribe(network, request)
	if err != nil {
		t.Fatal(err)
	}
	txsBytes, err := json.MarshalIndent(txs, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(txsBytes))
}

func TestRealInscribe(t *testing.T) {
	network := &chaincfg.TestNet3Params

	commitTxPrevOutputList := make([]*PrevOutput, 0)
	//1. deploy

	// commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
	// 	TxId:       "c7e50c42d8e98780daf1122dba09dd23b4d9c8c73c9bbf3f48a7d6d47945d8c3",
	// 	VOut:       1,
	// 	Amount:     31200,
	// 	Address:    "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
	// 	PrivateKey: "<WIF private key>",
	// })

	// inscriptionDataList := make([]InscriptionData, 0)
	// inscriptionDataList = append(inscriptionDataList, InscriptionData{
	// 	ContentType: "text/plain;charset=utf-8",
	// 	Body:        []byte(`{"p":"brc-20","op":"deploy","tick":"sino","max":"21000000000","lim":"100"}`),
	// 	RevealAddr:  "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
	// })

	// 2. mint
	// commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
	// 	TxId:       "e13113f95d26d98724280388deab7eaad6f3fc970c8995de6eedf3306019d419",
	// 	VOut:       1,
	// 	Amount:     29890,
	// 	Address:    "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
	// 	PrivateKey: "<WIF private key>",
	// })

	// inscriptionDataList := make([]InscriptionData, 0)
	// inscriptionDataList = append(inscriptionDataList, InscriptionData{
	// 	ContentType: "text/plain;charset=utf-8",
	// 	Body:        []byte(`{"p":"brc-20","op":"mint","tick":"sino","amt":"100"}`),
	// 	RevealAddr:  "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
	// })

	// 3. transfer
	commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
		TxId:       "59f6beb45927bb9b2c5426aa4e996491908d2cb5ec99581448d87175482cfa4e",
		VOut:       1,
		Amount:     28590,
		Address:    "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
		PrivateKey: "<WIF private key>",
	})

	inscriptionDataList := make([]InscriptionData, 0)
	inscriptionDataList = append(inscriptionDataList, InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"transfer","tick":"sino","amt":"10"}`),
		RevealAddr:  "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
	})

	request := &InscriptionRequest{
		CommitTxPrevOutputList: commitTxPrevOutputList,
		CommitFeeRate:          2,
		RevealFeeRate:          2,
		RevealOutValue:         546,
		InscriptionDataList:    inscriptionDataList,
		ChangeAddress:          "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
	}

	requestBytes, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(requestBytes))

	txs, err := Inscribe(network, request)
	if err != nil {
		t.Fatal(err)
	}
	txsBytes, err := json.MarshalIndent(txs, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(txsBytes))

	/**
			e13113f95d26d98724280388deab7eaad6f3fc970c8995de6eedf3306019d419
			4b73d9e1cc69323ebead5604566a3fd653581c31cb6ff74f4b1940161d3ba889
				{
			          "commitTx": "0200000001c3d84579d4d6a7483fbf9b3cc7c8d9b423dd09ba2d12f1da8087e9d8420ce5c7010000006a473044022060fc554dfa15ec6c423fc33fb9e72f6a01173a416c4c15a5f5fde1f8c420a4ea02206217ff88a1289fe1773fea8e41663a293a99061a29f7db79ebd99e3ad92363ac012102cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dfdffffff0248030000000000002251201292e9554b7e1621c021c3bc07b1f1599679268515c707ca534e65baa0767b77c2740000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac00000000",
			          "revealTxs": [
			            "0200000000010119d4196030f3ed6ede95890c97fcf3d6aa7eabde8803282487d9265df91331e10000000000fdffffff0122020000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac03400b9b1dfad163b218e0f20081d41d08c56103157d9dda8854a72367eebcc157356b0c7a8087a79b7b23a669b73964addb99a27377dc2bec628f351649ea22f1dc9020cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dac0063036f7264010118746578742f706c61696e3b636861727365743d7574662d38004a7b2270223a226272632d3230222c226f70223a226465706c6f79222c227469636b223a2273696e6f222c226d6178223a223231303030303030303030222c226c696d223a22313030227d6821c0cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8d00000000"
			          ],
			          "commitTxFee": 470,
			          "revealTxFees": [
			            294
			          ],
			          "commitAddrs": [
			            "tb1pz2fwj42t0ctzrsppcw7q0v03txt8jf59zhrs0jjnfejm4grk0dmszc4wxt"
			          ]
			        }
					//mint
					//59f6beb45927bb9b2c5426aa4e996491908d2cb5ec99581448d87175482cfa4e
					1aa28f86184374d293ea6ffd19e3403cd857b1e6fe0f6d31da6e6917ab444810
			{
		          "commitTx": "020000000119d4196030f3ed6ede95890c97fcf3d6aa7eabde8803282487d9265df91331e1010000006a4730440220749b9e2c9574058a9ef6ffc9f8c1933ebefd8de781f8df61d9180e60a53a71b50220665296de7d6deb3d5f3231c44d19b5a28fa750650f12970372a5d8e94258c8dc012102cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dfdffffff023e030000000000002251200263b5ac6c97992d3529b71d571f92f7c8c195d06a8fcf302944523b00160fccae6f0000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac00000000",
		          "revealTxs": [
		            "020000000001014efa2c487571d848145899ecb52c8d909164994eaa26542c9bbb2759b4bef6590000000000fdffffff0122020000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac034047a589f3749a21bff3aa3d6bde3fa55aa7c27b0fff648e02165a221331bacfaf48f7dbf9a00c1b8e34b914ca2c0186ced57f4f7fa00132080cfe51155cec76757a20cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dac0063036f7264010118746578742f706c61696e3b636861727365743d7574662d3800347b2270223a226272632d3230222c226f70223a226d696e74222c227469636b223a2273696e6f222c22616d74223a22313030227d6821c0cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8d00000000"
		          ],
		          "commitTxFee": 470,
		          "revealTxFees": [
		            284
		          ],
		          "commitAddrs": [
		            "tb1pqf3mttrvj7vj6dffkuw4w8uj7lyvr9wsd28u7vpfg3frkqqkplxq3dtynk"
		          ]
		        }

				// transfer
				bb6082135befb4fdc4ee64bee3d6631284c4598b152492c895e50daf6e4f15cc
				3452f20d3e6d19a2d57b4fa751c4ff23c1debe00fa9e1c60b56dc7a72b96d394
				{
	          "commitTx": "02000000014efa2c487571d848145899ecb52c8d909164994eaa26542c9bbb2759b4bef659010000006b4830450221008b06b97f9c1346abc09bccb84e268d82bb7ff815b497a0446ef0cfb33dc59395022035d606333077b3d0c83b58ce18fb53292886903ba6f9d94e51a868a2862f9eed012102cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dfdffffff023e030000000000002251205d2d83eba5b0f425dd51a371e89bdb642c0788109c08f3f02bf448d0223206749a6a0000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac00000000",
	          "revealTxs": [
	            "02000000000101cc154f6eaf0de595c89224158b59c4841263d6e3be64eec4fdb4ef5b138260bb0000000000fdffffff0122020000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac034087b66204b438bd74c7e5a1daaf2e469fbc9619e23a11ddc0a640fa8109b01c986aa758c2839fcabca505ffb1146bae51d4f1132398d299cbbf5e63ae4c4059657d20cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dac0063036f7264010118746578742f706c61696e3b636861727365743d7574662d3800377b2270223a226272632d3230222c226f70223a227472616e73666572222c227469636b223a2273696e6f222c22616d74223a223130227d6821c1cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8d00000000"
	          ],
	          "commitTxFee": 470,
	          "revealTxFees": [
	            284
	          ],
	          "commitAddrs": [
	            "tb1pt5kc86a9kr6zth235dc73x7mvskq0zqsnsy08upt73ydqg3jqe6qnefv82"
	          ]
	        }
			// transfer
			58ce0330df6ca65e6e25a25578308639d53caeffc6b00ea5ec7c51885a585c94
	*/
}

func TestBuildTx(t *testing.T) {
	// legacy address
	txBuild := NewTxBuild(1, &chaincfg.TestNet3Params)
	txBuild.AddInput2("1aa28f86184374d293ea6ffd19e3403cd857b1e6fe0f6d31da6e6917ab444810", 0, "<WIF private key>", "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk", 546)
	txBuild.AddInput2("58ce0330df6ca65e6e25a25578308639d53caeffc6b00ea5ec7c51885a585c94", 1, "<WIF private key>", "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk", 26820)
	txBuild.AddOutput("tb1q7p4ensjfx674q06egyxtc26gcy7237p9xu8vq2", 546)
	txBuild.AddOutput("n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk", 26070)
	tx, err := txBuild.Build()
	assert.Nil(t, err)
	txHex, err := GetTxHex(tx)
	assert.Nil(t, err)
	fmt.Println(txHex)
}

func TestGenkey(t *testing.T) {
	network := &chaincfg.TestNet3Params

	pri, err := btcec.NewPrivateKey()
	require.Nil(t, err)
	wif, err := btcutil.NewWIF(pri, network, true)
	require.Nil(t, err)

	fmt.Println("wif: ", wif.String())
	publicKey := pri.PubKey().SerializeCompressed()
	fmt.Printf("pubkey: %x\n", publicKey)

	p2pkh, err := PubKeyToAddr(publicKey, LEGACY, network)
	assert.Nil(t, err)
	fmt.Printf("p2pkh: %s\n", p2pkh)

	p2wpkh, err := PubKeyToAddr(publicKey, SEGWIT_NATIVE, network)
	assert.Nil(t, err)
	fmt.Printf("p2wpkh: %s\n", p2wpkh)

	p2sh, err := PubKeyToAddr(publicKey, SEGWIT_NESTED, network)
	assert.Nil(t, err)
	fmt.Printf("p2sh: %s\n", p2sh)

	p2tr, err := PubKeyToAddr(publicKey, TAPROOT, network)
	assert.Nil(t, err)
	fmt.Printf("p2tr: %s\n", p2tr)

}
