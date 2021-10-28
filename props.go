package main

import "time"

type Basic struct {
	Title    string
	Type     string
	ImageUrl string
	Url      string
}

type Optional struct {
	Audio, Description, Determiner,
	Locale, SiteName, Video string
	Locales []string
}

type Image struct {
	Url, Secure_url, Type, Width, Height, Alt string
}

type Audio struct {
	Url, Secure_url, Type string
}

type Music struct {
	Song          song
	Album         album
	Playlist      playlist
	Radio_Station radio_Station
}

type song struct {
	Duration, Disc, Track int
	Album                 []album
	Musician              []Profile
}

type album struct {
	Song         song
	Disc, track  int
	Musician     Profile
	release_date time.Time
}

type playlist struct {
	Song        song
	creator     Profile
	Disc, Track int
}

type radio_Station struct {
	creator string
}

type Profile struct {
	firstName, LastName, UserName, Gender string
}

type movie struct {
	Actor, Director, Writer []Profile
	ActorRole               string
	Duration                int
	Release_Date            time.Time
	Tag                     []string
}

type episode struct {
	ActorRole, Actor, Director,
	Writer, Duration, Release_Date,
	Tag, Series string
}

type Article struct {
	Published_time, Modified_time, Expiration_time time.Time
	Author                                         []Profile
	Section                                        string
	Tag                                            []string
}

type Book struct {
	Author       []Profile
	ISBN         string
	Release_Date time.Time
	Tag          []string
}

type Website struct {
	Url string
}

func init() {

}
