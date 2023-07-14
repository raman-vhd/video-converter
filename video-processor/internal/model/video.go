package model

type Video struct {
    VideoID     string           `json:"video_id"`
    Ext         string           `json:"ext"`
    Size        int              `json:"size"`
    Versions    map[string]ConvertedVideo `json:"versions"`
}

type ConvertedVideo struct {
    Size  int    `json:"size"`
    State string `json:"State"`
}

type AMQPMsg struct {
    VideoID string  `json:"videoid"`
    Quality  string `json:"quality"`
    Ext string      `json:"ext"`
}
