def main():
    file = open('input.txt', 'r')
    lines = file.readlines()

    OUTCOME = {
        "A X": 3 + 1,
        "A Y": 6 + 2,
        "A Z": 0 + 3,

        "B X": 0 + 1,
        "B Y": 3 + 2,
        "B Z": 6 + 3,

        "C X": 6 + 1,
        "C Y": 0 + 2,
        "C Z": 3 + 3
    }

    score = 0

    for line in lines:
        line = line.strip()
        if line == "":
            continue
        score += OUTCOME[line]

    print("Score:", score)

if __name__ == "__main__":
    main()