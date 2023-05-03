package skills

import (
	"fmt"

	skills "github.com/scottraio/plum/skills"
)

type EpicuriousSkill struct {
	skills.Skill
}

func Epicurious() *skills.Skill {
	var skill *EpicuriousSkill
	// create the model
	skill = &EpicuriousSkill{
		// Model is the base model that you want to use
		Skill: skills.Skill{

			HowTo: `
				You are given a JSON response from a REST API. The JSON object will have a Name, 
				Date Sold, and Warranty Expire field. If the Warranty Expire is in the past, the 
				machine is not under Warranty. If the Warranty Expire is null, the machine is still 
				under Warranty.

				Always just return a sentence like this: The "Name" was sold on "Date Sold" and is "Warranty Status".
			`,

			Return: func(query string) string {
				return skill.Search(query)
			},
		},
	}

	return &skill.Skill
}

func (s *EpicuriousSkill) Search(text string) string {
	// return response from json endpoint
	endpoint := fmt.Sprintf("https://proluxe-portal-api-djsktex6ea-uc.a.run.app/serial/%s", text)
	resp, _ := s.GetJSON(endpoint)
	return s.Describe(text, resp)
}
