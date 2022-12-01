def main():
    elves = open("input.txt").read().split("\n\n")
    calories = [sum(map(int, elf.split())) for elf in elves]
    print(sum(sorted(calories)[-3:]))


if __name__ == "__main__":
    main()