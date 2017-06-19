from itertools import permutations


def main():
    inp_values = []
    with open("input.txt") as f:
        for line in f:
            ln = []
            lst = line.split()
            for elem in lst:
                ln.append(int(elem))
            inp_values.append(ln)
    for line in inp_values:
        ops, num = get_ops(line)
        buff = print_ops_nums(num, ops)
        print(buff)


def print_ops_nums(nums, ops):
    result = nums[-1]
    nums = nums[:-1]
    buff = "{}".format(nums[0])
    nums = nums[1:]
    for num, op in zip(nums, ops):
        buff = buff + op + str(num)
    buff + '=' + str(result)
    return buff


def get_ops(inp):
    buff = inp[-1]
    inp = inp[:-1]
    for num in permutations(inp):
        for ops in all_ops_combinations(len(num)):
            result = build_and_calc(ops, num)
            if result == buff:
                return ops, num + (buff,)


def all_ops_combinations(length):
    OPS = [
        ['+'],
        ['-'],
        ['*'],
        ['/'],
    ]
    if length == 2:
        for op in OPS:
            yield op
        return

    for op in OPS:
        for comb in all_ops_combinations(length - 1):
            yield op + comb


def build_and_calc(ops, lst):
    buff = lst[0]
    lst = lst[1:]
    for number, oper in zip(lst, ops):
        if oper == '+':
            buff += number
        if oper == '-':
            buff -= number
        if oper == '*':
            buff *= number
        if oper == '/':
            buff /= number
    return buff


if __name__ == '__main__':
    main()
