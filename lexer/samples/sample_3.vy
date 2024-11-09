audio1 = load('audio1.wav')

for {
    audio1 = loop(audio1, count: 1)
    duration = getDuration(audio1)
    if (duration > 10.0) {
        break
    }
}

export(audio1, 'looped_audio1.wav')