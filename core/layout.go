package core

type PropsAndChildren any

func Row(stylePropsAndChildren ...PropsAndChildren) View {
	return ComponentFunc(func(ctx *Context) *Node {
		base := ctx.Theme().Components.Row
		style := &base
		var children []View

		for _, item := range stylePropsAndChildren {
			switch v := item.(type) {
			case StyleProp:
				v.Apply(style)
			case View:
				children = append(children, v)
			}
		}

		return &Node{
			Type:     "Row",
			Style:    style,
			Children: renderAll(ctx, children),
		}
	})
}

func Card(stylePropsAndChildren ...PropsAndChildren) View {
	return ComponentFunc(func(ctx *Context) *Node {
		base := ctx.Theme().Components.Card
		style := &base
		var children []View

		for _, item := range stylePropsAndChildren {
			switch v := item.(type) {
			case StyleProp:
				v.Apply(style)
			case View:
				children = append(children, v)
			}
		}

		return &Node{
			Type:     "Card",
			Style:    style,
			Children: renderAll(ctx, children),
		}
	})
}

func Spacer(size int) View {
	return ComponentFunc(func(ctx *Context) *Node {
		return &Node{
			Type: "Spacer",
			Props: map[string]any{
				"size": size,
			},
		}
	})
}

func Scroll(children ...View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		var nodes []*Node
		for _, child := range children {
			nodes = append(nodes, child.Render(ctx))
		}
		return &Node{
			Type:     "Scroll",
			Props:    map[string]any{},
			Children: nodes,
		}
	})
}

func SafeArea(child View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		return &Node{
			Type:     "SafeArea",
			Props:    map[string]any{},
			Children: []*Node{child.Render(ctx)},
		}
	})
}

func Fragment(children ...View) View {
	return ComponentFunc(func(ctx *Context) *Node {
		if len(children) == 1 {
			return children[0].Render(ctx)
		}
		return &Node{
			Type:     "Fragment",
			Children: renderAll(ctx, children),
		}
	})
}

func Column(stylePropsAndChildren ...PropsAndChildren) View {
	return ComponentFunc(func(ctx *Context) *Node {
		base := ctx.Theme().Components.Column
		style := &base
		var children []View

		for _, item := range stylePropsAndChildren {
			switch v := item.(type) {
			case StyleProp:
				v.Apply(style)
			case View:
				children = append(children, v)
			}
		}

		return &Node{
			Type:     "Column",
			Style:    style,
			Children: renderAll(ctx, children),
		}
	})
}

func renderAll(ctx *Context, views []View) []*Node {
	var out []*Node
	for _, v := range views {
		out = append(out, v.Render(ctx))
	}
	return out
}
