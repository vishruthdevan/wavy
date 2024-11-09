aud1 = load("input1.wav")
format = "aac"

if (format == 'wav') {
    export(aud1, 'output1.wav')
} else if (format == 'mp3') {
    changeFormat(aud1, format: 'wav')
    export(aud1, 'output1.mp3')
} else if (format == 'aac') {
    export(aud1, 'output1.aac')
} else {
    print("Invalid format)
}