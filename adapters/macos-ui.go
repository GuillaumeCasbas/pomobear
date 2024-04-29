package adapters

import (
	"github.com/guillaumecasbas/pomobear/usecases"
	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/objc"
)

// TODO: need to refactor after the repo system change
func StartUI(repo *PomodoroFileRepository) {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		item := appkit.StatusBar_SystemStatusBar().StatusItemWithLength(-1)
		objc.Retain(&item)
		item.Button().SetTitle("Bonjour le monde")

		go func() {
			for {
				u := usecases.NewPomodoroUsecases(repo)
				p := NewCmdPresenter()
				t, _ := u.Status()

				dispatch.MainQueue().DispatchAsync(func() {
					item.Button().SetTitle(p.DisplayStatus(t))
				})
			}
		}()

		itemQuit := appkit.NewMenuItem()
		itemQuit.SetTitle("Quit")
		itemQuit.SetAction(objc.Sel("terminate:"))

		menu := appkit.NewMenu()
		menu.AddItem(itemQuit)
		item.SetMenu(menu)
	})
}
