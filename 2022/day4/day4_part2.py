def main():
    file = open("input.txt", "r")
    lines = file.readlines()
    counter = 0
    for line in lines:
        elve1, elve2 = line.strip().split(",")
        start_elve1, end_elve1 = list(map(int, elve1.split("-")))
        start_elve2, end_elve2 = list(map(int, elve2.split("-")))

        if (
            start_elve2 <= start_elve1 <= end_elve2
            or start_elve2 <= end_elve1 <= end_elve2
        ) or (
            start_elve1 <= start_elve2 <= end_elve1
            or start_elve1 <= end_elve2 <= end_elve1
        ):
            counter += 1
    print(counter)


if __name__ == "__main__":
    main()
