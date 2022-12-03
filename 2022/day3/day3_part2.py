from string import ascii_lowercase, ascii_uppercase


def in_all_groups(s1, s2, s3):
    s1_chars = set(s1)
    s2_chars = set(s2)
    s3_chars = set(s3)
    result = s1_chars.intersection(s2_chars).intersection(s3_chars)
    for item in result:
        return item


def get_priorities():
    result = {}
    for index, character in enumerate(ascii_lowercase):
        result[character] = index + 1
    for index, character in enumerate(ascii_uppercase):
        result[character] = index + 27
    return result


def split_groups(lines):
    return [lines[x : x + 3] for x in range(0, len(lines), 3)]


def main():
    file = open("input1.txt", "r")
    lines = file.readlines()
    groups = split_groups(lines)
    priorities = get_priorities()
    sum_priorities = 0
    for group in groups:
        badge_item = in_all_groups(group[0].strip(), group[1].strip(), group[2].strip())
        sum_priorities += priorities[badge_item]
    print(sum_priorities)


if __name__ == "__main__":
    main()
