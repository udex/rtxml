package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Comment struct {
	ID    string    `json:"id"`
	Pid   string    `json:"pid"`
	Text  string    `json:"text"`
	User  User      `json:"user"`
	Loc   Locator   `json:"locator"`
	Score int       `json:"score"`
	Votes Votes     `json:"votes"`
	Time  time.Time `json:"time"`
}

type User struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Pic   string `json:"picture"`
	Admin bool   `json:"admin"`
	IP    string `json:"ip"`
}

type Locator struct {
	Site string `json:"site"`
	URL  string `json:"url"`
}

type Votes struct{}

func parseTimestamp(s string) time.Time {
	const format = "2006-01-02 15:04:05"
	t, err := time.Parse(format, s)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func newComment(id, pid, text, timestamp string, user User, loc Locator, score int) Comment {
	return Comment{
		ID:    idbID("", id),
		Pid:   idbID("", pid),
		Text:  text,
		User:  user,
		Loc:   loc,
		Score: score,
		Votes: Votes{},
		Time:  parseTimestamp(timestamp),
	}
}

func newUser(name, uid, ip string) User {
	return User{
		Name:  name,
		ID:    userID(name, uid),
		Pic:   "",
		Admin: false,
		IP:    ip,
	}
}

func newLocator(url string) Locator {
	return Locator{
		Site: "radiot",
		URL:  url,
	}
}

func userID(name, uid string) string {
	prefix := "idb"
	if uid == "" && name == "" {
		return idbID(prefix, "unknown")
	}
	if uid == "" {
		return idbID(prefix, name)
	}
	return idbID(prefix, uid)
}

func idbID(prefix, s string) string {
	if s == "" || s == "0" {
		return ""
	}
	if prefix != "" {
		prefix = fmt.Sprintf("%s_", prefix)
	}
	return fmt.Sprintf("%s%s", prefix, s)
}

// commgen генерирует Comment структуры
func commgen(b Blogposts, r *rt) <-chan Comment {
	ch := make(chan Comment)
	go func() {
		for _, post := range b.Blogposts {
			title := strings.TrimSpace(post.Title)
			// Ищем url поста
			url, err := r.url(title)
			if err != nil {
				// Пост не является ни записью подкаста, ни темами для записи
				log.Println(err)
				continue
			}
			loc := newLocator(url)
			for _, c := range post.Comments.Comments {
				name := c.Name
				uid := c.UserID
				ip := c.IP
				u := newUser(name, uid, ip)
				comment := newComment(c.ID, c.Pid, c.Text, c.Time, u, loc, c.Score)
				ch <- comment
			}
		}
		close(ch)
	}()
	return ch
}

// jsongen генерирует строковые json представления каждого комментария
func jsongen(cs <-chan Comment) <-chan string {
	ch := make(chan string)
	go func() {
		for c := range cs {
			bt, err := json.Marshal(c)
			if err != nil {
				log.Printf("[ERROR] Comment: [%v], error: %s", c, err)
				close(ch)
				break
			}
			ch <- string(bt)
		}
		close(ch)
	}()
	return ch

}
