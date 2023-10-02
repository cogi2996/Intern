package resources

type Notification struct {
	ID         int64  `json:"id"`
	UUID       string `json:"uuid"`
	Title      string `json:"title"`
	ReviewText string `json:"review_text"`
	Message    string `json:"message"`
	PublicTask *Task  `json:"public_task"`
	Sender     *User  `json:"sender"`
	SendTime   string `json:"send_time"`
}

type ListNotificationResponse struct {
	Notifications []*Notification `json:"notifications"`
}
