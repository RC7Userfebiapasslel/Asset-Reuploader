func IsFriend(cookie string, myUserID int64, targetUserID int64) (bool, error) {
    url := fmt.Sprintf(
        "https://friends.roblox.com/v1/users/%d/friends", myUserID,
    )
    // GET request with cookie header
    // Parse the response JSON — it returns a list of friend objects with .id fields
    // Check if targetUserID appears in that list
    // Return true/false
}
