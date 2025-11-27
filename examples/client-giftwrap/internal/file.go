package internal

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

type FileMetadata struct {
	Filename string `json:"filename"`
	MimeType string `json:"mimeType"`
	Size     int64  `json:"size"`
}

type Message struct {
	Type     string        `json:"type"`
	Content  string        `json:"content"`
	Metadata *FileMetadata `json:"metadata,omitempty"`
}

func EncodeTextMessage(text string) Message {
	return Message{
		Type:    "text",
		Content: text,
	}
}

func EncodeFileMessage(data []byte, filename, mimeType string) Message {
	return Message{
		Type:    "file",
		Content: base64.StdEncoding.EncodeToString(data),
		Metadata: &FileMetadata{
			Filename: filename,
			MimeType: mimeType,
			Size:     int64(len(data)),
		},
	}
}

func DecodeMessage(msg Message) (msgType string, content []byte, filename string, err error) {
	msgType = msg.Type
	
	switch msg.Type {
	case "text":
		content = []byte(msg.Content)
	case "file":
		content, err = base64.StdEncoding.DecodeString(msg.Content)
		if err != nil {
			return "", nil, "", err
		}
		if msg.Metadata != nil {
			filename = msg.Metadata.Filename
		}
	default:
		return "", nil, "", errors.New("unknown message type")
	}
	
	return msgType, content, filename, nil
}

func EncodeMessageJSON(msg Message) (string, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func DecodeMessageJSON(data string) (Message, error) {
	var msg Message
	err := json.Unmarshal([]byte(data), &msg)
	return msg, err
}
