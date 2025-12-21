# Code Cracker CLI

A fun command-line number guessing game built in Go where you crack the secret code by guessing the correct digits and their positions.

## 🎮 Game Overview

Code Cracker challenges you to guess a randomly generated number within the fewest attempts possible. After each guess, you'll receive feedback in the form of:

- **Bulls-eye** 🎯: Correct digit in the correct position
- **Hits** 🎲: Correct digit in the wrong position

Use this feedback strategically to deduce the secret code!

## 🚀 Features

- **Three Difficulty Levels**:
  - Easy (3 digits)
  - Medium (4 digits)
  - Hard (5 digits)
- **Real-time Feedback**: Get immediate bulls-eye and hit counts after each guess
- **Performance Tracking**: See your total attempts and completion time
- **Input Validation**: Ensures guesses match the required digit length

## 📋 Prerequisites

- Go 1.25.2 or higher

## 🔧 Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd code-cracker-cli
```

2. Build the application:
```bash
go build
```

## 🎯 How to Play

1. Run the game:
```bash
./code-cracker-cli
```

2. Select your difficulty level (1-3)

3. Start guessing! Enter a number with the correct number of digits

4. Use the feedback to refine your next guess:
   - More bulls-eyes mean you have correct digits in correct positions
   - Hits indicate correct digits that need to be repositioned

5. Keep guessing until you crack the code!

## 📖 Example Gameplay

```
Welcome to the Code Cracker!

Please select the difficulty level:
1. Easy (3 digits)
2. Medium (4 digits)
3. Hard (5 digits)

Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Guess 1: 1234
Incorrect! You got 1 bullseyes and 2 hits.

Guess 2: 5678
Incorrect! You got 0 bullseyes and 1 hits.

Guess 3: 1324
Congratulations! You guessed the correct number in 3 attempts and 15.42 seconds.
```

## 🏗️ Project Structure

```
code-cracker-cli/
├── main.go          # Entry point and game loop
├── checker.go       # Logic for comparing guesses with target
├── difficulty.go    # Difficulty selection handler
├── random.go        # Random number generation
└── go.mod          # Module dependencies
```

## 🧩 Key Components

### checker.go
Compares player guesses with the target number, calculating bulls-eyes and hits using efficient map-based counting to handle duplicate digits correctly.

### difficulty.go
Handles difficulty selection with input validation, mapping user choices to the appropriate number of digits.

### random.go
Generates random numbers with proper zero-padding to ensure consistent digit lengths across all difficulty levels.

## 🎓 Game Strategy Tips

- Start with guesses that cover different digits to maximize information
- Track which digits are confirmed through bulls-eyes
- Use hits to determine which digits need repositioning
- Eliminate impossible combinations based on previous feedback

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 📄 License

This project is open source and available under the MIT License.

## 🎉 Have Fun!

Challenge yourself to crack the code in the fewest attempts possible. Happy guessing!