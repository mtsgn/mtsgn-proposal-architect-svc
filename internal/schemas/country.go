package schemas

type Country struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CountryDialCode struct {
	ID       uint64  `json:"id"`
	DialCode string  `json:"dial_code"`
	Country  Country `json:"country"`
}
