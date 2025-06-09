package com.govinci.app

import android.os.Bundle
import android.widget.FrameLayout
import androidx.appcompat.app.AppCompatActivity

class MainActivity : AppCompatActivity() {
    private lateinit var root: FrameLayout
    private lateinit var renderer: PatchRenderer

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        root = FrameLayout(this)
        setContentView(root)
        renderer = PatchRenderer(this)

        GovinciBridge.InitApp()
        val initial = GovinciBridge.RenderInitial()
        renderer.renderInitial(initial, root)
    }
}
