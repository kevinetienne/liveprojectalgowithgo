package main

import "fmt"

const numDisks = 3

func push(post []int, disk int) []int {
	return append([]int{disk}, post...)
}

func pop(post []int) (int, []int) {
	return post[0], post[1:]
}

func moveDisk(posts [][]int, from, to int) {
	disk, post := pop(posts[from])
	posts[from] = post
	posts[to] = push(posts[to], disk)
}

func drawPosts(posts [][]int) {
	pp := make([][]int, 3)
	for i := range posts {
		copy(pp[i], posts[i])

	}
	copy(pp, posts)

	for i, post := range posts {
		for j := 3 - len(post); j > 0; j-- {
			pp[i] = push(pp[i], 0)
		}
	}

	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			fmt.Printf("%2d", pp[r][c])
		}
		fmt.Println()
	}
	fmt.Println("-------")
}

// Move the disks from from_post to to_post
// using temp_post as temporary storage.
func moveDisks(posts [][]int, numToMove, from, to, tmp int) {
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, from, tmp, to)
	}

	moveDisk(posts, from, to)
	drawPosts(posts)

	if numToMove > 1 {
		moveDisks(posts, numToMove-1, tmp, to, from)
	}
}

func main() {
	// Make three posts.
	posts := [][]int{}

	// Push the disks onto post 0 biggest first.
	posts = append(posts, []int{})
	for disk := numDisks; disk > 0; disk-- {
		posts[0] = push(posts[0], disk)
	}

	// Make the other posts empty.
	for p := 1; p < 3; p++ {
		posts = append(posts, []int{})
	}

	// Draw the initial setup.
	drawPosts(posts)

	// Move the disks.
	moveDisks(posts, numDisks, 0, 1, 2)
}
