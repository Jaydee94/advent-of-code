def main():
    file = open('input.txt', 'r')
    lines = file.readlines()

    OUTCOME = {
        "A X": 0 + 3,
        "A Y": 3 + 1,
        "A Z": 6 + 2,

        "B X": 0 + 1,
        "B Y": 3 + 2,
        "B Z": 6 + 3,

        "C X": 0 + 2,
        "C Y": 3 + 3,
        "C Z": 6 + 1
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