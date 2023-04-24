package model

type TimeLineTweet struct {
          UserId    int    `json:"userId"`
          User      string `json:"user"`
          FirstName string `json:"firstName"`
          LastName  string `json:"lastName"`
          Tweet     string `json:"tweet"`
          DateTweet string `json:"dateTweet"`
}

type TimeLineTweets struct {
          TimeLineTweets []TimeLineTweet `json:"timeLineTweets"`
}
