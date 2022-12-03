from string import ascii_lowercase, ascii_uppercase


def in_both(s1, s2):
    s1_chars = set(s1)
    s2_chars = set(s2)
    result = s1_chars.intersection(s2_chars)
    for item in result:
        return item


def get_priorities():
    result = {}
    for index, character in enumerate(ascii_lowercase):
        result[character] = index + 1
    for index, character in enumerate(ascii_uppercase):
        result[character] = index + 27
    return result


def split_compartments(item):
    string1, string2 = item[: len(item) // 2].strip(), item[len(item) // 2 :].strip()
    return (string1, string2)


def main():
    file = open("input1.txt", "r")
    lines = file.readlines()
    priorities = get_priorities()
    sum_priorities = 0
    for line in lines:
        rucksack_inventory = split_compartments(line)
        mutuality = in_both(rucksack_inventory[0], rucksack_inventory[1])
        sum_priorities += priorities[mutuality]
    print(sum_priorities)


if __name__ == "__main__":
    main()
