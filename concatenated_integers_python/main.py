from typing import List, Tuple, Any

# https://www.reddit.com/r/dailyprogrammer/comments/69y21t/20170508_challenge_314_easy_concatenated_integers/

def main():
    min_max = []
    with open("input.txt") as f:
        for line in f:
            elems = line.split()
            int_elems = [int(elem) for elem in elems]
            min_max.append(process_line(int_elems))
    with open("output.txt", mode='w') as f:
        for minim, maxmim in min_max:
            f.write("{} {}\n".format(minim, maxmim))

def concat_numbers(nums: List[int]) -> int:
    out = ""
    for num in nums:
        out += str(num)
    return int(out)

def permutations(l: List[Any]) -> List[List[Any]]:
    if len(l) == 0:
        return []

    if len(l) == 1:
        return [l]

    perms = []
    for head in l:
        tail = l[:]
        tail.remove(head)
        tail_perms = permutations(tail)
        head_perms = [[head]+p for p in tail_perms]
        perms += head_perms
    return perms

def process_line(nums: List[int]) -> Tuple[int, int]:
    permut = permutations(nums)
    out = [concat_numbers(elem) for elem in permut]
    return (min(out), max(out))

if __name__ == '__main__':
    main()
