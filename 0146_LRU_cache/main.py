# coding=utf-8

KEY_NOT_EXIST_FLAG = -1


class Node(object):
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.pre = None
        self.next = None

    def remove_from_link(self):
        if self.pre is not None:
            self.pre.next = self.next

        if self.next is not None:
            self.next.pre = self.pre


class LRUCache(object):
    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.cache = {}
        self.capacity = capacity
        self.root = Node(0, 0)
        self.root.next = self.root
        self.root.pre = self.root

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key not in self.cache:
            return KEY_NOT_EXIST_FLAG
        node = self.cache[key]
        if node != KEY_NOT_EXIST_FLAG:
            self.move_to_head(node)

        return node.val

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: None
        """
        if key not in self.cache:
            if len(self.cache) >= self.capacity:
                self.remove_tail()
            self.cache[key] = Node(key=key, val=value)

        self.cache[key].val = value

        self.move_to_head(self.cache[key])

    def move_to_head(self, node):
        node.remove_from_link()
        node.next = self.root.next
        node.pre = self.root
        self.root.next = node
        node.next.pre = node

    def remove_tail(self):
        if self.root.next == self.root:
            return

        target = self.root.pre
        target.remove_from_link()

        del self.cache[target.key]


if __name__ == "__main__":
    cache = LRUCache(2)
    cache.put(1, 1)
    cache.put(2, 2)
    assert cache.get(1) == 1
    cache.put(3, 3)
    assert cache.get(2) == -1
    cache.put(4, 4)
    assert cache.get(1) == -1
    assert cache.get(3) == 3
    assert cache.get(4) == 4
