scaleVolume = function(volume, scale) {
    return volume * scale
}

aud1 = load("input1.wav")

louder_aud1 = scaleVolume(aud1, 2.0)
fainter_aud1 = scaleVolume(aud1, 0.5)

export(louder_aud1, 'louder_aud1.wav')
export(fainter_aud1, 'fainter_aud1.wav')