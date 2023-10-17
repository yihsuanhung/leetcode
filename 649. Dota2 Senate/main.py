class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        radiant_q, dire_q = [], []
        senate_list = list(senate)
        for i in range(len(senate_list)):
            if senate_list[i] is "R":
                radiant_q.append(i)
            else:
                dire_q.append(i)
        next_round = len(senate)
        while radiant_q and dire_q:
            if radiant_q[0] < dire_q[0]:
                radiant_q.append(next_round)
            else:
                dire_q.append(next_round)
            next_round += 1
            radiant_q.pop(0)
            dire_q.pop(0)
        return "Radiant" if radiant_q else "Dire"
