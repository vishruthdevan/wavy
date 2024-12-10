let load = fn(file) {
    puts("Loading file: " + file);
};

let changeFormat = fn(audio, format) {
    puts("Changing format of audio to: " + format);
};

let export = fn(audio, outputFile) {
    puts("Exporting audio to: " + outputFile);
    return outputFile;
};

let aud1 = "input1.wav";
let format = "aac";

load(aud1);
changeFormat(aud1, "aac");
let output1 = export(aud1, "output1.aac");

puts(output1)
output1