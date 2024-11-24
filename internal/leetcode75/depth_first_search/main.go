package main

// 841. 钥匙和房间
// https://leetcode.cn/problems/keys-and-rooms
func canVisitAllRooms(rooms [][]int) bool {
	visitedRooms := make(map[int]struct{}, len(rooms))
	visitedRooms[0] = struct{}{}

	nextRooms := rooms[0]

	for len(nextRooms) > 0 {
		nextSearchRooms := make([]int, 0, len(nextRooms))
		for _, roomIdx := range nextRooms {
			if _, ok := visitedRooms[roomIdx]; !ok {
				visitedRooms[roomIdx] = struct{}{}
				nextSearchRooms = append(nextSearchRooms, rooms[roomIdx]...)
			}
		}

		nextRooms = nextSearchRooms
	}

	for i := 0; i < len(rooms); i++ {
		if _, ok := visitedRooms[i]; !ok {
			return false
		}
	}

	return true
}
