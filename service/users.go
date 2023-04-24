package service

import (
          `github.com/gofiber/fiber/v2`
          
          `MiniTwitterClone/database`
          `MiniTwitterClone/model`
          `MiniTwitterClone/service/util`
)

func GetUser(c *fiber.Ctx) error {
          res, err := database.DB.Query("SELECT *FROM users")
          if err != nil {
                    return util.DefaultErrorHandler(c, err)
          }
          defer res.Close()
          var users []model.User
          for res.Next() {
                    user := model.User{}
                    err := res.Scan(&user.Id, &user.User, &user.PassHash, &user.Email, &user.FirstName, &user.Lastname, &user.Dob)
                    if err != nil {
                              return err
                    }
                    users = append(users, user)
          }
          util.ResponseHelperJSON(c, users, "users", "No Users Found")
          return err
}
func GetUserByAge(c *fiber.Ctx) error {
          res, err := database.DB.Query("SELECT user_id, user, passhash, email, first_name, last_name, TIMESTAMPDIFF(YEAR, dob, CURDATE()) AS age FROM users ORDER BY age ASC")
          if err != nil {
                    return util.DefaultErrorHandler(c, err)
          }
          defer res.Close()
          var users []model.UserWithAge
          for res.Next() {
                    user := model.UserWithAge{}
                    err := res.Scan(&user.Id, &user.User, &user.PassHash, &user.Email, &user.FirstName, &user.LastName, &user.Age)
                    if err != nil {
                              return err
                    }
                    users = append(users, user)
          }
          util.AtLeastTwice(users)
          util.AtLEastTwiceAlt(users)
          util.ExactlyTwice(users)
          util.ResponseHelperJSON(c, users, "user", "Not Found")
          return err
}
