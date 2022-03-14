package entities

type Blacklist struct {
	Id        int    `json:"id"`
	Uniqid    string `json:"uniqid"`
	ShopId    int    `json:"shop_id"`
	Type      string `json:"type"`
	Data      string `json:"data"`
	Note      string `json:"note"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	UpdatedBy int    `json:"updated_by"`
}

type BlacklistResponse struct {
	Status int `json:"status"`
	Data   struct {
		Blacklist struct {
			Id        int    `json:"id"`
			Uniqid    string `json:"uniqid"`
			ShopId    int    `json:"shop_id"`
			Type      string `json:"type"`
			Data      string `json:"data"`
			Note      string `json:"note"`
			CreatedAt int    `json:"created_at"`
			UpdatedAt int    `json:"updated_at"`
			UpdatedBy int    `json:"updated_by"`
		} `json:"blacklist"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type BlacklistListResponse struct {
	Status int `json:"status"`
	Data   struct {
		Blacklists Blacklist `json:"blacklists"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type CreateBlacklistResponse struct {
	Status int `json:"status"`
	Data   struct {
		Uniqid string `json:"uniqid"`
	} `json:"data"`
	Message string      `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}
