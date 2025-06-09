package com.govinci.app

object GovinciBridge {
    init {
        System.loadLibrary("govinci")
    }

    external fun InitApp()
    external fun RenderInitial(): String
    external fun TriggerCallback(id: String): String
    external fun TriggerTextCallback(id: String, value: String): String
}
