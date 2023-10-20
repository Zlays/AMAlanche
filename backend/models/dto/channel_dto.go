package dto

// requests
type ChannelCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

// response
type ChannelCreateResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type ChannelResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	OwnerId uint   `json:"ownerId"`
}

type ChannelGetAllResponse struct {
	Channels []*ChannelResponse `json:"channels"`
}
