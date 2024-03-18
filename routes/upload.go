package routes

import (
	"DiscordUpload/config"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"time"
)

type Attachment struct {
	Id                 string `json:"id"`
	Filename           string `json:"filename"`
	Size               int    `json:"size"`
	Url                string `json:"url"`
	ProxyUrl           string `json:"proxy_url"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	ContentType        string `json:"content_type"`
	Placeholder        string `json:"placeholder"`
	PlaceholderVersion int    `json:"placeholder_version"`
}

type DiscordResponse struct {
	Id        string `json:"id"`
	Type      int    `json:"type"`
	Content   string `json:"content"`
	ChannelId string `json:"channel_id"`
	Author    struct {
		Id            string      `json:"id"`
		Username      string      `json:"username"`
		Avatar        interface{} `json:"avatar"`
		Discriminator string      `json:"discriminator"`
		PublicFlags   int         `json:"public_flags"`
		Flags         int         `json:"flags"`
		Bot           bool        `json:"bot"`
		GlobalName    interface{} `json:"global_name"`
	} `json:"author"`
	Attachments     []Attachment  `json:"attachments"`
	Embeds          []interface{} `json:"embeds"`
	Mentions        []interface{} `json:"mentions"`
	MentionRoles    []interface{} `json:"mention_roles"`
	Pinned          bool          `json:"pinned"`
	MentionEveryone bool          `json:"mention_everyone"`
	Tts             bool          `json:"tts"`
	Timestamp       time.Time     `json:"timestamp"`
	EditedTimestamp interface{}   `json:"edited_timestamp"`
	Flags           int           `json:"flags"`
	Components      []interface{} `json:"components"`
	WebhookId       string        `json:"webhook_id"`
}

var domain = "http://" + config.ReadEnv("DOMAIN") + ":" + config.ReadEnv("PORT") + "/"

func Upload(writer http.ResponseWriter, request *http.Request) {
	startTime := time.Now()

	var filePaths = make(map[string]string)

	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// get the files from the request
	files := request.MultipartForm.File["files[]"]

	// loop through the files
	for _, file := range files {

		fileFolder, err := uuid.NewUUID()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return

		}

		uploadedFile, err := file.Open()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("storage/"+fileFolder.String(), os.ModePerm)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// create a new file in the storage folder
		newFile, err := os.Create("storage/" + fileFolder.String() + "/" + file.Filename)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// copy the uploaded file to the new file
		_, err = newFile.ReadFrom(uploadedFile)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// close the new file
		err = newFile.Close()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// close the uploaded file
		err = uploadedFile.Close()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		filePaths[file.Filename] = "storage/" + fileFolder.String() + "/" + file.Filename
	}

	var response DiscordResponse
	response.Id = "1234567890"
	response.Type = 0
	response.Content = "Uploaded files"
	response.ChannelId = "1234567890"
	response.Author.Id = "1234567890"
	response.Author.Username = "DiscordUpload"
	response.Author.Avatar = nil
	response.Author.Discriminator = "1234"
	response.Author.PublicFlags = 0
	response.Author.Flags = 0
	response.Author.Bot = true
	response.Author.GlobalName = nil

	for originalName, name := range filePaths {
		var attachedFile Attachment
		attachedFile.Id = "1234567890"
		attachedFile.Filename = originalName
		attachedFile.Size = 1234567890
		attachedFile.Url = domain + name
		attachedFile.ProxyUrl = domain + name
		attachedFile.Width = 0
		attachedFile.Height = 0
		attachedFile.ContentType = "image/png"
		attachedFile.Placeholder = domain + name
		attachedFile.PlaceholderVersion = 1

		response.Attachments = append(response.Attachments, attachedFile)
	}

	response.Embeds = nil
	response.Mentions = nil
	response.MentionRoles = nil
	response.Pinned = false
	response.MentionEveryone = false
	response.Tts = false
	response.Timestamp = time.Now()
	response.EditedTimestamp = nil
	response.Flags = 0
	response.Components = nil
	response.WebhookId = "1234567890"

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	log.Printf("\nUploaded %o files\nDuration: %s", len(filePaths), time.Since(startTime))

	// send the response
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
