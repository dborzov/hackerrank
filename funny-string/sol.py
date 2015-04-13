try:
    from contest_helpers import log
except:
    def log(*args):
        pass


def is_funny(arr):
    log("is_funny called for ", arr)
    for i in range(1, (len(arr)+1)/2):
        if not abs(arr[i] - arr[i-1]) == abs(arr[-1-i] - arr[-i]):
            return False
    return True


for t in range(input()):
    string = raw_input().rstrip()
    if is_funny([ord(sym) for sym in string]):
        print "Funny"
    else:
        print "Not Funny"
