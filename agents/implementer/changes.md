# Changes: bottle-song exercise

## Implemented `Recite` function in `bottle_song.go`

- **`Recite(startBottles, takeDown int) []string`**: Generates the lyrics for the "Ten Green Bottles" song. Takes the starting number of bottles and how many verses to recite. Returns a slice of strings, one per line, with blank lines separating verses.
- **`numberWord(n int) string`**: Converts an integer (0–10) to its English word equivalent using a lookup slice.
- **`bottleStr(n int) string`**: Returns "bottle" (singular) when n == 1, otherwise "bottles" (plural).
- Uses the `Title` function (provided in the test file) to capitalize the first letter of the number word for the first two lines of each verse.
