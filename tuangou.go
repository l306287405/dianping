package dianping

type ReceiptPrepareRespBox struct {
	Resp
	Data ReceiptPrepareResp `json:"data"`
}

type ReceiptPrepareResp struct {
	ReceiptValidateResultDTOWithEndDate
	Count          int                                            `json:"count"`                      //可验证的张数
	PaymentDetail  []*PaymentDetailDTO                            `json:"payment_detail"`             //支付明细
	ReceiptInfoMap map[int64]*ReceiptValidateResultDTOWithEndDate `json:"receipt_info_map,omitempty"` //多团单维度卷信息,如果为null则为单团单,key为product_item_id
}

type PaymentDetailDTO struct {
	PaymentDetailId string  `json:"payment_detail_id"` //支付详情id
	Amount          float64 `json:"amount"`            //支付金额
	AmountType      int64   `json:"amount_type"`       //金额类型
}

type ReceiptValidateResultDTOWithEndDate struct {
	ReceiptValidateResultDTO
	ReceiptEndDate string `json:"receipt_end_date"` //卷过期时间
}

type ReceiptValidateResultDTO struct {
	ReceiptCode     string   `json:"receipt_code"`              //验证券码
	DealId          *int64   `json:"deal_id,omitempty"`         //套餐id
	DealgroupId     *int64   `json:"dealgroup_id,omitempty"`    //团购id
	ProductItemId   *int64   `json:"product_item_id,omitempty"` //商品id
	ProductType     *int     `json:"product_type,omitempty"`    //商品类型
	DealTitle       string   `json:"deal_title"`                //商品名称
	DealPrice       *float64 `json:"deal_price,omitempty"`      //商品售卖价格
	DealMarketprice *float64 `json:"deal_marketprice"`          //商品市场价格
	BizType         int      `json:"biz_type"`                  //业务类型
	Mobile          *string  `json:"mobile,omitempty"`          //用户手机号
}

//输码验券校验接口
//http://open.dianping.com/document/v2?docId=6000176&rootDocId=5000
func (s *Service) ReceiptPrepare(r *ReqParams) (resp *ReceiptPrepareRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/receipt/prepare"
	)

	err = r.CheckKeys("session", "receipt_code")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &ReceiptPrepareRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = PostForm(u, r, resp)
	return
}

type ReceiptScanprepareRespBox struct {
	Resp
	Data []*ReceiptPrepareResp `json:"data"`
}

type ReceiptScanprepareResp struct {
	AvailableCount int      `json:"available_count"`           //可用券数量
	ReceiptCode    string   `json:"receipt_code"`              //其中一张券号
	DealId         *int64   `json:"deal_id,omitempty"`         //套餐id
	DealgroupId    *int64   `json:"dealgroup_id,omitempty"`    //团购id
	DealTitle      string   `json:"deal_title"`                //商品名称
	DealPrice      *float64 `json:"deal_price,omitempty"`      //商品售卖价格
	ProductItemId  *int64   `json:"product_item_id,omitempty"` //商品id
	BizType        int      `json:"biz_type"`                  //业务类型
	ProductType    *int     `json:"product_type,omitempty"`    //商品类型
}

//扫码验券校验接口
//http://open.dianping.com/document/v2?docId=6000181&rootDocId=5000
func (s *Service) ReceiptScanprepare(r *ReqParams) (resp *ReceiptScanprepareRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/receipt/scanprepare"
	)

	err = r.CheckKeys("session", "qr_code")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &ReceiptScanprepareRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = PostForm(u, r, resp)
	return
}

type ReceiptConsumeResp struct {
	ReceiptValidateResultDTOWithEndDate
	AppShopId     *string             `json:"app_shop_id,omitempty"`
	OpenShopUuid  *string             `json:"open_shop_uuid,omitempty"`
	PaymentDetail []*PaymentDetailDTO `json:"payment_detail"` //支付明细
}

type ReceiptConsumeRespBox struct {
	Resp
	Data []*ReceiptConsumeResp `json:"data"`
}

//验券接口
//http://open.dianping.com/document/v2?docId=6000177&rootDocId=5000
func (s *Service) ReceiptConsume(r *ReqParams) (resp *ReceiptConsumeRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/receipt/consume"
	)

	err = r.CheckKeys("session", "requestid", "receipt_code", "count", "app_shop_account", "app_shop_accountname")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &ReceiptConsumeRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = PostForm(u, r, resp)
	return
}

type ReceiptBatchconsumeRespBox struct {
	Resp
	Data []*ReceiptBatchconsumeResp `json:"data"`
}

type ReceiptBatchconsumeResp struct {
	SingleConsumeCode int                  `json:"single_consume_code"` //单笔验券code
	SingleConsumeMsg  string               `json:"single_consume_msg"`  //单笔验券错误信息
	ConsumeDetailList []*ConsumeDetailList `json:"consume_detail_list"` //验券详情
}

type ConsumeDetailList struct {
	OrderId      string `json:"order_id"`       //美团点评订单号
	AppShopId    string `json:"app_shop_id"`    //三方shopid
	OpenShopUuid string `json:"open_shop_uuid"` //点评加密门店id
	ReceiptValidateResultDTO
}

//次卡批量验券接口
//http://open.dianping.com/document/v2?docId=6000521&rootDocId=5000
func (s *Service) ReceiptBatchconsume(r *ReqParams) (resp *ReceiptBatchconsumeRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/receipt/batchconsume"
	)

	err = r.CheckKeys("session", "requestid", "app_shop_account", "app_shop_accountname", "receipt_code_infos")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &ReceiptBatchconsumeRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = PostForm(u, r, resp)
	return
}

type ReceiptReverseconsumeRespBox struct {
	Resp
	Data []*ReceiptReverseconsumeResp `json:"data"`
}

type ReceiptReverseconsumeResp struct {
	ReceiptCode     string   `json:"receipt_code"`             //验证券码
	DealId          *int64   `json:"deal_id,omitempty"`        //套餐id
	DealgroupId     *int64   `json:"dealgroup_id,omitempty"`   //团购id
	DealTitle       string   `json:"deal_title"`               //商品名称
	DealPrice       *float64 `json:"deal_price,omitempty"`     //商品售卖价格
	DealMarketprice *float64 `json:"deal_marketprice"`         //商品市场价格
	Mobile          *string  `json:"mobile,omitempty"`         //用户手机号
	AppShopId       *string  `json:"app_shop_id,omitempty"`    //第三方的店铺id，基于商户授权时需要填写
	OpenShopUuid    *string  `json:"open_shop_uuid,omitempty"` //开放平台加密店铺id，基于客户授权时需要填写
}

//撤销验券接口
//http://open.dianping.com/document/v2?docId=6000180&rootDocId=5000
func (s *Service) ReceiptReverseconsume(r *ReqParams) (resp *ReceiptReverseconsumeRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/receipt/reverseconsume"
	)

	err = r.CheckKeys("session", "app_deal_id", "receipt_code", "app_shop_account", "app_shop_accountname")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &ReceiptReverseconsumeRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = PostForm(u, r, resp)
	return
}

type ReceiptGetconsumedRespBox struct {
	Resp
	Data ReceiptGetconsumedResp `json:"data"`
}

type OrderShopPromoDetails struct {
	PromoAmount float64 `json:"promo_amount"`   //优惠金额
	PromoType   int     `json:"promo_type"`     //优惠类型 1-抵用券 2-立减
	Desc        *string `json:"desc,omitempty"` //说明信息
}

type ReceiptGetconsumedResp struct {
	ReceiptValidateResultDTO
	OrderShoppromoAmount  *float64                 `json:"order_shoppromo_amount,omitempty"`   //所在订单商家营销金额
	OrderShopPromoDetails []*OrderShopPromoDetails `json:"order_shop_promo_details,omitempty"` //所在订单商家营销详情
	VerifyAccount         *string                  `json:"verify_account,omitempty"`           //验证账号
	VerifyChannel         string                   `json:"verify_channel"`                     //验证方式 15-第三方验证;其他验证方式等
	VerifyTime            string                   `json:"verify_time"`                        //验券时间
}

//查询已验券信息接口
//http://open.dianping.com/document/v2?docId=6000178&rootDocId=5000
func (s *Service) ReceiptGetconsumed(r *ReqParams) (resp *ReceiptGetconsumedRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/receipt/getconsumed"
	)

	err = r.CheckKeys("session", "receipt_code")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &ReceiptGetconsumedRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = GetRequest(u, r, resp)
	return
}

type ReceiptQuerylistbydateRespBox struct {
	Resp
	Data ReceiptQuerylistbydateResp `json:"data"`
}

type ReceiptQuerylistbydateResp struct {
	TotalCount string                    `json:"total_count"` //总数
	Records    []*ReceiptGetconsumedResp `json:"records"`     //记录列表
}

type ReceiptQuerylistBydateRespRecords struct {
	ReceiptValidateResultDTO
	VerifyAccount *string `json:"verify_account,omitempty"` //验证账号
	VerifyChannel string  `json:"verify_channel"`           //验证方式 15-第三方验证;其他验证方式等
	VerifyTime    string  `json:"verify_time"`              //验券时间
}

//验券记录
//http://open.dianping.com/document/v2?docId=6000179&rootDocId=5000
func (s *Service) ReceiptQuerylistbydate(r *ReqParams) (resp *ReceiptQuerylistbydateRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/receipt/querylistbydate"
	)

	err = r.CheckKeys("session", "date", "offset", "limit")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &ReceiptQuerylistbydateRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = GetRequest(u, r, resp)
	return
}

type DealQueryshopdealRespBox struct {
	Resp
	Data []*DealQueryshopdealResp `json:"data"`
}

type DealQueryshopdealResp struct {
	DealId           int64    `json:"deal_id"`            //套餐id
	DealgroupId      int64    `json:"dealgroup_id"`       //团购id
	BeginDate        string   `json:"begin_date"`         //团购开始售卖时间	格式如 2017-05-24 12:00 ，数据精确到分钟。
	EndDate          string   `json:"end_date"`           //团购结束售卖时间	格式如 2017-05-24 12:00 ，数据精确到分钟。
	Title            string   `json:"title"`              //套餐名称
	Price            float64  `json:"price"`              //套餐价格
	Marketprice      *float64 `json:"marketprice"`        //套餐原价
	ReceiptBeginDate string   `json:"receipt_begin_date"` //团购券开始服务时间	格式如 2017-05-24 12:00 ，数据精确到分钟。
	ReceiptEndDate   string   `json:"reveipt_end_date"`   //团购券结束服务时间	格式如 2017-05-24 12:00 ，数据精确到分钟。
	SaleStatus       int      `json:"sale_status"`        //售卖状态	1-未开始售卖，2-售卖中，3-售卖结束
}

//获取团购信息接口
//http://open.dianping.com/document/v2?docId=6000182&rootDocId=5000
func (s *Service) DealQueryshopdeal(r *ReqParams) (resp *DealQueryshopdealRespBox, err error) {
	var (
		u = OPENAPI_URL + "/tuangou/deal/queryshopdeal"
	)

	err = r.CheckKeys("session")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &DealQueryshopdealRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = GetRequest(u, r, resp)
	return
}

type ProductQueryproductbytypeRespBox struct {
	Resp
	Data []*ProductQueryproductbytypeResp `json:"data"`
}

type ProductQueryproductbytypeResp struct {
	ProductId       int            `json:"product_id"`
	Type            int            `json:"type"`
	ProductItemList []*ProductItem `json:"product_item_list"`
}

type ProductItem struct {
	ProductItemId  int           `json:"product_item_id"`
	ProductId      int           `json:"product_id"`
	Name           string        `json:"name"`
	Price          float64       `json:"price"`
	MarketPrice    float64       `json:"market_price"`
	ServiceSkuList []*ServiceSku `json:"service_sku_list,omitempty"`
}

type ServiceSku struct {
	SkuId    int     `json:"sku_id"`
	SkuName  string  `json:"sku_name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

//按商品类型查询门店商品信息
//http://open.dianping.com/document/v2?docId=6000675&rootDocId=5000
func (s *Service) ProductQueryproductbytype(r *ReqParams) (resp *ProductQueryproductbytypeRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/deal/queryproductbytype"
	)

	err = r.CheckKeys("session", "type")
	if err != nil {
		return
	}
	err = r.ChooseOne("app_shop_id", "open_shop_uuid")
	if err != nil {
		return
	}
	resp = &ProductQueryproductbytypeRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = GetRequest(u, r, resp)
	return
}

type ReceiptQuerybymobileRespBox struct {
	Resp
	Data []*ReceiptQuerybymobileResp `json:"data"`
}

type ReceiptQuerybymobileResp struct {
	SerialNumber string `json:"serial_number"` //团购券码
}

//手机号查询可用团购券
//http://open.dianping.com/document/v2?docId=6000682&rootDocId=5000
func (s *Service) ReceiptQuerybymobile(r *ReqParams) (resp *ReceiptQuerybymobileRespBox, err error) {
	var (
		u = OPENAPI_ROUTER + "/tuangou/receipt/querybymobile"
	)

	err = r.CheckKeys("session", "mobile", "deal_group_id", "deal_id", "offset", "limit", "open_shop_uuid", "platform")
	if err != nil {
		return
	}

	resp = &ReceiptQuerybymobileRespBox{}
	r.AddPublicParams(&s.Config)
	r.Sign(s.Secret)
	err = GetRequest(u, r, resp)
	return
}
