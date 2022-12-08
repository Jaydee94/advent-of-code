def findStartOfPacketMarker(lines, marker_length):
    for element in range(marker_length - 1, len(lines)):
        chars = []
        for item in range(element - (marker_length - 1), element + 1):
            if lines[item] not in chars:
                chars.append(lines[item])
            else:
                break
        if len(chars) == marker_length:
            return element + 1


def main():
    with open("input.txt") as file:
        lines = file.readline()
        print(findStartOfPacketMarker(lines, 14))


if __name__ == "__main__":
    main()
