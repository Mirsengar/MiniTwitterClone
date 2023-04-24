package service

import (
          `github.com/gofiber/fiber/v2`
          
          `MiniTwitterClone/database`
          `MiniTwitterClone/model`
          `MiniTwitterClone/service/util`
)

func GetTweets(c *fiber.Ctx) error {
          res, err := database.DB.Query("SELECT  *FROM tweets")
          if err != nil {
                    return err
          }
          defer res.Close()
          var tweets []model.Tweet
          for res.Next() {
                    tweet := model.Tweet{}
                    err := res.Scan(&tweet.Id, &tweet.UserId, &tweet.Tweet, &tweet.DataTweet)
                    if err != nil {
                              return err
                    }
                    tweets = append(tweets, tweet)
          }
          util.ResponseHelperJSON(c, tweets, "tweets", "No Tweets Found")
          return err
}
