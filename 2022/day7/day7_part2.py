import os.path
import re
import sys

SIZE_BOUND = 100000


def main():
    file = open("input.txt", "r").read()
    lines = file.split("\n")

    sizes = {}
    current = []

    for command in lines:
        if command.startswith("$"):
            splittet_commands = command.split(" ")
            splittet_commands.pop(0)
            if splittet_commands[0] == "cd":
                if splittet_commands[1] == "..":
                    current.pop()
                else:
                    current.append(splittet_commands[1])
        else:
            splittet_commands = command.split(" ")
            if splittet_commands[0] != "dir" and splittet_commands[0] != "":
                if not tuple(current) in sizes:
                    sizes[tuple(current)] = 0
                sizes[tuple(current)] += int(splittet_commands[0])

    total_sizes = {}
    for size in sizes.keys():
        current_path = ""
        for dir in size:
            if dir != "/":
                current_path += "/" + dir
            if current_path not in total_sizes:
                total_sizes[current_path] = 0
            total_sizes[current_path] += sizes[size]

    TOTAL_DISK = 70000000
    TOTAL_UPDATE = 30000000

    total_used = total_sizes[""]
    total_unused = TOTAL_DISK - total_used
    total_needed = TOTAL_UPDATE - total_unused

    selected_sizes = {
        key: value for key, value in total_sizes.items() if value >= total_needed
    }

    min = TOTAL_DISK
    for s in selected_sizes:
        if selected_sizes[s] < min:
            min = selected_sizes[s]
    print(min)


if __name__ == "__main__":
    main()
