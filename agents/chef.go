package agents

import (
	plum "github.com/scottraio/plum"
	agents "github.com/scottraio/plum/agents"
)

func Chef() agents.Engine {
	return plum.ChatAgent(`
		As a Plum Agent you are also a Chef. Not only can you cook great food, but it's healthy too!. 
		You're allowed to make substitutions to the recipes, especially for allergies.
		You can learn more about their tastes by asking the user questions. 
		You can play a fun game of "What's for lunch", where you recommend a type of food for lunch.
		Since lunch is the best meal of the day, you can recommend burgers, tacos, or pizza.
	`, ChefTools())
}

func ChefTools() []agents.Tool {
	// Tools
	return []agents.Tool{
		{
			Name:        "Recipes",
			Description: "Useful for finding existing recipes",
			HowTo:       `Example Input: "chicken taco recipe"`,
			Func: func(input agents.Input) string {
				return ""
			},
		},

		{
			Name:        "Epicurous",
			Description: "Useful for the latest recipes and food trends.",
			HowTo:       `Example Input: "top 5 recipes this week"`,
			Func: func(input agents.Input) string {
				// lookup := plum.App.Skills["epicurous"].Return(query)
				// return lookup
				return ""
			},
		},
	}
}
