package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

//  id, first_name, last_name, username, photo_url, auth_date and hash
func checkTelegramAuthorization(params map[string][]string) bool {
	token := "536417314:AAHjFTwRZ5puLNSQCAg2QiQA-WX4Lq0Vms4"
	keyHash := sha256.New()
	keyHash.Write([]byte(token))
	secretkey := keyHash.Sum(nil)

	var checkparams []string
	for k, v := range params {
		if k != "hash" {
			checkparams = append(checkparams, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}
	sort.Strings(checkparams)
	checkString := strings.Join(checkparams, "\n")
	hash := hmac.New(sha256.New, secretkey)
	hash.Write([]byte(checkString))
	hashstr := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(hashstr)
	if hashstr == params["hash"][0] {
		return true
	}
	return false
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	})

	r.GET("/auth", func(c *gin.Context) {
		params := c.Request.URL.Query()
		ok := checkTelegramAuthorization(params)
		if ok {
			info, _ := json.MarshalIndent(params, "", "  ")
			c.String(http.StatusOK, "%s", info)
		} else {
			c.String(http.StatusBadRequest, "bad request")
		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
