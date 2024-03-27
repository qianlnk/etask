package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

func main() {
	goodID := os.Args[1]

	var param string
	if len(os.Args) >= 3 {
		param = os.Args[2]
	}

	u, err := url.Parse("https://www.iesdouyin.com")
	if err != nil {
		panic(err)
	}

	objectID := fmt.Sprintf("0_%s_%s_0", goodID, goodID)

	sp := SearchParams{
		SearchID:       param,
		SearchResultID: goodID,
		CardStatus:     "",
		AnchorID:       "0",
	}

	spData, err := json.Marshal(sp)
	if err != nil {
		panic(err)
	}

	ei := EntranceInfo{
		EcomSceneID:   "1082",
		CarrierSource: "search_ecommerce",
		SourceMethod:  "product_card",
		EcomGroupType: "video",
		SearchParams:  string(spData),
	}

	eiData, err := json.Marshal(ei)
	if err != nil {
		panic(err)
	}

	gmp := GoodsMetaParams{
		EntranceInfo: string(eiData),
	}

	gmpData, err := json.Marshal(gmp)
	if err != nil {
		panic(err)
	}

	mp := MetaParams{
		SecUid:          "",
		GoodsMetaParams: string(gmpData),
	}
	mpData, err := json.Marshal(mp)
	if err != nil {
		panic(err)
	}

	querys := map[string]string{
		"schema_type":  "20",
		"object_id":    objectID,
		"utm_campaign": "client_scan_share",
		"app":          "aweme",
		"utm_medium":   "android",
		"tt_from":      "scan_share",
		"iid":          "MS4wLjABAAAAGPpcau3yfGfA_DQdXZfuWHWPE8uCohWZWmriCmtBJDMGJYZXNUHnFKbcImjv3fvX",
		"utm_source":   "scan_share",
		"meta_params":  string(mpData),
		"with_sec_did": "1",
		"is_qrcode":    "1",
	}

	q := u.Query()
	for k, v := range querys {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	fmt.Println(u.String())
}

type MetaParams struct {
	SecUid          string `json:"sec_uid"`
	GoodsMetaParams string `json:"goods_meta_params"`
}

type GoodsMetaParams struct {
	EntranceInfo string `json:"entrance_info"`
}

type EntranceInfo struct {
	EcomSceneID   string `json:"ecom_scene_id"`
	CarrierSource string `json:"carrier_source"`
	SourceMethod  string `json:"source_method"`
	EcomGroupType string `json:"ecom_group_type"`
	SearchParams  string `json:"search_params"`
}

type SearchParams struct {
	SearchID       string `json:"search_id"`
	SearchResultID string `json:"search_result_id"`
	CardStatus     string `json:"card_status"`
	AnchorID       string `json:"anchor_id"`
}

// {"sec_uid":"","goods_meta_params":"{\"entrance_info\":\"{\\\"ecom_scene_id\\\":\\\"1082\\\",\\\"carrier_source\\\":\\\"search_ecommerce\\\",\\\"source_method\\\":\\\"product_card\\\",\\\"ecom_group_type\\\":\\\"video\\\",\\\"search_params\\\":\\\"{\\\\\\\"search_id\\\\\\\":\\\\\\\"\\\\\\\",\\\\\\\"search_result_id\\\\\\\":\\\\\\\"3604468181667825642\\\\\\\"}\\\",\\\"card_status\\\":\\\"\\\",\\\"anchor_id\\\":\\\"0\\\"}\"}"}
// https://www.iesdouyin.com/?schema_type=20&object_id=0_3604468181667825642_3604468181667825642_0&utm_campaign=client_scan_share&app=aweme&utm_medium=android&tt_from=scan_share&iid=MS4wLjABAAAAGPpcau3yfGfA_DQdXZfuWHWPE8uCohWZWmriCmtBJDMGJYZXNUHnFKbcImjv3fvX&utm_source=scan_share&meta_params={"sec_uid":"","goods_meta_params":"{\"entrance_info\":
// \"{\\\"ecom_scene_id\\\":\\\"1082\\\",\\\"carrier_source\\\":\\\"search_ecommerce\\\",\\\"source_method\\\":\\\"product_card\\\",\\\"ecom_group_type\\\":\\\"video\\\",\\\"search_params\\\":
// \\\"{\\\\\\\"search_id\\\\\\\":\\\\\\\"\\\\\\\",\\\\\\\"search_result_id\\\\\\\":\\\\\\\"3604468181667825642\\\\\\\"}\\\",\\\"card_status\\\":\\\"\\\",\\\"anchor_id\\\":\\\"0\\\"}\"}"}&with_sec_did=1&is_qrcode=1
