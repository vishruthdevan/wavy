audioFiles = ["input1.wav", "input2.wav", "input3.wav"]
audioClips = []

foreach file in audioFiles {
    clip = load(file)
    fadeIn(clip, 3.0)
    append(audioClips, clip)
}

total_duration = getDuration(audioClips[0]) + getDuration(audioClips[1]) + getDuration(audioClips[2])