package main

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name   string
	Enable bool
	Votes  []Vote
}

func (s *Player) GetScore() []Vote {
	fromMain := make([]Vote, 0)
	other := Vote{
		VoteBy: -1,
	}
	count := 0
	for _, v := range s.Votes {
		if v.IsMain {
			fromMain = append(fromMain, v)
		} else {
			other.Scores.First += v.Scores.First
			other.Scores.Second += v.Scores.Second
			other.Scores.Third += v.Scores.Third
			other.Scores.Fourth += v.Scores.Fourth
			other.Scores.Fifth += v.Scores.Fifth
			count += 1
		}
	}
	if count > 0 {
		other.Scores.First /= count
		other.Scores.Second /= count
		other.Scores.Third /= count
		other.Scores.Fourth /= count
		other.Scores.Fifth /= count
	}
	fromMain = append(fromMain, other)
	return fromMain
}
