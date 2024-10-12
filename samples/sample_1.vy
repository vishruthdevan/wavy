audioFiles = ["input1.wav", "input2.wav", "input3.wav"]
audioClips = []

for (file in audioFiles) {
clip = load(file)
fadeIn(clip, duration: 3.0)
audioClips.append(clip)
}

duration = getDuration(audioClips[0])

if (duration > 30.0) {
trimmedClip = trim(audioClips[0], start: 5.0, end: 30.0)
} else {
trimmedClip = trim(audioClips[0], start: 5.0, end: duration)
}

volumeAdjustedClip = adjustVolume(trimmedClip, db: 40)

finalClip = join(audioClips)

export(appendedClip, 'output.wav')