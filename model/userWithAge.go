package model

type UserWithAge struct {
          Id        int    `json:"id"`
          User      string `json:"user"`
          PassHash  string `json:"passHash"`
          Email     string `json:"email"`
          FirstName string `json:"firstName"`
          LastName  string `json:"lastName"`
          Age       int    `json:"age"`
}
