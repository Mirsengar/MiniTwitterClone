package model

type User struct {
          Id        int    `json:"id"`
          User      string `json:"user"`
          PassHash  string `json:"passHash"`
          Email     string `json:"email"`
          FirstName string `json:"firstName"`
          Lastname  string `json:"Lastname"`
          Dob       string `json:"dob"`
}
type Users struct {
          Users []User `json:"users"`
}
