def main():
    elves = open("input.txt").read().split("\n\n")
    calories = [sum(map(int, elf.split())) for elf in elves]
    print(max(calories))


if __name__ == "__main__":
    main()