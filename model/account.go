package model

type FeedEntry struct {
        Id string `json:"id"`
        Name string  `json:"name"`
}

func (a *FeedEntry) ToString() string {
    return a.Id + " " + a.Name
}