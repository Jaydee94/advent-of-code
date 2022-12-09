def main():
    file = open("input.txt", "r").read()
    lines = file.split("\n")

    HEIGHT = len(lines)
    WIDTH = len(lines[0])

    count = 0

    for element in range(len(lines)):
        line = lines[element]
        if element == 0 or element == HEIGHT - 1:
            continue

        for item in range(len(line)):
            tree = line[item]
            if item == 0 or item == WIDTH - 1:
                continue

            for right in range(item + 1, WIDTH):
                if line[right] >= tree:
                    break

            for left in range(item - 1, -1, -1):
                if line[left] >= tree:
                    break

            for down in range(element + 1, HEIGHT):
                if lines[down][item] >= tree:
                    break

            for up in range(element - 1, -1, -1):
                if lines[up][item] >= tree:
                    break

            score = (right - item) * (item - left) * (down - element) * (element - up)

            if score > count:
                count = score

    print(count)


if __name__ == "__main__":
    main()
