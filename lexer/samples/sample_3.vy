audio1 = load('audio1.wav')
duration = 0.0
for ( duration < 10.0) {
    audio1 = loop(audio1, 1)
    duration = getDuration(audio1)
}

export(audio1, "looped_audio1.wav")