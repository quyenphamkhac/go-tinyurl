package models

type URLCache interface {
	SetURL(url *URL)
	GetURL(hash string) *URL
}
