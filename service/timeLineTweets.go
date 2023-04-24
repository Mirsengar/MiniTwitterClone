package service

import (
          "fmt"
          "log"
          
          "github.com/gofiber/fiber/v2"
          
          `MiniTwitterClone/database`
          `MiniTwitterClone/model`
          `MiniTwitterClone/service/util`
)

func GetFeedTweets(c *fiber.Ctx) error {
          dbQuery := "SELECT users.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = ? ORDER BY tweets.date_tweet DESC;"
          rows, err := database.DB.Query(dbQuery, c.Params("id"))
          if err != nil {
                    return util.DefaultErrorHandler(c, err)
          }
          defer rows.Close()
          var timelineTweets []model.TimeLineTweet
          for rows.Next() {
                    timelineTweet := model.TimeLineTweet{}
                    err := rows.Scan(&timelineTweet.UserId, &timelineTweet.User, &timelineTweet.FirstName, &timelineTweet.LastName, &timelineTweet.Tweet, &timelineTweet.DateTweet)
                    if err != nil {
                              log.Fatal(err)
                    }
                    timelineTweets = append(timelineTweets, timelineTweet)
          }
          fmt.Print(timelineTweets)
          util.ResponseHelperJSON(c, timelineTweets, "timeline", "No timeline found")
          return err
}

func GetFeedTweetsPaginated(c *fiber.Ctx) error {
          dbQuery := "SELECT user.user_id, users.user, users.first_name, users.last_name, tweets.tweet, tweets.date_tweet FROM users INNER JOIN tweets ON users.user_id = tweets.user_id INNER JOIN followers ON users.user_id = followers.id_user WHERE followers.id_follower = ? ORDER BY tweets.date_tweet DESC LIMIT ? OFFSET ?;"
          rows, err := database.DB.Query(dbQuery, c.Params("id"), c.Params("limit"), c.Params("offset"))
          if err != nil {
                    return util.DefaultErrorHandler(c, err)
          }
          defer rows.Close()
          var timelineTweets []model.TimeLineTweet
          for rows.Next() {
                    timelineTweet := model.TimeLineTweet{}
                    err := rows.Scan(&timelineTweet.UserId, &timelineTweet.User, &timelineTweet.FirstName, &timelineTweet.LastName, &timelineTweet.Tweet, &timelineTweet.DateTweet)
                    if err != nil {
                              log.Fatal(err)
                    }
                    timelineTweets = append(timelineTweets, timelineTweet)
          }
          util.ResponseHelperJSON(c, timelineTweets, "timeline", "No timeline found")
          return err
}
