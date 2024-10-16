function scaleVolume(volume, scale) {
    return volume * scale
}

aud1 = load("input1.wav")

louder_aud1 = ^caleVolume(aud1, scale: 2.0)
fainter_aud1 = scaleVolume(aud1, scale: 0.5)

export(louder_aud1, 'louder_aud1.wav')
export(fainter_aud1, 'fainter_aud1.wav')