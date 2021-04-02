package models

type SessionResponse struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestmap"`
}

type HiRezServerStatusResponse struct {
	EntryDatetime string `json:"entry_datetime"`
	Environment   string `json:"environment"`
	LimitedAccess bool   `json:"limited_access"`
	Platform 		  string `json:"platform"`
	RetMsg 	 			string `json:"ret_msg"`
	Status 				string `json:"status"`
	Version 			string `json:"version"`
}

type DataUsedResponse struct {
	ActiveSessions     int    `json:"Active_Sessions"`
	ConcurrentSessions int    `json:"Concurrent_Sessions"`
	RequestDailyLimit  int    `json:"Request_Limit_Daily"`
	SessionCap 				 int    `json:"Session_Cap"`
	SessionTimeLimit   int    `json:"Session_Time_Limit"`
	TotalRequestsToday int    `json:"Total_Requests_Today"`
	TotalSessionsToday int    `json:"Total_Sessions_Today"`
	RetMsg             string `json:"ret_msg"`
}