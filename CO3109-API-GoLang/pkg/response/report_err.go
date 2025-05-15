package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/pkg/telegram"
)

func sendServerTelegramMessageAsync(message string, t telegram.TeleBot, chatID int64) {
	go func() {
		splitMessages := splitMessageForTelegram(message)
		for _, message := range splitMessages {
			_, err := t.SendMessage(chatID, message)
			if err != nil {
				log.Printf("Error sending Telegram message: %v\n", err)
			}
		}
	}()
}

func splitMessageForTelegram(message string) []string {
	const maxMessageLength = 4096
	var messages []string
	for len(message) > maxMessageLength {
		messages = append(messages, message[:maxMessageLength])
		message = message[maxMessageLength:]
	}
	messages = append(messages, message)
	return messages
}

func buildInternalServerErrorDataForReportBug(errString string, backtrace []string, c *gin.Context) string {
	url := c.Request.URL.String()
	method := c.Request.Method
	params := c.Request.URL.Query().Encode()

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return ""
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	body := string(bodyBytes)

	var sb strings.Builder
	sb.WriteString("================ SMAP SERVICE ERROR ================\n")
	sb.WriteString(fmt.Sprintf("Route   : %s\n", url))
	sb.WriteString(fmt.Sprintf("Method  : %s\n", method))
	sb.WriteString("---------------------------------------------------------------------------\n")

	if len(c.Request.Header) > 0 {
		sb.WriteString("Headers :\n")
		for key, values := range c.Request.Header {
			sb.WriteString(fmt.Sprintf("    %s: %s\n", key, strings.Join(values, ", ")))
		}
		sb.WriteString("---------------------------------------------------------------------------\n")
	}

	if params != "" {
		sb.WriteString(fmt.Sprintf("Params  : %s\n", params))
	}

	if body != "" {
		sb.WriteString("Body    :\n")
		// Pretty print JSON if possible
		var prettyBody bytes.Buffer
		if err := json.Indent(&prettyBody, bodyBytes, "    ", "  "); err == nil {
			sb.WriteString(prettyBody.String() + "\n")
		} else {
			sb.WriteString("    " + body + "\n")
		}
		sb.WriteString("---------------------------------------------------------------------------\n")
	}

	sb.WriteString(fmt.Sprintf("Error   : %s\n", errString))

	if len(backtrace) > 0 {
		sb.WriteString("\nBacktrace:\n")
		for i, line := range backtrace {
			sb.WriteString(fmt.Sprintf("[%d]: %s\n", i, line))
		}
	}

	sb.WriteString("====================================================\n")
	return sb.String()
}
