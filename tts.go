package golang_tts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bmizerany/aws4"
)

const api = "https://polly.us-west-2.amazonaws.com"
const (
	MP3 format = iota
	OGG
)

const (
	NEURAL   engine = "neural"
	STANDARD engine = "standard"
)

const (
	RATE_8000  rate = 8000
	RATE_16000 rate = 16000
	RATE_22050 rate = 22050
)

const (
	Geraint   = "Geraint"
	Gwyneth   = "Gwyneth"
	Mads      = "Mads"
	Naja      = "Naja"
	Hans      = "Hans"
	Marlene   = "Marlene"
	Nicole    = "Nicole"
	Russell   = "Russell"
	Amy       = "Amy"
	Brian     = "Brian"
	Emma      = "Emma"
	Raveena   = "Raveena"
	Ivy       = "Ivy"
	Joanna    = "Joanna"
	Joey      = "Joey"
	Justin    = "Justin"
	Kendra    = "Kendra"
	Kimberly  = "Kimberly"
	Salli     = "Salli"
	Conchita  = "Conchita"
	Enrique   = "Enrique"
	Miguel    = "Miguel"
	Penelope  = "Penelope"
	Chantal   = "Chantal"
	Celine    = "Celine"
	Mathieu   = "Mathieu"
	Dora      = "Dora"
	Karl      = "Karl"
	Carla     = "Carla"
	Giorgio   = "Giorgio"
	Mizuki    = "Mizuki"
	Liv       = "Liv"
	Lotte     = "Lotte"
	Ruben     = "Ruben"
	Ewa       = "Ewa"
	Jacek     = "Jacek"
	Jan       = "Jan"
	Maja      = "Maja"
	Ricardo   = "Ricardo"
	Vitoria   = "Vitoria"
	Cristiano = "Cristiano"
	Ines      = "Ines"
	Carmen    = "Carmen"
	Maxim     = "Maxim"
	Tatyana   = "Tatyana"
	Astrid    = "Astrid"
	Filiz     = "Filiz"
	Aditi     = "Aditi"
	Matthew   = "Matthew"
)

type format int
type rate int
type voice int
type engine string

type TTS struct {
	accessKey string
	secretKey string
	request   request
}

type request struct {
	Engine       string
	LanguageCode string
	OutputFormat string
	SampleRate   string
	Text         string
	VoiceId      string
	TextType     string
}

func New(accessKey string, secretKey string) *TTS {
	return &TTS{
		accessKey: accessKey,
		secretKey: secretKey,
		request: request{
			Engine:       "neural",
			LanguageCode: "en-US",
			OutputFormat: "mp3",
			SampleRate:   "22050",
			Text:         "",
			TextType:     "text",
			VoiceId:      "Brian"}}
}

func (tts *TTS) Format(format format) {
	switch format {
	case MP3:
		tts.request.OutputFormat = "mp3"
	case OGG:
		tts.request.OutputFormat = "ogg_vorbis"
	}
}

func (tts *TTS) Engine(engine engine) {
	tts.request.Engine = fmt.Sprintf("%s", engine)
}

func (tts *TTS) SampleRate(rate rate) {
	tts.request.SampleRate = fmt.Sprintf("%d", rate)
}

func (tts *TTS) Voice(voice string) {
	tts.request.VoiceId = fmt.Sprintf("%s", voice)
}

func (tts *TTS) TextType(textType string) {
	tts.request.TextType = fmt.Sprintf("%s", textType)
}

func (tts *TTS) Language(lang string) {
	tts.request.LanguageCode = fmt.Sprintf("%s", lang)
}

func (tts *TTS) Speech(text string) (data []byte, err error) {
	tts.request.Text = text

	b, err := json.Marshal(tts.request)
	if err != nil {
		return []byte{}, err
	}

	r, _ := http.NewRequest("POST", api+"/v1/speech", bytes.NewReader(b))
	r.Header.Set("Content-Type", "application/json")

	client := aws4.Client{Keys: &aws4.Keys{
		AccessKey: tts.accessKey,
		SecretKey: tts.secretKey}}

	res, err := client.Do(r)
	if err != nil {
		return []byte{}, err
	}

	defer func() {
		err = res.Body.Close()
	}()

	data, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return []byte{}, err
	} else if res.StatusCode != 200 {
		return []byte{}, fmt.Errorf("returned status code: %s %q", res.Status, data)
	}

	return data, nil
}
