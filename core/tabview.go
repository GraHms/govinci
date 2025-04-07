package core

import "fmt"

type TabViewNode struct {
	SelectedIndex int
	OnTabChange   func(int)
	Tabs          []TabItem
	Content       []View
}

type TabItem struct {
	Label string
	Icon  string
}

type TabViewProp interface {
	Apply(*TabViewNode)
}

func TabView(props ...TabViewProp) View {
	return ComponentFunc(func(ctx *Context) *Node {
		node := &TabViewNode{}

		for _, p := range props {
			p.Apply(node)
		}

		// serialização dos tabs
		tabs := []map[string]string{}
		for _, t := range node.Tabs {
			tabs = append(tabs, map[string]string{
				"label": t.Label,
				"icon":  t.Icon,
			})
		}

		propMap := map[string]any{
			"selectedIndex": node.SelectedIndex,
			"tabs":          tabs,
		}
		if node.OnTabChange != nil {
			propMap["onTabChange"] = registerIntCallback(node.OnTabChange)
		}

		return &Node{
			Type:     "TabView",
			Props:    propMap,
			Children: renderAll(ctx, node.Content),
		}
	})
}

type tabViewFunc func(*TabViewNode)

func (f tabViewFunc) Apply(t *TabViewNode) { f(t) }

func SelectedIndex(i int) TabViewProp {
	return tabViewFunc(func(t *TabViewNode) {
		t.SelectedIndex = i
	})
}

func OnTabChange(fn func(int)) TabViewProp {
	return tabViewFunc(func(t *TabViewNode) {
		t.OnTabChange = fn
	})
}

func Tabs(tabs ...TabItem) TabViewProp {
	return tabViewFunc(func(t *TabViewNode) {
		t.Tabs = tabs
	})
}

func Content(views ...View) TabViewProp {
	return tabViewFunc(func(t *TabViewNode) {
		t.Content = views
	})
}

func Tab(label string, icon string) TabItem {
	return TabItem{Label: label, Icon: icon}
}

var (
	intCallbacks = map[string]func(int){}
	intCounter   int
)

func registerIntCallback(fn func(int)) string {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	id := fmt.Sprintf("int_cb_%d", intCounter)
	intCounter++
	intCallbacks[id] = fn
	return id
}

func TriggerIntCallback(id string, val int) {
	callbackMux.Lock()
	defer callbackMux.Unlock()

	if fn, ok := intCallbacks[id]; ok {
		fn(val)
	}
}
