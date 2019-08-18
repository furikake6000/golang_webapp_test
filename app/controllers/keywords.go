package controllers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/groovili/gogtrends"
)

func ShowTrend(c *gin.Context) {
	ctx := context.Background()
	explore, err := gogtrends.Explore(ctx,
		&gogtrends.ExploreRequest{
			ComparisonItems: []*gogtrends.ComparisonItem{
				{
					Keyword: "Go",
					Geo:     "US",
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
	fmt.Println(overTime)
}
