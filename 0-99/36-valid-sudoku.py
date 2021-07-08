# coding: utf-8

class Solution:
	# @param board, a 9x9 2D array
	# @return a boolean
	def isValidSudoku(self, board):
		row = [0]*9
		numl = []
		for i in range(9):
			numl.append(row[:])

		sub = [(3*i, 3*j) for i in range(3) for j in range(3)]	#每个小方块的右上角元素

		for i in range(9):
			for j in range(9):
				if board[i][j] == '.':
					numl[i][j] = 0
				else:
					numl[i][j] = int(board[i][j])

		for i in range(9):										#对每一行进行判定
			visited = [False]*10
			for j in range(9):
				if numl[i][j] > 0:
					if visited[numl[i][j]]:
						return False
					else:
						visited[numl[i][j]] = True

		for j in range(9):										#对每一列进行判定
			visited = [False]*10
			for i in range(9):
				if numl[i][j] > 0:
					if visited[numl[i][j]]:
						return False
					else:
						visited[numl[i][j]] = True

		for x, y in sub:										#对每一小方块进行判定
			visited = [False]*10
			for i in range(3):
				for j in range(3):
					if numl[x+i][y+j] > 0:
						if visited[numl[x+i][y+j]]:
							return False
						else:
							visited[numl[x+i][y+j]] = True
							
		return True

a = Solution()
a.isValidSudoku(1)