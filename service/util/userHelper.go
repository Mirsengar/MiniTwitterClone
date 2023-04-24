package util

import (
          "fmt"
          "sort"
          
          `MiniTwitterClone/model`
)

func AtLeastTwice(users []model.UserWithAge) bool {
          fmt.Println("------------------------------")
          fmt.Printf("Inside AtLEastTwice with users: %v\n", users)
          for i := 0; i < len(users); i++ {
                    for j := i + 1; j < len(users); j++ {
                              if users[i].Age*2 <= users[j].Age {
                                        fmt.Printf("ZOOMER: %v\n", users[i])
                                        fmt.Printf("BOOMER: %v\n", users[j])
                                        return true
                              }
                    }
          }
          return false
}
func AtLEastTwiceAlt(users []model.UserWithAge) bool {
          fmt.Println("------------------------------")
          fmt.Printf("Inside AtLEastTwiceAlt with users: %v\n", users)
          for _, s := range users {
                    for _, t := range users {
                              if s.Age*2 <= t.Age {
                                        fmt.Printf("ZOOMER: %v\n", s)
                                        fmt.Printf("BOOMER: %v\n", t)
                                        return true
                              }
                    }
          }
          return false
}
func ExactlyTwice(users []model.UserWithAge) bool {
          fmt.Println("------------------------------")
          fmt.Printf("Inside ExactlyTwice with users: %v\n", users)
          for i := 0; i < len(users); i++ {
                    for j := 0; j < len(users); j++ {
                              if users[i].Age == users[j].Age*2 {
                                        fmt.Printf("user who is exactly twice younger is: %v\n", users[j])
                                        fmt.Printf("user who is exactly double age is: %v\n", users[i])
                                        return true
                              }
                    }
          }
          return false
}
func ConstrainedExactlyTwice(users []model.UserWithAge) bool {
          fmt.Println("------------------------------")
          fmt.Printf("Inside ConstrainedExactlyTwice with users: %v\n", users)
          for i, s := range users {
                    fmt.Printf("checking user I: %v\n", s)
                    return BinarySearchAge(users[i+1:], s.Age*2)
          }
          return false
}
func BinarySearchAge(users []model.UserWithAge, age int) bool {
          fmt.Println("------------------------------")
          fmt.Printf("Inside BinarySearch with users: %v\n", users)
          fmt.Printf("age is %v\n", age)
          var left int = 0
          var right int = len(users) - 1
          var mid int = 0
          for left <= right {
                    mid = (left + right) / 2
                    if users[mid].Age == age {
                              fmt.Printf("user who is exactly twice older is: %v\n", users[mid])
                              return true
                    } else if users[mid].Age < age {
                              left = mid + 1
                    } else {
                              right = mid - 1
                    }
          }
          return false
}
func GetUsersWithinLimits(users []model.UserWithAge, limit1 int, limit2 int) []model.UserWithAge {
          fmt.Println("------------------------------")
          fmt.Printf("Inside getUsersWithinLimits with users: %v\n", users)
          var usersWithinLimits []model.UserWithAge
          for _, s := range users {
                    if s.Age >= limit1 && s.Age <= limit2 {
                              usersWithinLimits = append(usersWithinLimits, s)
                    }
          }
          return usersWithinLimits
}
func SortUsersByAge(users []model.User) []model.User {
          fmt.Println("------------------------------")
          fmt.Printf("Inside sortUsersByAge with users: %v\n", users)
          sort.Slice(users, func(i, j int) bool {
                    return users[i].Dob > users[j].Dob
          })
          return users
}
