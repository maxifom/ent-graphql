// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewNote struct {
	Body string `json:"body"`
}

type Note struct {
	ID         string `json:"id"`
	Body       string `json:"body"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
