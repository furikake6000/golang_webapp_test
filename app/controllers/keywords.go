package controllers

import (
	"context"
	"log"
	"my/models"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/groovili/gogtrends"
)

func ShowKeyword(c *gin.Context) {
	user := models.CurrentUser(c)

	ctx := context.Background()
	explore, err := gogtrends.Explore(ctx,
		&gogtrends.ExploreRequest{
			ComparisonItems: []*gogtrends.ComparisonItem{
				{
					Keyword: c.Param("keyword"),
					Geo:     "JP",
					Time:    "today+12-m",
				},
			},
			Category: 31, // Programming category
			Property: "",
		}, "EN")
	if err != nil {
		panic(err)
	}
	overTime, err := gogtrends.InterestOverTime(ctx, explore[0], "EN")
	if err != nil {
		panic(err)
	}

	c.HTML(200, "keywords.html", gin.H{
		"keyword":  c.Param("keyword"),
		"user":     user,
		"timeline": overTime,
	})
}
func printItems(items interface{}) {
	ref := reflect.ValueOf(items)

	if ref.Kind() != reflect.Slice {
		log.Fatalf("Failed to print %s. It's not a slice type.", ref.Kind())
	}

	for i := 0; i < ref.Len(); i++ {
		log.Println(ref.Index(i).Interface())
	}
}
