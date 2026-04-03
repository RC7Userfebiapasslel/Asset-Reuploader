package friends

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/kartFr/Asset-Reuploader/internal/roblox"
)

type friendEntry struct {
    ID int64 `json:"id"`
}

type friendsResponse struct {
    Data []friendEntry `json:"data"`
}

// IsFriend checks if targetUserID is in myUserID's friends list.
func IsFriend(client *roblox.Client, myUserID int64, targetUserID int64) (bool, error) {
    url := fmt.Sprintf("https://friends.roblox.com/v1/users/%d/friends", myUserID)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return false, err
    }
    req.Header.Set("Cookie", ".ROBLOSECURITY="+client.Cookie)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    var result friendsResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return false, err
    }

    for _, f := range result.Data {
        if f.ID == targetUserID {
            return true, nil
        }
    }
    return false, nil
}
