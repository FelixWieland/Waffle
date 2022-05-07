package database

type FriendEntry struct {
	User1 uint64
	User2 uint64
}

// FriendsGetFriendsList retrieves all friends from the user with ID `userId`
func FriendsGetFriendsList(userId uint64) (result int, friendsList []FriendEntry) {
	var friends = []FriendEntry{}

	queryResult, queryErr := Database.Query("SELECT user_1, user_2 FROM waffle.friends WHERE user_1 = ?", userId)
	defer queryResult.Close()

	if queryErr != nil {
		return -1, friends
	}

	for queryResult.Next() {
		friendEntry := FriendEntry{}

		queryResult.Scan(&friendEntry.User1, &friendEntry.User2)

		friends = append(friends, friendEntry)
	}

	return 0, friends
}

// FriendsAddFriend stores a new friendship in the Database
func FriendsAddFriend(userId uint64, friendId uint64) bool {
	query, queryErr := Database.Query("INSERT INTO waffle.friends (user_1, user_2) VALUES (?, ?)", userId, friendId)
	defer query.Close()

	if queryErr != nil {
		return false
	}

	return true
}

// FriendsRemoveFriend removes a friendship from the Database
func FriendsRemoveFriend(userId uint64, friendId uint64) bool {
	query, queryErr := Database.Query("DELETE FROM waffle.friends WHERE user_1 = ? AND user_2 = ?", userId, friendId)
	defer query.Close()

	if queryErr != nil {
		return false
	}

	return true
}
