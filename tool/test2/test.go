package tool

type Hero struct {
    HeroId int `json:"heroId"`
    Hp int `json:"hp"`
    Attack int `json:"attack"`
}

// sdsds
type EnterGameReq struct {
    UserId2 int `json:"userId2"`
    OpenId string `json:"openId"`  // 123
    Heros []Hero `json:"heros"`
    TimeStamp int64 `json:"timeStamp"`
    Flag []int `json:"flag"`
}

type EnterGameAck struct {
}

