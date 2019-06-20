package day08

import (
	"strconv"
	"strings"
)

type node struct {
	Children []*node
	Metadata []int
}

// Part1 returns the sum of all the metadata entries in all nodes of the tree.
func Part1(input string) (string, error) {
	nums, err := parseInput(input)
	if err != nil {
		return "", err
	}
	tree, _ := makeTree(nums, 0)
	sum := sumTree(tree)
	return strconv.Itoa(sum), nil
}

// Part2 returns the
func Part2(input string) (string, error) {
	nums, err := parseInput(input)
	if err != nil {
		return "", err
	}
	tree, _ := makeTree(nums, 0)
	sum := sumTreeByIndex(tree)
	return strconv.Itoa(sum), nil
}

func parseInput(input string) ([]int, error) {
	rawNums := strings.Split(strings.TrimSpace(input), " ")
	nums := make([]int, len(rawNums))
	for i, rawNum := range rawNums {
		num, err := strconv.Atoi(rawNum)
		if err != nil {
			return nums, err
		}
		nums[i] = num
	}
	return nums, nil
}

// makeTree assumes a valid input array and will panic if the input does not
// form a valid tree.
func makeTree(nums []int, startIndex int) (root *node, nextIndex int) {
	if startIndex >= len(nums) {
		return
	}
	childCount := nums[startIndex]
	metadataCount := nums[startIndex+1]
	root = &node{
		Children: make([]*node, childCount),
		Metadata: make([]int, metadataCount),
	}

	nextIndex = startIndex + 2
	var child *node
	for i := range root.Children {
		child, nextIndex = makeTree(nums, nextIndex)
		root.Children[i] = child
	}
	for i := range root.Metadata {
		root.Metadata[i] = nums[nextIndex]
		nextIndex++
	}
	return
}

func sumTree(root *node) (sum int) {
	for _, child := range root.Children {
		sum += sumTree(child)
	}
	for _, num := range root.Metadata {
		sum += num
	}
	return
}

func sumTreeByIndex(root *node) (sum int) {
	if len(root.Children) == 0 {
		return sumTree(root)
	}

	for _, childNum := range root.Metadata {
		childIndex := childNum - 1
		if 0 <= childIndex && childIndex < len(root.Children) {
			sum += sumTreeByIndex(root.Children[childIndex])
		}
	}
	return
}
