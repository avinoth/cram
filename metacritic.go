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

  metascore := page.Find(".product_scores .metascore_summary a span").First().Text()
  user_score := page.Find(".product_scores .side_details .score_summary a div").First().Text()

  final_resp.Ratings["critic"]["metacritic_metascore"] = metascore
  final_resp.Ratings["user"]["metacritic_userscore"] = user_score

}
