package treenode


var (
	ExampleTree1 = &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5},
					{Val: 6},
				},
			},
			{Val: 2},
			{Val: 4},
		},
	}

	ExampleTree2 = &Node{
		Val: 1,
		Children: []*Node{
			{Val: 2},
			{
				Val: 3,
				Children: []*Node{
					{Val: 6},
					{Val: 7, Children: []*Node{
						{Val: 11, Children: []*Node{
							{Val: 14},
						}},
					}},
				},
			},
			{Val: 4, Children: []*Node{
				{Val: 8, Children: []*Node{
					{Val: 12},
				}},
			}},
			{Val: 5, Children: []*Node{
				{Val: 9, Children: []*Node{
					{Val: 13},
				}},
				{Val: 10},
			}},
		},
	}
)