package Gsensor

type AiQiChaSearchResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		QType      int    `json:"qType"`
		QueryStr   string `json:"queryStr"`
		PageNum    int    `json:"pageNum"`
		ResultList []struct {
			Pid     string `json:"pid"`
			EntName string `json:"entName"`
		} `json:"resultList"`
		TotalNumFound string `json:"totalNumFound"`
		TotalPageNum  int    `json:"totalPageNum"`
		VipTotalNum   int    `json:"vipTotalNum"`
		TotalPage     string `json:"totalPage"`
	} `json:"data"`
}

type AiQiChaInfoResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Pid           string `json:"pid"`
		EntName       string `json:"entName"`
		RegAddr       string `json:"regAddr"`
		Addr          string `json:"addr"`
		Scope         string `json:"scope"`
		OpenStatus    string `json:"openStatus"`
		LegalPerson   string `json:"legalPerson"`
		PersonTitle   string `json:"personTitle"`
		StartDate     string `json:"startDate"`
		RegCapital    string `json:"regCapital"`
		PrinType      int    `json:"prinType"`
		LicenseNumber string `json:"licenseNumber"`
		UnifiedCode   string `json:"unifiedCode"`
		TaxNo         string `json:"taxNo"`
		EntLogoWord   string `json:"entLogoWord"`
		PersonLink    string `json:"personLink"`
		PersonId      string `json:"personId"`
		GeoInfo       struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"geoInfo"`
		IsClaim   int    `json:"isClaim"`
		ClaimUrl  string `json:"claimUrl"`
		Email     string `json:"email"`
		Website   string `json:"website"`
		Telephone string `json:"telephone"`
		Phoneinfo []struct {
			Phone          string `json:"phone"`
			Desc           string `json:"desc"`
			Ismobile       int    `json:"ismobile"`
			HideButton     int    `json:"hideButton"`
			PhoneCompCount int    `json:"phoneCompCount"`
		} `json:"phoneinfo"`
		Emailinfo []struct {
			Email string `json:"email"`
			Desc  string `json:"desc"`
		} `json:"emailinfo"`
		Labels struct {
			Opening struct {
				Text      string `json:"text"`
				Style     string `json:"style"`
				FontColor string `json:"fontColor"`
				BgColor   string `json:"bgColor"`
			} `json:"opening"`
			Finance struct {
				Text      string `json:"text"`
				Style     string `json:"style"`
				FontColor string `json:"fontColor"`
				BgColor   string `json:"bgColor"`
			} `json:"finance"`
		} `json:"labels"`
		NewLabels []struct {
			Key   string `json:"key"`
			Value struct {
				Text      string `json:"text"`
				Style     string `json:"style"`
				FontColor string `json:"fontColor"`
				BgColor   string `json:"bgColor"`
			} `json:"value"`
		} `json:"newLabels"`
		Describe    string        `json:"describe"`
		Tags        []interface{} `json:"tags"`
		EntLogo     string        `json:"entLogo"`
		ShareLogo   string        `json:"shareLogo"`
		IsCollected bool          `json:"isCollected"`
		HideButton  int           `json:"hideButton"`
		IsMonitor   int           `json:"isMonitor"`
	} `json:"data"`
}

type AiQiChaICPResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Total     int `json:"total"`
		TotalNum  int `json:"totalNum"`
		PageCount int `json:"pageCount"`
		Page      int `json:"page"`
		Size      int `json:"size"`
		List      []struct {
			Domain   []string `json:"domain"`
			SiteName string   `json:"siteName"`
			HomeSite []string `json:"homeSite"`
			IcpNo    string   `json:"icpNo"`
		} `json:"list"`
	} `json:"data"`
}
type AiQiChaHoldResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Total      int `json:"total"`
		TotalNum   int `json:"totalNum"`
		PageCount  int `json:"pageCount"`
		Page       int `json:"page"`
		Size       int `json:"size"`
		DisplayNum int `json:"displayNum"`
		List       []struct {
			EntName  string `json:"entName"`
			Pid      string `json:"pid"`
			Logo     string `json:"logo"`
			LogoWord string `json:"logoWord"`
			PathData []struct {
				Pathpercent float64 `json:"pathpercent"`
				PathList    []struct {
					InvestComp string  `json:"investComp"`
					Percent    float64 `json:"percent"`
					Pid        string  `json:"pid"`
				} `json:"pathList"`
			} `json:"pathData"`
			Proportion float64 `json:"proportion"`
		} `json:"list"`
	} `json:"data"`
}
