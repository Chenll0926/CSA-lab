package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	var aliveCells []cell
	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {

			aroundCells := []byte{
				world[(i+p.imageHeight-1)%p.imageHeight][(j+p.imageWidth-1)%p.imageWidth],
				world[(i+p.imageHeight-1)%p.imageHeight][(j+p.imageWidth)%p.imageWidth],
				world[(i+p.imageHeight-1)%p.imageHeight][(j+p.imageWidth+1)%p.imageWidth],
				world[(i+p.imageHeight)%p.imageHeight][(j+p.imageWidth-1)%p.imageWidth],
				world[(i+p.imageHeight)%p.imageHeight][(j+p.imageWidth+1)%p.imageWidth],
				world[(i+p.imageHeight+1)%p.imageHeight][(j+p.imageWidth-1)%p.imageWidth],
				world[(i+p.imageHeight+1)%p.imageHeight][(j+p.imageWidth)%p.imageWidth],
				world[(i+p.imageHeight+1)%p.imageHeight][(j+p.imageWidth+1)%p.imageWidth],
			}

			numberOfAlive := 0
			for _, val := range aroundCells {
				if val == 255 {
					numberOfAlive++
				}
			}
			if world[i][j] == 255 {
				if numberOfAlive == 2 || numberOfAlive == 3 {
					//world[i][j] = 255
					aliveCells = append(aliveCells, cell{i, j})
				}
			} else if world[i][j] == 0 {
				if numberOfAlive == 3 {
					//world[i][j] = 255
					aliveCells = append(aliveCells, cell{i, j})
				}
			}
		}
	}

	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {
			world[i][j] = 0
		}
	}

	for _, aliveCell := range aliveCells {
		world[aliveCell.x][aliveCell.y] = 255
	}

	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var aliveCells []cell

	for i := 0; i < p.imageHeight; i++ {
		for j := 0; j < p.imageWidth; j++ {

			if world[i][j] == 255 {
				aliveCells = append(aliveCells, cell{j, i})
			}
		}
	}

	return aliveCells
}
