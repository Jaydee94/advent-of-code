
def main():
    with open("input.txt", "r") as file:
        input = file.read()

    stack, commands = input.split("\n\n")
    stacks = {}

    for line in stack.split("\n")[-2::-1]:
        for idx, element in enumerate(line[1:-1:4]):
            if element == " ":
                continue
            if idx + 1 not in stacks:
                stacks[idx + 1] = []
            stacks[idx + 1].append(element)


    stacks = stacks.copy()

    for line in commands.split("\n"):
        _, amount, _, source, _, destination = line.strip().split(" ")
        amount, source, destination = int(amount), int(source), int(destination)
        stacks[source], stacks[destination] = (
            stacks[source][:-amount],
            stacks[destination] + stacks[source][-amount:],
        )

    print("".join([x[-1] for x in stacks.values()]))


if __name__ == '__main__':
    main()