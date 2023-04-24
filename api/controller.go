package api

import (
          `github.com/gofiber/fiber/v2`
          
          `MiniTwitterClone/service`
)

func SetupRoutes(app *fiber.App) {
          api := app.Group("/api")
          api.Get("/users", service.GetUser)
          api.Get("/users/age", service.GetUserByAge)
          api.Get("/tweets", service.GetTweets)
          api.Get("/feed/:id", service.GetFeedTweets)
          api.Get("/feed/:id/:limit/:offset", service.GetFeedTweetsPaginated)
}
