package main

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

var Tree = &BinaryNode{
	Value: 20,
	Right: &BinaryNode{
		Value: 50,
		Right: &BinaryNode{
			Value: 100,
		},
		Left: &BinaryNode{
			Value: 30,
			Right: &BinaryNode{
				Value: 45,
			},
			Left: &BinaryNode{
				Value: 29,
			},
		},
	},
	Left: &BinaryNode{
		Value: 10,
		Right: &BinaryNode{
			Value: 15,
		},
		Left: &BinaryNode{
			Value: 5,
			Right: &BinaryNode{
				Value: 7,
			},
		},
	},
}

var Tree2 = &BinaryNode{
	Value: 20,
	Right: &BinaryNode{
		Value: 50,
		Left: &BinaryNode{
			Value: 30,
			Right: &BinaryNode{
				Value: 45,
				Right: &BinaryNode{
					Value: 49,
				},
			},
			Left: &BinaryNode{
				Value: 29,
				Left: &BinaryNode{
					Value: 21,
				},
			},
		},
	},
	Left: &BinaryNode{
		Value: 10,
		Right: &BinaryNode{
			Value: 15,
		},
		Left: &BinaryNode{
			Value: 5,
			Right: &BinaryNode{
				Value: 7,
			},
		},
	},
}
