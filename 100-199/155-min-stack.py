class MinStack:
    # @param x, an integer
    def __init__(self):
        self.stack = []
    # @return an integer
    def push(self, x):
        self.stack.append(x)

    # @return nothing
    def pop(self):
        self.stack.pop()

    # @return an integer
    def top(self):
        return self.stack[-1]

    # @return an integer
    def getMin(self):
        return min(self.stack)