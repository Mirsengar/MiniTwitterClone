package model

type Tweet struct {
          Id        int `json:"id"`
          UserId    int `json:"userId"`
          Tweet     string
          DataTweet string `json:"dataTweet"`
}
type Tweets struct {
          Tweets []Tweets `json:"tweets"`
}
