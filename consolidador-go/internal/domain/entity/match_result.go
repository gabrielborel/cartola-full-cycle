package entity

import "strconv"

type MatchResult struct {
	teamAScore int
	teamBScore int
}

func NewMatchResult(teamAScore, teamBScore int) *MatchResult {
	return &MatchResult{
		teamAScore: teamAScore,
		teamBScore: teamBScore,
	}
}

func (m *MatchResult) GetResult() string {
	return strconv.Itoa(m.teamAScore) + "-" + strconv.Itoa(m.teamBScore)
}
