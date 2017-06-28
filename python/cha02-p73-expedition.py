import sys
import heapq


def solve():
    n, l, p = map(int, sys.stdin.readline().split())
    a = list(map(int, sys.stdin.readline().split()))
    b = list(map(int, sys.stdin.readline().split()))
    a.append(l)
    b.append(0)
    n += 1

    que = []

    ans = 0
    pos = 0
    tank = p

    for i in range(n):
        d = a[i] - pos
        while tank - d < 0:
            if len(que) == 0:
                print(-1)
                return
            tank += -que[0]
            heapq.heappop(que)
            ans += 1

        tank -= d
        pos = a[i]
        heapq.heappush(que, -b[i])
    print(ans)


if __name__ == '__main__':
    solve()