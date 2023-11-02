class Solution:
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        keychain = rooms[0]
        opened = [False] * len(rooms)
        opened[0] = True

        while keychain:
            key = keychain.pop(0)
            if not opened[key]:
                # open the room and collect the keys
                opened[key] = True
                keys = rooms[key]
                for k in keys:
                    keychain.append(k)

        return all(opened)
