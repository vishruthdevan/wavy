let samples = [2, 4, -1, 5];
let volumeBoost = fn(sample) { return sample * 10; };

let adjustedSample1 = volumeBoost(samples[0]);
let adjustedSample2 = volumeBoost(samples[1]);
puts(adjustedSample1);
puts(adjustedSample2);