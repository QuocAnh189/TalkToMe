package migration

import (
	"fmt"
	"gochat/internal/domain/model"
	"gochat/internal/infrashstructrure/persistence/db"
)

func SeedData(db *db.Database) error {
	// Create 30 users (10x3)
	users := make([]model.User, 0)
	roles := []string{"user", "user", "admin"} // Keep ratio of 2 users : 1 admin

	for i := 0; i < 30; i++ {
		user := model.User{
			ID:        fmt.Sprintf("user-%d", i+1),
			Name:      fmt.Sprintf("User %d", i+1),
			Email:     fmt.Sprintf("user%d@example.com", i+1),
			Password:  "password123",
			AvatarURL: fmt.Sprintf("https://api.dicebear.com/7.x/avataaars/svg?seed=User%d", i+1),
			Role:      roles[i%3],
		}
		// user.Password = utils.HashAndSalt([]byte("password123"))
		users = append(users, user)
	}

	if err := db.GetDB().CreateInBatches(users, len(users)).Error; err != nil {
		return err
	}

	// Create 20 groups (10x2)
	groups := make([]model.Group, 0)
	for i := 0; i < 20; i++ {
		group := model.Group{
			ID:          fmt.Sprintf("group-%d", i+1),
			Name:        fmt.Sprintf("Group %d", i+1),
			AvatarURL:   fmt.Sprintf("https://api.dicebear.com/7.x/avataaars/svg?seed=Group%d", i+1),
			Description: fmt.Sprintf("Description for group %d", i+1),
			OwnerID:     users[i%len(users)].ID,
			Background:  fmt.Sprintf("https://example.com/backgrounds/group%d.jpg", i+1),
		}
		groups = append(groups, group)
	}

	if err := db.GetDB().CreateInBatches(groups, len(groups)).Error; err != nil {
		return err
	}

	// Create 30 group users (10x3)
	groupUsers := make([]model.GroupUser, 0)
	for i := 0; i < 30; i++ {
		groupUser := model.GroupUser{
			ID:       fmt.Sprintf("group-user-%d", i+1),
			UserID:   users[i%len(users)].ID,
			GroupID:  groups[i%len(groups)].ID,
			Nickname: fmt.Sprintf("User%dInGroup%d", i%len(users)+1, i%len(groups)+1),
			IsAdmin:  i%3 == 0, // Every third user is admin
		}
		groupUsers = append(groupUsers, groupUser)
	}

	if err := db.GetDB().CreateInBatches(groupUsers, len(groupUsers)).Error; err != nil {
		return err
	}

	// Create 20 friend relationships (10x2)
	friends := make([]model.Friend, 0)
	for i := 0; i < 20; i++ {
		friend := model.Friend{
			ID:         fmt.Sprintf("friend-%d", i+1),
			InviterID:  users[i%len(users)].ID,
			AccepterID: users[(i+1)%len(users)].ID,
		}
		friends = append(friends, friend)
	}

	if err := db.GetDB().CreateInBatches(friends, len(friends)).Error; err != nil {
		return err
	}

	// Create 20 conversations (10x2)
	conversations := make([]model.Conversation, 0)
	for i := 0; i < 20; i++ {
		conversation := model.Conversation{
			ID:              fmt.Sprintf("conv-%d", i+1),
			UserIDOne:       users[i%len(users)].ID,
			UserIDTwo:       users[(i+1)%len(users)].ID,
			Background:      fmt.Sprintf("https://example.com/backgrounds/conv%d.jpg", i+1),
			UserNicknameOne: fmt.Sprintf("User%dNick", i%len(users)+1),
			UserNicknameTwo: fmt.Sprintf("User%dNick", (i+1)%len(users)+1),
		}
		conversations = append(conversations, conversation)
	}

	if err := db.GetDB().CreateInBatches(conversations, len(conversations)).Error; err != nil {
		return err
	}

	// Create 20 messages (10x2)
	messages := make([]model.Message, 0)
	for i := 0; i < 20; i++ {
		var groupID *string
		var conversationID *string
		if i%2 == 0 {
			gID := groups[i%len(groups)].ID
			groupID = &gID
		} else {
			cID := conversations[i%len(conversations)].ID
			conversationID = &cID
		}

		message := model.Message{
			ID:             fmt.Sprintf("msg-%d", i+1),
			Message:        fmt.Sprintf("Message content %d", i+1),
			GroupID:        groupID,
			SenderID:       users[i%len(users)].ID,
			ConversationID: conversationID,
		}
		messages = append(messages, message)
	}

	if err := db.GetDB().CreateInBatches(messages, len(messages)).Error; err != nil {
		return err
	}

	// Create 20 message attachments (10x2)
	attachments := make([]model.MessageAttachment, 0)
	attachmentTypes := []string{"image", "document", "video", "audio"}
	for i := 0; i < 20; i++ {
		attachment := model.MessageAttachment{
			ID:        fmt.Sprintf("attach-%d", i+1),
			MessageID: messages[i%len(messages)].ID,
			Type:      attachmentTypes[i%len(attachmentTypes)],
			Filename:  fmt.Sprintf("file%d.%s", i+1, attachmentTypes[i%len(attachmentTypes)]),
			URL:       fmt.Sprintf("https://example.com/attachments/file%d.%s", i+1, attachmentTypes[i%len(attachmentTypes)]),
		}
		attachments = append(attachments, attachment)
	}

	if err := db.GetDB().CreateInBatches(attachments, len(attachments)).Error; err != nil {
		return err
	}

	// Create 20 notifications (10x2)
	notifications := make([]model.Notification, 0)
	notificationTypes := []string{"friend_request", "message", "group_invite", "group_message"}
	for i := 0; i < 20; i++ {
		notification := model.Notification{
			ID:       fmt.Sprintf("notif-%d", i+1),
			FromID:   users[i%len(users)].ID,
			ToID:     users[(i+1)%len(users)].ID,
			IsRead:   i%2 == 0,
			Type:     notificationTypes[i%len(notificationTypes)],
			Content:  fmt.Sprintf("Notification content %d", i+1),
			IsAccept: i%2 == 0,
			IsMarked: i%3 == 0,
		}
		notifications = append(notifications, notification)
	}

	if err := db.GetDB().CreateInBatches(notifications, len(notifications)).Error; err != nil {
		return err
	}

	// Update last messages for groups and conversations in batch
	if err := db.GetDB().Model(&model.Group{}).Where("id IN ?", getIDs(groups)).Updates(map[string]interface{}{
		"last_message_id": db.GetDB().Raw("CASE id " + createCaseStatement(groups, messages)),
	}).Error; err != nil {
		return err
	}

	if err := db.GetDB().Model(&model.Conversation{}).Where("id IN ?", getIDs(conversations)).Updates(map[string]interface{}{
		"last_message_id": db.GetDB().Raw("CASE id " + createCaseStatement(conversations, messages)),
	}).Error; err != nil {
		return err
	}

	return nil
}

// Helper functions for batch updates
func getIDs(items interface{}) []string {
	switch v := items.(type) {
	case []model.Group:
		ids := make([]string, len(v))
		for i, item := range v {
			ids[i] = item.ID
		}
		return ids
	case []model.Conversation:
		ids := make([]string, len(v))
		for i, item := range v {
			ids[i] = item.ID
		}
		return ids
	default:
		return nil
	}
}

func createCaseStatement(items interface{}, messages []model.Message) string {
	var result string
	switch v := items.(type) {
	case []model.Group:
		for i, item := range v {
			result += fmt.Sprintf("WHEN '%s' THEN '%s' ", item.ID, messages[i%len(messages)].ID)
		}
	case []model.Conversation:
		for i, item := range v {
			result += fmt.Sprintf("WHEN '%s' THEN '%s' ", item.ID, messages[i%len(messages)].ID)
		}
	}
	result += "END"
	return result
}
