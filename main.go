package main


import( 
	"fmt"
	"context"
	"log"
	"strconv"
	"os"
	"github.com/shomali11/slacker"
	"github.com/joho/godotenv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main(){
	godotenv.Load(".env")
	SLACK_BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	if SLACK_BOT_TOKEN == ""{
		log.Fatal("SLACK BOT TOKEN is not found in the environemnt")
	}
	SLACK_APP_TOKEN := os.Getenv("SLACK_APP_TOKEN")
	if SLACK_APP_TOKEN == "" {
		log.Fatal("SLACK APP TOKEN is not found in the environment")
	}

	bot := slacker.NewClient(SLACK_BOT_TOKEN,SLACK_APP_TOKEN)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{ 
		Description :"yob calculator",
		Examples : []string{"my yob is 2020"},
		Handler : func(botCtx slacker.BotContext,request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob ,err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}


}