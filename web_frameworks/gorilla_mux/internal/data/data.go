package data

import "fmt"

type Repo struct {
	teams   []Team
	members []Member
}

func NewRepo() *Repo {
	teams := []Team{
		{1, "warwicks", nil},
		{2, "maveriks", nil},
	}
	members := []Member{
		{1, "najam awan", "nsa@gmail.com", 1},
		{2, "najaf awan", "najaf@gmail.com", 2},
	}
	teams[0].Leader = &members[0]
	teams[1].Leader = &members[1]

	return &Repo{teams: teams, members: members}
}

type Member struct {
	Id     int
	Name   string
	Email  string
	teamId int
}

type Team struct {
	Id     int
	Name   string
	Leader *Member
}

func (r *Repo) GetTeams() []Team {
	return r.teams
}

func (r *Repo) GetMembers() []Member {
	return r.members
}

func init() {
	fmt.Println("setup data")
}
