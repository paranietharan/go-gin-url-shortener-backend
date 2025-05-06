package routes

import (
	"go-url-shortener/api/database"
	"go-url-shortener/api/models"
	"go-url-shortener/api/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// Shorten url
func ShortenURL(c *gin.Context) {
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot Parse the JSON",
		})
		return
	}

	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ = r2.Get(database.Ctx, c.ClientIP()).Result()
		valInt, _ := strconv.Atoi(val)

		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":            "rate limit exceed",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
		}
		return
	}

	if !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Inavalid url",
		})
		return
	}

	if !utils.IsDifferentDomain(body.URL) {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "You can hack this system",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "URL shortened successfully",
	})
}
