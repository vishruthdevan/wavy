let reverseString = fn(s, index) {
    if (index < 0) {
        "";
    } else {
        s[index] + reverseString(s, index - 1);
    }
};

let original = ["W", "a", "v", "y", " ", "L", "a", "n", "g", "u", "a", "g", "e"];
let reversed = reverseString(original, len(original) - 1);

puts("Original: ");
puts(original);
puts("Reversed: ");
puts(reversed);
reversed;