def main():
    file = open("input.txt", "r").read()
    lines = file.split("\n")

    HEIGHT = len(lines)
    WIDTH = len(lines[0])
    VISIBLE = HEIGHT * 2 + WIDTH * 2 - 4

    for element in range(len(lines)):
        line = lines[element]
        if element == 0 or element == HEIGHT - 1:
            continue

        for item in range(len(line)):
            tree = line[item]
            if item == 0 or item == WIDTH - 1:
                continue

            right = (tree > t for t in line[item + 1 :])
            up = (tree > lines[k][item] for k in range(element - 1, -1, -1))
            down = (tree > lines[k][item] for k in range(element + 1, len(lines)))
            left = (tree > t for t in line[:item])

            if all(up) or all(down) or all(left) or all(right):
                VISIBLE += 1

    print(VISIBLE)


if __name__ == "__main__":
    main()
