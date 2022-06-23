# Golang Text-to-Speech package base on AWS Polly

## Getting access

Just go to [https://console.aws.amazon.com/iam/home#/security_credential](https://console.aws.amazon.com/iam/home#/security_credential) (if haven't registered user then to have sign up) and open **Access Keys** tab

## Settings

```go
polly := golang_tts.New("AKIBI2FJJVG77M7OC3DQ", "/PEiT4T+27zG7E0Z4+8EJHASn92Au7JWMNrGwR8Z")

polly.Format(golang_tts.MP3)
polly.Voice(golang_tts.Brian)

text := "To read or not to read"
bytes, err := polly.Speech(text)
if err != nil {
    panic(err)
}

ioutil.WriteFile("./result.mp3", bytes, 0644)
```

## List of all methods

* golang_tts.New(accessKey, secretKey) - initializes new TTS client
* golang_tts.Engine(golang_ttl.NEURAL) - sets engine (standard is default)
* golang_tts.Format(golang_ttl.MP3) - sets output format (mp3 is default)
* golang_tts.SampleRate(golang_ttl.RATE_22050) - sets sample rate (22050 is default)
* golang_tts.Language("en-GB") - sets language (en-US is default)
* golang_tts.Voice(golang_ttl.Brian) - sets voice (Brian is default)
* golang_tts.Speech("text to speech") - sends text to synthesize
