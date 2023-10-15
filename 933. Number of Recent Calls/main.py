class RecentCounter:

    def __init__(self):
        self.q = []

    def ping(self, t: int) -> int:
        while len(self.q) > 0 and self.q[0] < t-3000:
            self.q.pop(0)  # dequeue
        self.q.append(t)  # enqueue
        return len(self.q)
