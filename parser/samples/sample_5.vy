aud1 = load("input1.wav")
format = "aac"

if (format == 'aac') {
    changeFormat(aud1, 'aac')
    export(aud1, 'output1.aac')
} else {
    changeFormat(aud1, 'mp3')
    export(aud1, 'output1.mp3')
}
