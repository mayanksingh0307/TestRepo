package co.bito.intellij.codeCompletion.toolwindow

import com.intellij.openapi.editor.RangeMarker
import co.bito.intellij.codeCompletion.model.InvocationContext
import co.bito.intellij.codeCompletion.model.SessionContext
import co.bito.intellij.codeCompletion.popup.UserActionListener

class CodeReferenceActionListener : UserActionListener {
    override fun afterAccept(states: InvocationContext, sessionContext: SessionContext, rangeMarker: RangeMarker) {
        val (project, editor) = states.requestContext
        val manager = CodeReferenceManager.getInstance(project)
        manager.insertCodeReference(states, sessionContext.selectedIndex)
        manager.addListeners(editor)
    }
}
