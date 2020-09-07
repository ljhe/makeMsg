package tool

type Hero struct {
    HeroId int `json:"heroId"`
    Hp int `json:"hp"`
    Attack int `json:"attack"`
}

// sdns
type EnterGameReq struct {
    UserId1 int `json:"userId1"`  // sd
    OpenId string `json:"openId"`
    Heros []Hero `json:"heros"`
    TimeStamp int64 `json:"timeStamp"`
    Flag []int `json:"flag"`
}

type EnterGameAck struct {
}

