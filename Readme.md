# Golang Text-to-Speech package base on AWS Polly

## Settings

```go
polly := golang_tts.New("AKIBI2FJJVG77M7OC3DQ", "/PEiT4T+27zG7E0Z4+8EJHASn92Au7JWMNrGwR8Z")

polly.Format(golang_tts.MP3)
polly.Voice(golang_tts.Brian)

text := "To read or not to read"
bytes, err := polly.Speech(text)
if err != nil {
    t.Fatal(err)
}

ioutil.WriteFile("./result.mp3", bytes, 0644)
```

## List all methods

* golang_tts.New(accessKey, secretKey) - initializes new TTS client
* golang_tts.Format(golang_ttl.MP3) - sets output format (mp3 is default)
* golang_tts.SampleRate(golang_ttl.RATE_22050) - sets sample rate (22050 is default)
* golang_tts.Voice(golang_ttl.Brian) - sets voice (Brian is default)
* golang_tts.Speech("text to speech") - sends text to synthesize
