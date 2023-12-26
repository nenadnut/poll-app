package main

// var now = time.Now()

// const (
// 	pollID  = 1
// 	pollID2 = 2
// )

// func regularUser() models.User {
// 	return models.User{
// 		ID:        1,
// 		FirstName: "first name",
// 		LastName:  "last name",
// 		Email:     "email",
// 		Password:  "pass",
// 		Role:      models.ADMIN,
// 		CreatedAt: now,
// 		UpdatedAt: now,
// 	}
// }

// func adminUser() models.User {
// 	return models.User{
// 		ID:        1,
// 		FirstName: "first name",
// 		LastName:  "last name",
// 		Email:     "email",
// 		Password:  "pass",
// 		Role:      models.ADMIN,
// 		CreatedAt: now,
// 		UpdatedAt: now,
// 	}
// }

// func options() []models.Option {
// 	return []models.Option{
// 		{
// 			ID:        1,
// 			Text:      "Text 1",
// 			Chosen:    false,
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		{
// 			ID:        2,
// 			Text:      "Text 2",
// 			Chosen:    true,
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		{
// 			ID:        3,
// 			Text:      "Text 3",
// 			Chosen:    false,
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 		{
// 			ID:        4,
// 			Text:      "Text 4",
// 			Chosen:    true,
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 	}
// }

// func questions() []models.SerializableQuestion {
// 	options := options()

// 	return []models.SerializableQuestion{
// 		{
// 			ID:      1,
// 			Text:    "Text 1",
// 			PollID:  pollID,
// 			Options: options,
// 		},
// 		{
// 			ID:      2,
// 			Text:    "Text 2",
// 			PollID:  pollID,
// 			Options: options,
// 		},
// 		{
// 			ID:      3,
// 			Text:    "Text 3",
// 			PollID:  pollID,
// 			Options: options,
// 		},
// 		{
// 			ID:      4,
// 			Text:    "Text 4",
// 			PollID:  pollID,
// 			Options: options,
// 		},
// 	}
// }

// func polls() []models.SerializablePoll {
// 	adminUser := adminUser()
// 	regularUser := regularUser()
// 	questions := questions()

// 	return []models.SerializablePoll{
// 		{
// 			ID:          1,
// 			Title:       "Test poll",
// 			Description: "Test description",
// 			Questions:   questions,
// 			CreatorID:   adminUser.ID,
// 			Completed:   false,
// 		},
// 		{
// 			ID:          2,
// 			Title:       "Test poll 2",
// 			Description: "Test description 2",
// 			Questions:   questions,
// 			CreatorID:   regularUser.ID,
// 			Completed:   true,
// 		},
// 	}
// }
