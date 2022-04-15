package iamport

type AccessTokenPayload struct {
	ImpKey    string `json:"imp_key"`
	ImpSecret string `json:"imp_secret"`
}

type AccessTokenResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		AccessToken string `json:"access_token"`
		ExpiredAt   int    `json:"expired_at"`
		Now         int    `json:"now"`
	} `json:"response"`
}

type IamportBillingKeyPayload struct {
	CustomerUID string `json:"customer_uid"`
	CardNumber  string `json:"card_number"`
	Expiry      string `json:"expiry"`
	Birth       string `json:"birth"`
	Password    string `json:"pwd_2digit"`
}

type IamportBillingKeyResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		CustomerUID      string `json:"customer_uid"`
		PgProvider       string `json:"pg_provider"`
		PgID             string `json:"pg_id"`
		CardName         string `json:"card_name"`
		CardCode         string `json:"card_code"`
		CardNumber       string `json:"card_number"`
		CardType         string `json:"card_type"`
		CustomerName     string `json:"customer_name"`
		CustomerTel      string `json:"customer_tel"`
		CustomerEmail    string `json:"customer_email"`
		CustomerAddr     string `json:"customer_addr"`
		CustomerPostcode string `json:"customer_postcode"`
		Inserted         int    `json:"inserted"`
		Updated          int    `json:"updated"`
	} `json:"response"`
}

type IamportPaymentPayload struct {
	CustomerUID string `json:"customer_uid"`
	MerchantUID string `json:"merchant_uid"`
	Amount      int    `json:"amount"`
	Name        string `json:"name"`
}

type IamportPaymentResp struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Response struct {
		ImpUID        string `json:"imp_uid"`
		MerchantUID   string `json:"merchant_uid"`
		PayMethod     string `json:"pay_method"`
		Channel       string `json:"channel"`
		PgProvider    string `json:"pg_provider"`
		EmbPgProvider string `json:"emb_pg_provider"`
		PgTID         string `json:"pg_tid"`
		PgID          string `json:"pg_id"`
		Escrow        bool   `json:"escrow"`
		ApplyNum      string `json:"apply_num"`
		BankCode      string `json:"bank_code"`
		BankName      string `json:"bank_name"`
		CardCode      string `json:"card_code"`
		CardName      string `json:"card_name"`
		CardQuota     int    `json:"card_quota"`
		CardNumber    string `json:"card_number"`
		CardType      string `json:"card_type"`
	} `json:"response"`
}
