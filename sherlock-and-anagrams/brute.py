try:
    from contest_helpers import log
except:
    def log(*args):
        pass

from collections import Counter


def anagramic_pairs(string):
    count = 0
    for n in range(1,len(string)):
        substrings= Counter()
        for i in range(len(string)-n+1):
            substrings["".join(sorted([sym for sym in string[i:i+n]]))] += 1
        log(" all the combinations for n=",n,": ", substrings)
        for substr in substrings:
            count += (substrings[substr] * (substrings[substr]-1))/2
    return count

for t in range(input()):
    print anagramic_pairs(raw_input())
