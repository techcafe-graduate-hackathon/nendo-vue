package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// 日付フォーマット
var layout = "2006-01-02 15:04:05"

// Qiitaのブログ一覧JSON形式
type QiitaJSON struct {
	RenderedBody   string      `json:"rendered_body"`
	Body           string      `json:"body"`
	Coediting      bool        `json:"coediting"`
	CommentsCount  int         `json:"comments_count"`
	CreatedAt      time.Time   `json:"created_at"`
	Group          interface{} `json:"group"`
	ID             string      `json:"id"`
	LikesCount     int         `json:"likes_count"`
	Private        bool        `json:"private"`
	ReactionsCount int         `json:"reactions_count"`
	Tags           []struct {
		Name     string        `json:"name"`
		Versions []interface{} `json:"versions"`
	} `json:"tags"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	User      struct {
		Description       string `json:"description"`
		FacebookID        string `json:"facebook_id"`
		FolloweesCount    int    `json:"followees_count"`
		FollowersCount    int    `json:"followers_count"`
		GithubLoginName   string `json:"github_login_name"`
		ID                string `json:"id"`
		ItemsCount        int    `json:"items_count"`
		LinkedinID        string `json:"linkedin_id"`
		Location          string `json:"location"`
		Name              string `json:"name"`
		Organization      string `json:"organization"`
		PermanentID       int    `json:"permanent_id"`
		ProfileImageURL   string `json:"profile_image_url"`
		TeamOnly          bool   `json:"team_only"`
		TwitterScreenName string `json:"twitter_screen_name"`
		WebsiteURL        string `json:"website_url"`
	} `json:"user"`
	PageViewsCount interface{} `json:"page_views_count"`
}

// メイン関数
func main() {
	getQiita("toRisouP")
}

// Qiitaの取得
func getQiita (userID string) {
	// APIのURLを発行
	url := "http://qiita.com/api/v2/users/" + userID + "/items"

	fmt.Println("get : " + url)

	// GET
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	// レスポンスからBodyのみを取得
	byteArray, _ := ioutil.ReadAll(resp.Body)

	// 取得したJSONをパース
	var blogData []QiitaJSON
	err := json.Unmarshal(byteArray, &blogData)

	// JSONパース時のエラー処理
	if err != nil {
		fmt.Println(err)
		return
	}

	// 今日の日付を取得
	today := timeToDay(time.Now())

	// 一覧を表示
	for i := 0; i < len(blogData); i++ {
		createDay := timeToDay(blogData[i].CreatedAt)

		if stringToTime(createDay).Unix() == stringToTime(today).Unix() - 86400 {
			fmt.Println(blogData[i].ID)
			fmt.Println("\tタイトル : " + blogData[i].Title)
			fmt.Println("\t投稿日　 ：" + blogData[i].CreatedAt.Format(layout))
		} else {
			fmt.Println((stringToTime(today).Unix() - stringToTime(createDay).Unix()) / 86400)
		}
	}

	// TODO: WordPressへPOSTする処理
}

// 日付だけ取得
func timeToDay(t time.Time) string {
	str := t.Format("2006-01-02 00:00:00")
	return str
}

func stringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}