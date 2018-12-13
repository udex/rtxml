package main

import (
	"crypto/sha256"
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
		ID:    hash("idb", id),
		Pid:   hash("idb", pid),
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
		ID:    hash("disqus", uid),
		Pic:   "",
		Admin: false,
		IP:    hash("", ip),
	}
}

func newLocator(url string) Locator {
	return Locator{
		Site: "radiot",
		URL:  url,
	}
}

func hash(prefix, s string) string {
	if s == "" {
		return s
	}
	h := sha256.New224()
	h.Write([]byte(s))
	return fmt.Sprintf("%s_%x", prefix, h.Sum(nil))
}

func encodeJSON(b Blogposts, r *rt) ([]byte, error) {
	result := []Comment{}
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
			result = append(result, comment)
		}
	}
	bt, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return bt, nil
}
