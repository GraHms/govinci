package core

func Image(src string, styleProps ...StyleProp) View {
	return ComponentFunc(func(ctx *Context) *Node {
		base := ctx.Theme().Components.Camera
		style := &base
		for _, sp := range styleProps {
			sp.Apply(style)
		}

		return &Node{
			Type: "Image",
			Props: map[string]any{
				"src": src,
			},
			Style: style,
		}
	})
}
