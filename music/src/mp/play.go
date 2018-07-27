package mp

import "fmt"

type Player interface {
    Play(source string)
}

func Play(source, mType string) {
    var p Player

    switch mType {
        case "MP3":
            p = &MP3Player{}
        case "WAV":
            p = &WAVPlayer{}
        default:
            fmt.Println("Unsuppoerted music type", mType)
            return
    }

    p.Play(source)
}
