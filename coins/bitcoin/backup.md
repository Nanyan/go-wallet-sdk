func TestRealInscribe(t *testing.T) {
	network := &chaincfg.TestNet3Params

	commitTxPrevOutputList := make([]*PrevOutput, 0)
	//1. deploy

	// commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
	// 	TxId:       "1c9c11819115a15c97e2efb3230c7ef728f7e325e47134fca6b38ac564f63c40",
	// 	VOut:       1,
	// 	Amount:     33560,
	// 	Address:    "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
	// 	PrivateKey: "<WIF private key>",
	// })

	// inscriptionDataList := make([]InscriptionData, 0)
	// inscriptionDataList = append(inscriptionDataList, InscriptionData{
	// 	ContentType: "text/plain;charset=utf-8",
	// 	Body:        []byte(`{"p":"brc-20","op":"deploy","tick":"新1","max":"21000000000","lim":"100"}`),
	// 	RevealAddr:  "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
	// })

	// 2. mint
	commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
		TxId:       "e875a2692e57e9568295324e06768a8e11890084f5480f1dd458ffb938034a3d",
		VOut:       0,
		Amount:     546,
		Address:    "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
		PrivateKey: "<WIF private key>",
	})
	commitTxPrevOutputList = append(commitTxPrevOutputList, &PrevOutput{
		TxId:       "0c5fe0cc06a11dbc09b13f693d5c7ad830fa2b271feac4406e13d7417b77a6a0",
		VOut:       1,
		Amount:     32250,
		Address:    "n13PdoYfou4bDxEuRzipRvAB8Tnk91WZJk",
		PrivateKey: "<WIF private key>",
	})

	inscriptionDataList := make([]InscriptionData, 0)
	inscriptionDataList = append(inscriptionDataList, InscriptionData{
		ContentType: "text/plain;charset=utf-8",
		Body:        []byte(`{"p":"brc-20","op":"mint","tick":"新1","amt":"100"}`),
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
					0c5fe0cc06a11dbc09b13f693d5c7ad830fa2b271feac4406e13d7417b77a6a0
					e875a2692e57e9568295324e06768a8e11890084f5480f1dd458ffb938034a3d
					{
				          "commitTx": "0200000001403cf664c58ab3a6fc3471e425e3f728f77e0c23b3efe2975ca1159181119c1c010000006a47304402205c47521123b627e642857eb4c41f7117a8121206db13de397d5afffc1c11e1d302203930eddf723babb5e8efda22d83f859d3c4bfd10ca76d1bb886a7465fee01018012102cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dfdffffff024803000000000000225120420d8e3676d10108fd9de81d99dbcb39ff2d146bd925a1819c3beb2a3d11e084fa7d0000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac00000000",
				          "revealTxs": [
				            "02000000000101a0a6777b41d7136e40c4ea1f272bfa30d87a5c3d693fb109bc1da106cce05f0c0000000000fdffffff0122020000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac0340f4cf9996687260c5f314fb9c165ff2b0f485082e3d6b57e0b6826dc449a5d21c977e76d4bdddc22dd7cacdf30ec47e0f5d495527cb3e649e050be0da838d7f9c9020cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dac0063036f7264010118746578742f706c61696e3b636861727365743d7574662d38004a7b2270223a226272632d3230222c226f70223a226465706c6f79222c227469636b223a22e696b031222c226d6178223a223231303030303030303030222c226c696d223a22313030227d6821c1cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8d00000000"
				          ],
				          "commitTxFee": 470,
				          "revealTxFees": [
				            294
				          ],
				          "commitAddrs": [
				            "tb1pggxcudnk6yqs3lvaaqwenk7t88lj69rtmyj6rqvu804j50g3uzzq5ppujr"
				          ]
				        }
	curl -X 'GET' \
	  'https://open-api-testnet.unisat.io/v1/indexer/brc20/%E6%96%B01/tx/e875a2692e57e9568295324e06768a8e11890084f5480f1dd458ffb938034a3d/history?type=inscribe-deploy&start=0&limit=16' \
	  -H 'accept: application/json' \
	  -H 'Authorization: Bearer 0e13e816bf165727f0751063d967ef018f84e1b2615acf0b14ba24e5de9b2c13'

				// mint
				c7e50c42d8e98780daf1122dba09dd23b4d9c8c73c9bbf3f48a7d6d47945d8c3
				2dbcf8e6a8cffcd81a4044c172dcc25a56bdd9ba29f0801fad3de572500bf9db
				{
		          "commitTx": "02000000023d4a0338b9ff58d41d0f48f5840089118e8a76064e32958256e9572e69a275e8000000006a4730440220465b3ebe0caa1fecf504050c1b9a482576a5222b18d86b0a401e2ac0a5dc0905022076ec499d0aa1344846e2625658a7b77ac36e460476417bba7d2bfca099b4e2e7012102cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dfdffffffa0a6777b41d7136e40c4ea1f272bfa30d87a5c3d693fb109bc1da106cce05f0c010000006b483045022100b768edc1cb0e9ebbf222c07cbf8fea7fc7b744ba8339834ccce9cf82b5d3cf27022002fbcd839141af8a941fe343abff360d03a229f20d12a7b887f9de5231b9e7aa012102cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dfdffffff023e03000000000000225120fc437879eee2bc541aa75b061caecf68db9180f70c758b2fc23611603d4f3485e0790000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac00000000",
		          "revealTxs": [
		            "02000000000101c3d84579d4d6a7483fbf9b3cc7c8d9b423dd09ba2d12f1da8087e9d8420ce5c70000000000fdffffff0122020000000000001976a914d62c059d7907222ca6fc5930dc8135daa2c801f188ac0340497dbba53a33bab5f5649f3043278021a9b19d43aa0bedff0ee40f4278710482e2fa9c4b891c130bf69fa220b870e19bde259fe05ca9651bac5debcc93fd2e837a20cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8dac0063036f7264010118746578742f706c61696e3b636861727365743d7574662d3800347b2270223a226272632d3230222c226f70223a226d696e74222c227469636b223a22e696b031222c22616d74223a22313030227d6821c0cc75aaac9152412d6f61597d921bb14ca25daa7aac4438927d36e3655d7a8d8d00000000"
		          ],
		          "commitTxFee": 766,
		          "revealTxFees": [
		            284
		          ],
		          "commitAddrs": [
		            "tb1pl3phs70wu279gx48tvrpetk0drderq8hp36ckt7zxcgkq020xjzsswfysp"
		          ]
		        }
	*/
}
