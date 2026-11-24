# Level Order Traversal (and Spiral)

> Striver SDE Sheet 116 · Day 18 Binary Tree II — BFS with the queue length snapshotted per round to cut levels; spiral is the same plus alternate reversal. Go, stdlib only, tested.

## Problem

Return the tree's values grouped by level, left to right — plus the spiral variant, where alternate levels run right to left.

## Approach

Plain BFS loses level boundaries: the queue mixes the current level's tail with the next level's head. The standard fix is to snapshot `len(queue)` before each round — that count is exactly the current level's population, so draining that many nodes (while enqueueing their children behind) yields one clean level per iteration. The spiral form doesn't need a deque or direction-aware traversal: reverse every second level after the fact, same asymptotics, a third of the code. Level order is also the serialization order used by LeetCode-style tree literals — worth having as a tool, not just an answer.

**Complexity:** O(n) time, O(widest level) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
