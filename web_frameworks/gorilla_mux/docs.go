// Package classification My Sample API New
//
// Documentation for Sample API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

// Member defines the structure for an API product
// DMember request model
type DMember struct {
	// the id for the product
	//
	// required: true
	// min: 1
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	TeamId int    `json:"teamID"`
}

// Team defines the structure for an API product
// DTeam request model
type DTeam struct {
	// the id for the product
	//
	// required: false
	// min: 1
	Id int `json:"id"`

	// the name for this team
	//
	// required: true
	// max length: 255
	Name string `json:"name"`

	// the leader for this team
	//
	// required: true
	Leader *DMember `json:"leader"`
}

// A list of teams
// swagger:response teamsResponse
type teamsResponseWrapper struct {
	// All current teams
	// in: body
	Body []DTeam
}

// A list of members
// swagger:response membersResponse
type membersResponseWrapper struct {
	// All current members
	// in: body
	Body []DMember
}
