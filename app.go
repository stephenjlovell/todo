package main

type Item struct {
	Id        int    `json:"-"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Url       string `json:"url"`
}

var mockApp *App

func init() {
	mockApp = &App{
		items: make([]Item, 10, 10),
	}
}

type App struct {
	items []Item
}

func (app *App) mockItem() *Item {
	return &Item{}
}
