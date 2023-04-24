package model

type Follower struct {
          IdUser      int `json:"idUser"`
          IdFollowers int `json:"idFollower"`
}

type Followers struct {
          Followers []Followers `json:"followers"`
}
