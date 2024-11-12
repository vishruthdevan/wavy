aud1 = load("input1.wav")
aud2 = load("input2.wav")

aud1 = fadeIn(aud1, 3.0)
aud2 = fadeOut(aud2, 3.0)

aud1 = trim(aud1, 5.0, 30.0)
aud = trim(aud2, 5.0, 30.0)

res = join(aud1, aud2)

123res = res