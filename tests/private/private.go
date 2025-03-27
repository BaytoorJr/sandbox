package private

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var client = http.Client{}

var notifications = []string{
	"Check out our latest deals! 🎉",
	"New arrivals just dropped – don’t miss out!",
	"Your daily inspiration is here! Tap to read.",
	"Special offer! Limited time only – act fast!",
	"🚀 Boost your productivity today with these tips!",
	"Reminder: Don’t forget to complete your goals!",
	"You have a new friend request! 📬",
	"Only a few hours left! Grab it before it’s gone!",
	"Get your exclusive discount now! Limited stock!",
	"We’ve got something exciting just for you!",
	"Your wishlist item is back in stock!",
	"Flash Sale! Don’t miss this amazing offer!",
	"🌟 Achieve your goals with these quick steps!",
	"You’ve received a reward! Tap to claim!",
	"Just for you – a personalized deal awaits!",
	"⚡ Hurry, your offer expires soon!",
	"Take a break and enjoy a quick read!",
	"Your favorite item is on sale now!",
	"Congratulations! You’ve reached a new level!",
	"Stay tuned! Exciting updates are coming!",
	"Need motivation? Here’s a quick tip for you!",
	"Your order has shipped! Track it here.",
	"📅 Don’t forget to schedule your next task!",
	"Check your app for a special surprise!",
	"Limited time: Free shipping on all orders!",
	"Unlock exclusive content! Tap to explore.",
	"Your subscription renews soon! Stay connected.",
	"Start your day with a positive note! ☀️",
	"New notification! Click here to view it.",
	"Celebrate with us! A special gift inside.",
	"Limited seats available – book now!",
	"Enjoy your favorite items with discounts!",
	"Quick reminder: Check out our latest post!",
	"Ready for a challenge? Start now!",
	"Exclusive content just for you – tap to view!",
	"Your favorite brand has new items in stock!",
	"Enjoy free access to our premium feature!",
	"Good news! Your reward points are updated.",
	"Only a few left – get yours now!",
	"Your package has been delivered! 🎉",
	"Feeling stuck? Here’s some inspiration.",
	"Your order has been processed successfully!",
	"Time for a quick break – check this out!",
	"Something special is waiting for you!",
	"Big savings inside! Don’t miss out!",
	"📢 Important update on your account.",
	"Stay motivated – new tips available now!",
	"Reward yourself with this exclusive deal!",
	"Update available! Tap to download now.",
	"Enjoy our latest articles on the go!",
	"💼 Check out new job opportunities!",
	"Today’s trending news – stay informed!",
	"Your next adventure awaits! Tap to plan.",
	"Big announcement coming soon – stay tuned!",
	"Time-limited offer: Save big today!",
	"Meet our new collection – shop now!",
	"Reminder: Your appointment is tomorrow!",
	"Good news! We’ve extended our sale!",
	"Your feedback is important! Rate us now.",
	"New update: Check out these added features!",
	"Take a sneak peek at our upcoming launch!",
	"Keep going! You’re closer to your goal!",
	"Tap here to explore today’s top picks.",
	"💡 New tips to improve your skills!",
	"Reminder: Update your profile today.",
	"Feeling lucky? Spin the wheel and win!",
	"Invite friends and earn rewards!",
	"Enjoy our content even without WiFi!",
	"You’re one step away from unlocking this!",
	"New products added! See what’s fresh.",
	"Check out our exclusive deals of the day!",
	"🚀 Boost your efficiency with these tools!",
	"Read our latest success stories!",
	"⏰ Time-sensitive offer – act now!",
	"Plan your day with our helpful insights!",
	"Need help? Our support is here for you!",
	"Relax with these fun activities!",
	"Looking for deals? We’ve got you covered!",
	"🔥 Trending now! Don’t miss out!",
	"Your daily goal update: Keep it up!",
	"Tap to save on your next purchase!",
	"New arrivals waiting for you – browse now!",
	"Learn something new every day! Start here.",
	"Good morning! Start your day with us.",
	"New features just launched – check them out!",
	"Get ahead with these pro tips!",
	"Your weekly roundup is ready!",
	"Get inspired – here’s your daily dose!",
	"Reminder: Your trial ends soon.",
	"Personalized recommendations just for you!",
	"Take your skills to the next level!",
	"Limited edition items now available!",
	"Tap here to find nearby events!",
	"You’ve unlocked a new badge – congrats!",
	"Stay on top of your tasks with reminders!",
	"Time for a new adventure – start here!",
	"Find your next read with our book picks!",
	"Surprise! You’ve unlocked a secret feature!",
	"Your day just got better – check this out!",
	"Check out our top stories for today!",
	"Exclusive rewards just for our members!",
	"New challenges are waiting for you!",
	"Discover today’s top trending items!",
	"We’ve got big news! Tap to learn more.",
	"Reminder: You have an unfinished task.",
	"Tap to view your personalized dashboard!",
	"Stay organized with our latest update!",
}

type SendNotificationRequest struct {
	DataMap map[string]string `json:"dataMap"`
}

func SendNotifications() {
	reqMap := map[string]string{
		"title":    "",
		"body":     "",
		"category": "FREEDOM_BUSINESS",
	}

	req := SendNotificationRequest{
		DataMap: reqMap,
	}

	for i := 0; i < 100; i++ {
		req.DataMap["title"] = "Monday Push #" + strconv.Itoa(i)
		req.DataMap["body"] = notifications[i]

		requestBody, err := json.Marshal(req)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("sending request: ", string(requestBody))

		request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "https://ibul.trafficwave.kz/producer/notification/write", bytes.NewBuffer(requestBody))
		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = client.Do(request)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("finished sending 100 requests")
	}
}

func SendCheckAccountAvailability() {
	url := "https://ngx-proxy-ibul.trafficwave.kz/colvir/mobile-service/api/m1/has-account?iinBin=800524302334"
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	request.SetBasicAuth("mobibul", "Bmde8wk0Zzf16WphapTK")

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	response := false

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("url: ", url, "response: ", response)
}
