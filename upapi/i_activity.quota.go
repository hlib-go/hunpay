package upapi

/*
5.8.12  优惠券活动剩余名额查询
返回营销系统配置的优惠券活动剩余名额
*/
func ActivityQuota(c *Config, transSeqId, activityNo, activityType, backendToken string) (r *ActivityQuotaResult, err error) {
	bm := NewBodyMap()
	bm.Set("appId", c.AppId)
	bm.Set("transSeqId", transSeqId)     // 请求ID，可以使用uuid
	bm.Set("activityNo", activityNo)     //活动编号
	bm.Set("activityType", activityType) // 活动类型: 1.u点全场立减（同原线上立减）、2.u点全场券、3线下立减、4折扣券、5代金券、6满抵券、7随机立减券、8凭证券、9提货券、10送货券、11精准营销展示券、12单品券、13单品立减
	bm.Set("backendToken", backendToken)

	resp, err := Post(c, "/activity.quota", bm)
	if err != nil {
		return
	}
	if resp.Resp != E00.Code {
		err = ErrNew(resp.Resp, resp.Msg)
		return
	}
	return
}

type ActivityQuotaResult struct {
	TransSeqId           string `json:"transSeqId          "`
	ActivityNo           string `json:"activityNo          "`
	ActivityName         string `json:"activityName        "`
	ActivityType         string `json:"activityType        "`
	BeginTime            string `json:"beginTime           "`
	EndTime              string `json:"endTime             "`
	ActivityStatus       string `json:"activityStatus      "`
	LimitType            string `json:"limitType           "`
	ActivityMark         string `json:"activityMark        "`
	StartMark            string `json:"startMark           "`
	AllAmount            string `json:"allAmount           "`
	AllRemainAmount      string `json:"allRemainAmount     "`
	AllAmountUseupTime   string `json:"allAmountUseupTime  "`
	MonthAmount          string `json:"monthAmount         "`
	MonthRemainAmount    string `json:"monthRemainAmount   "`
	MonthAmountUseupTime string `json:"monthAmountUseupTime"`
	DayAmount            string `json:"dayAmount           "`
	DayRemainAmount      string `json:"dayRemainAmount     "`
	DayAmountUseupTime   string `json:"dayAmountUseupTime  "`
	AllCount             string `json:"allCount            "`
	AllRemainCount       string `json:"allRemainCount      "`
	AllCountUseupTime    string `json:"allCountUseupTime   "`
	MonthCount           string `json:"monthCount          "`
	MonthRemainCount     string `json:"monthRemainCount    "`
	MonthCountUseupTime  string `json:"monthCountUseupTime "`
	DayCount             string `json:"dayCount            "`
	DayRemainCount       string `json:"dayRemainCount      "`
	DayCountUseupTime    string `json:"dayCountUseupTime   "`
	YearAmount           string `json:"yearAmount          "`
	YearRemainAmount     string `json:"yearRemainAmount    "`
	YearAmountUseupTime  string `json:"yearAmountUseupTime "`
	YearCount            string `json:"yearCount           "`
	YearRemainCount      string `json:"yearRemainCount     "`
	YearCountUseupTime   string `json:"yearCountUseupTime  "`
	AwardShowInfoList    []*ActivityQuotaResultAwardShowInfo
}

type ActivityQuotaResultAwardShowInfo struct {
	ActivityNo           string `json:"activityNo          "`
	ActivityName         string `json:"activityName        "`
	ActivityType         string `json:"activityType        "`
	BeginTime            string `json:"beginTime           "`
	EndTime              string `json:"endTime             "`
	ActivityStatus       string `json:"activityStatus      "`
	LimitType            string `json:"limitType           "`
	ActivityMark         string `json:"activityMark        "`
	StartMark            string `json:"startMark           "`
	AllAmount            string `json:"allAmount           "`
	AllRemainAmount      string `json:"allRemainAmount     "`
	AllAmountUseupTime   string `json:"allAmountUseupTime  "`
	MonthAmount          string `json:"monthAmount         "`
	MonthRemainAmount    string `json:"monthRemainAmount   "`
	MonthAmountUseupTime string `json:"monthAmountUseupTime"`
	DayAmount            string `json:"dayAmount           "`
	DayRemainAmount      string `json:"dayRemainAmount     "`
	DayAmountUseupTime   string `json:"dayAmountUseupTime  "`
	AllCount             string `json:"allCount            "`
	AllRemainCount       string `json:"allRemainCount      "`
	AllCountUseupTime    string `json:"allCountUseupTime   "`
	MonthCount           string `json:"monthCount          "`
	MonthRemainCount     string `json:"monthRemainCount    "`
	MonthCountUseupTime  string `json:"monthCountUseupTime "`
	DayCount             string `json:"dayCount            "`
	DayRemainCount       string `json:"dayRemainCount      "`
	DayCountUseupTime    string `json:"dayCountUseupTime   "`
	YearAmount           string `json:"yearAmount          "`
	YearRemainAmount     string `json:"yearRemainAmount    "`
	YearAmountUseupTime  string `json:"yearAmountUseupTime "`
	YearCount            string `json:"yearCount           "`
	YearRemainCount      string `json:"yearRemainCount     "`
	YearCountUseupTime   string `json:"yearCountUseupTime  "`
}
