from __future__ import print_function
import sys

def log(*args, **kwargs):
    """
        for printing all the execution progress messages into stderr
        (and the output results go to stdout)
    """
    kwargs["file"] = sys.stderr
    print(*args, **kwargs)