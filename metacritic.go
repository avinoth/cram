package main

import(
      "log"
      "strings"

       "github.com/PuerkitoBio/goquery"
      )

const METACRITIC_URL = "http://www.metacritic.com"

func metacritic(final_resp *Final) {
  title := final_resp.Title

  title = strings.Replace(title, " ", "-", -1)
  title = strings.ToLower(title)

  url := METACRITIC_URL + "/movie/" + title

  page, err := goquery.NewDocument(url)

  if err != nil {
    log.Fatal("Something went wrong while parsing Metaacritic: " + err.Error())
  }

  // name := page.Find("#main .movie_content_head .product_title a span").First().Text()

  // name = strings.Replace(name, " ", "-", -1)

  // fmt.Println(name)
  // if name != title {
  //   log.Fatal("Unable to fetch the movie " + title + " from metacritic..")
  // }

  metascore := page.Find(".product_scores .metascore_summary a span").First().Text()
  user_score := page.Find(".product_scores .side_details .score_summary a div").First().Text()

  final_resp.Ratings["metacritic_metascore"] = metascore
  final_resp.Ratings["metacritic_userscore"] = user_score

}
