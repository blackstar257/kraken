package kraken

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetStream(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		statusCode int
		slug       string
		respBody   string
	}{
		{
			http.StatusNotFound,
			"000000000",
			`{"error":"Not Found","status":404,"message":"Stream does not exist"}`,
		},
		{
			http.StatusOK,
			"26490481",
			`{"_id":37205889344,"game":"Call of Duty: Modern Warfare","broadcast_platform":"live","community_id":"","community_ids":[],"viewers":36975,"video_height":1080,"average_fps":60,"delay":0,"created_at":"2020-03-20T17:54:25Z","is_playlist":false,"stream_type":"live","preview":{"small":"https://static-cdn.jtvnw.net/previews-ttv/live_user_summit1g-80x45.jpg","medium":"https://static-cdn.jtvnw.net/previews-ttv/live_user_summit1g-320x180.jpg","large":"https://static-cdn.jtvnw.net/previews-ttv/live_user_summit1g-640x360.jpg","template":"https://static-cdn.jtvnw.net/previews-ttv/live_user_summit1g-{width}x{height}.jpg"},"channel":{"mature":false,"status":"GAMIN [ @summit1g ]","broadcaster_language":"en","broadcaster_software":"","display_name":"summit1g","game":"Call of Duty: Modern Warfare","language":"en","_id":26490481,"name":"summit1g","created_at":"2011-12-01T06:33:31.487567Z","updated_at":"2020-03-21T04:16:34.88348Z","partner":true,"logo":"https://static-cdn.jtvnw.net/jtv_user_pictures/200cea12142f2384-profile_image-300x300.png","video_banner":"https://static-cdn.jtvnw.net/jtv_user_pictures/82cfc7d89b6d5401-channel_offline_image-1920x1080.png","profile_banner":"https://static-cdn.jtvnw.net/jtv_user_pictures/summit1g-profile_banner-da83b29c4dede9e4-480.png","profile_banner_background_color":"","url":"https://www.twitch.tv/summit1g","views":340384343,"followers":4368313,"broadcaster_type":"","description":"I'm a competitive CounterStrike player who likes to play casually now and many other games. You will mostly see me play CS, H1Z1, and single player games at night. There will be many other games played on this stream in the future as they come out :D","private_video":false,"privacy_options_enabled":false}}`,
		},
	}

	for _, testCase := range testCases {
		c := newMockClient("cid", newMockHandler(testCase.statusCode, testCase.respBody))
		resp, err := c.GetStream(testCase.slug)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != testCase.statusCode {
			t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		fmt.Printf("%+v\n", resp)

		if resp.StatusCode == http.StatusNotFound {
			if resp.Error != "Not Found" {
				t.Errorf("expected error to be %s, got %s", "Not Found", resp.Error)
			}

			if resp.StatusCode != http.StatusNotFound {
				t.Errorf("expected error code to be %d, got %d", http.StatusNotFound, resp.StatusCode)
			}

			if resp.ErrorMessage != "Stream does not exist" {
				t.Errorf("expected error message to be %s, got %s", "Stream does not exist", resp.ErrorMessage)
			}
		}
	}
}
