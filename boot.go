package main

import (
	plum "github.com/scottraio/plum"

	agents "github.com/scottraio/plum-starter-app/agents"
	models "github.com/scottraio/plum-starter-app/models"
	skills "github.com/scottraio/plum-starter-app/skills"
)

func Boot() plum.AppConfig {
	boot := plum.Boot(plum.Initialize{
		// Embedding function
		Embedding: plum.InitEmbeddings("openai"),
		// Which LLMS to use
		LLM: "openai",
	})

	// register the models
	boot.RegisterModel("recipe", models.Recipe())

	// register the skills
	boot.RegisterSkill("epicurous_lookup", skills.Epicurious())

	// register the agents
	boot.RegisterAgent("chef", agents.Chef())

	return boot
}
