class Solution(object):
    def canVisitAllRooms(self, rooms):
        """
        :type rooms: List[List[int]]
        :rtype: bool
        """
        if len(rooms) == 0:
            return True

        room_visited = [False] * len(rooms)
        # enter first room
        self.enter_room(rooms, 0, room_visited)

        return all(room_visited)

    def enter_room(self, rooms, key, room_visited):
        room_visited[key] = True

        if len(rooms[key]) == 0:
            return

        for next_key in rooms[key]:
            # next room is not visit before
            if room_visited[next_key] is False:
                self.enter_room(rooms, next_key, room_visited)


if __name__ == "__main__":
    s = Solution()
    assert s.canVisitAllRooms([[1], [2], [3], []]) is True
    assert s.canVisitAllRooms([[1, 3], [3, 0, 1], [2], [0]]) is False
