package main

import (
	"log"
	"regexp"
	// "strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"encoding/json"
    "fmt"
    "net/http"
    "io"
    "time"
	"math/rand"
	"strconv"
)

// Markdown is send message parse mode
const Markdown = tgbotapi.ModeMarkdown

var cityName = &CityName{}
var weatherInfo = &WeatherInfo{}
var forest = &Forest{}
var callbackConfing = &tgbotapi.CallbackConfig{}
var responseMessage string


type Librari struct {
    Root []struct {
        Nama_program string `json:"NMPRGRM"`
        Nama_unit    string `json:"NMUNIT"`
        Ket    string `json:"KET"`
    } `json:"root"`
}

func check(result string  ) bool {
    // faulty regex   
    // m, err := regexp.MatchString("b\\ello w\\b",result)
    m, err := regexp.MatchString("Penyediaan Gaji dan Tunjangan ASN",result)
    if err != nil {
      fmt.Println("your regex is faulty")
      // you should log it or throw an error 
    //   return err.Error()
    }
    if (m) {
        // fmt.Println("Found it ")
        return true
    } else {
        return false
    }
}

func main() {
	token, err := ReadBotToken("./token.json")

	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil || update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "today":
				inlineButton := tgbotapi.NewInlineKeyboardButtonData("未來一週天氣", "forest")
				inlineKeyboard := []tgbotapi.InlineKeyboardButton{
					inlineButton,
				}
				inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)

				e := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					responseMessage,
				)
				e.BaseEdit.ReplyMarkup = &inlineKeyboardMarkup
				e.ParseMode = Markdown

				bot.Send(e)
			case "forest":
				inlineButton := tgbotapi.NewInlineKeyboardButtonData("今日天氣狀況", "today")
				inlineKeyboard := []tgbotapi.InlineKeyboardButton{
					inlineButton,
				}
				inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)

				e := tgbotapi.NewEditMessageText(
					update.CallbackQuery.Message.Chat.ID,
					update.CallbackQuery.Message.MessageID,
					callbackConfing.Text,
				)
				e.BaseEdit.ReplyMarkup = &inlineKeyboardMarkup
				e.ParseMode = Markdown

				bot.Send(e)
			}

			continue
		}

		command := regexp.MustCompile("/[a-z]+").FindString(update.Message.Text)

		switch command {
		case "/start":
			message := "Hello！ini adalah bot telegram 😉 \nKetik `/help` untuk menampilkan perintah yang tersedia!"
			response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			response.ParseMode = Markdown

			bot.Send(response)
		case "/help":
			message := "Berikut beberapa perintah yang tersedia:\n`/ceksuhu` menampilkan status suhu ruangan DC.\n`/cekpower` menampilkan status power listrik.\n`/cektpp` menampilkan status TPP 😉"
			response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			response.ParseMode = Markdown

			bot.Send(response)
		case "/ceksuhu":
			min := 20
			max := 25
			// fmt.Println(rand.Intn(max - min) + min)
			message := "Suhu ruang server saat ini adalah "
			message += strconv.Itoa(rand.Intn(max - min) + min)
			message += "°C"
			response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			response.ParseMode = Markdown

			bot.Send(response)

		case "/cektpp":
			resp, err := http.Get("http://e-sp2d.bandaacehkota.go.id/storeweb.php?_dc=1677312100805&SKPD=DINAS%20KOMUNIKASI%2C%20INFORMATIKA%20DAN%20STATISTIK&NOSPM=&REKTUJUAN=&NAMAKEGIATAN=&limit=30&token=tcOVdreOvLaPtr%2FDkna3t52Hi6%2B3yZmCpcK6soOBtbuVVbfCsruLtMXDkliH&start=NaN")

			if err != nil {
				time.Sleep(1 * time.Second)
				message := "Error..Tidak ada data!"
				response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				response.ParseMode = Markdown
				bot.Send(response)
			}else{
				defer resp.Body.Close()

				b, _ := io.ReadAll(resp.Body)
			
				// var res map[string]interface{}
				var res map[string]any
				jsonErr := json.Unmarshal(b, &res)
				if jsonErr != nil {
					log.Fatal(jsonErr)
					message := "Maaf ada kesalahan!"
					response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
					response.ParseMode = Markdown
					bot.Send(response)
				}else{
					var librariesInformation Librari
					err := json.Unmarshal(b, &librariesInformation)
					if err != nil {
						message := "error unmarshaling json"
						response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
						response.ParseMode = Markdown
						bot.Send(response)
						log.Fatal("error unmarshaling json: ", err)
					}

					var cair bool
					cair = false
					for _, data := range librariesInformation.Root {
						// fmt.Printf("%s (%d)\n", person.Name, person.Age)
						if check(data.Nama_program) {
							// strBuilder.WriteString("%s, %s %s \n", data.Nama_unit, data.Nama_program, data.Ket)
							// fmt.Printf("%s, %s %s \n", data.Nama_unit, data.Nama_program, data.Ket)
							cair = true
						}
						
						// fmt.Println(data.Nama_unit)
					}

					if cair {
						message := "TPP bulan ini sudah cair!"
						response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
						response.ParseMode = Markdown
						bot.Send(response)
					}else{
						message := "TPP bulan ini belum cair!"
						response := tgbotapi.NewMessage(update.Message.Chat.ID, message)
						response.ParseMode = Markdown
						bot.Send(response)
					}
					
				}
			} // end if err http
		default:
			if update.Message.Location != nil {
				weatherAPIURL := BuildURL(*update.Message.Location)
				body, err := HTTPGet(weatherAPIURL)

				if err != nil {
					log.Fatal(err)
				}

				forestInfo, _ := weatherInfo.HandleQueryResult(body)
				forestResponse := forest.HandleQueryForest(forestInfo)

				responseMessage = weatherInfo.ResponseWeatherText(weatherInfo)

				callbackConfing.Text = forestResponse

				inlineButton := tgbotapi.NewInlineKeyboardButtonData("未來一週天氣", "forest")
				inlineKeyboard := []tgbotapi.InlineKeyboardButton{
					inlineButton,
				}
				inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup(inlineKeyboard)

				response := tgbotapi.NewMessage(update.Message.Chat.ID, responseMessage)
				response.BaseChat.ReplyMarkup = inlineKeyboardMarkup
				response.ParseMode = Markdown

				bot.Send(response)
			}
		}
	}
}
