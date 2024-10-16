aud1 = load("input1.wav")
aud2 = load("input2.wav")

aud1 = fadeIn(aud1, duration: 3.0)
aud2 = fadeOut(aud2, duration: 3.0)

aud1 = trim(aud1, start: 5.0, end: 30.0)
aud = trim(aud2, start: 5.0, end: 30.0)

res = join(aud1, aud2)

123res = res