package com.govinci.app

import android.content.Context
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.FrameLayout
import android.widget.LinearLayout
import android.widget.TextView
import org.json.JSONArray
import org.json.JSONObject

class PatchRenderer(private val context: Context) {
    private val viewMap = mutableMapOf<String, View>()

    fun renderInitial(json: String, container: FrameLayout) {
        val node = JSONObject(json)
        val rootView = createView(node, "root")
        container.removeAllViews()
        container.addView(rootView)
    }

    fun applyPatches(json: String) {
        val patches = JSONArray(json)
        for (i in 0 until patches.length()) {
            val p = patches.getJSONObject(i)
            val type = p.getString("Type")
            val target = p.getString("TargetID")
            when (type) {
                "replace" -> {
                    val changes = p.getJSONObject("Changes")
                    val newView = createView(changes, target)
                    val old = viewMap[target] ?: continue
                    val parent = old.parent as? ViewGroup ?: continue
                    val index = parent.indexOfChild(old)
                    parent.removeViewAt(index)
                    parent.addView(newView, index)
                }
                "remove", "remove-child" -> {
                    val view = viewMap[target] ?: continue
                    val parent = view.parent as? ViewGroup ?: continue
                    parent.removeView(view)
                    viewMap.remove(target)
                }
                "add-child" -> {
                    val changes = p.getJSONObject("Changes")
                    val parent = viewMap[target] as? ViewGroup ?: continue
                    val path = "$target/${parent.childCount}"
                    val child = createView(changes, path)
                    parent.addView(child)
                }
                "update-props" -> {
                    val changes = p.getJSONObject("Changes")
                    val view = viewMap[target] ?: continue
                    updateProps(view, changes)
                }
                // update-style omitted for brevity
            }
        }
    }

    private fun createView(node: JSONObject, path: String): View {
        val type = node.getString("Type")
        val props = node.optJSONObject("Props")
        val view: View = when (type) {
            "Text" -> TextView(context).apply {
                text = props?.optString("content", "")
            }
            "Button" -> Button(context).apply {
                text = props?.optString("label", "")
                val cb = props?.optString("onClick")
                if (cb != null) {
                    setOnClickListener {
                        val patch = GovinciBridge.TriggerCallback(cb)
                        applyPatches(patch)
                    }
                }
            }
            "Column" -> LinearLayout(context).apply { orientation = LinearLayout.VERTICAL }
            "Row" -> LinearLayout(context).apply { orientation = LinearLayout.HORIZONTAL }
            else -> FrameLayout(context)
        }
        view.tag = path
        viewMap[path] = view
        val children = node.optJSONArray("Children")
        if (children != null && view is ViewGroup) {
            for (i in 0 until children.length()) {
                val child = children.getJSONObject(i)
                val childView = createView(child, "$path/$i")
                view.addView(childView)
            }
        }
        return view
    }

    private fun updateProps(view: View, props: JSONObject) {
        if (view is TextView) {
            props.optString("content")?.let { view.text = it }
        }
        if (view is Button) {
            props.optString("label")?.let { view.text = it }
            props.optString("onClick")?.let { id ->
                view.setOnClickListener {
                    val patch = GovinciBridge.TriggerCallback(id)
                    applyPatches(patch)
                }
            }
        }
    }
}
