package database

import (
          `database/sql`
)

var DB *sql.DB

func Connect() error {
          var err error
          DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/twitterdb")
          if err != nil {
                    return err
          }
          if err = DB.Ping(); err != nil {
                    return err
          }
          return nil
}
