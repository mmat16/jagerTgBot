package telegram

const msgHelp = `I can save and keep you pages. Also i can offer you them to read.

In order to save the page, just send me a link to it.

In order to get a random page from you list, send me command /rnd.

Caution! After /rnd the sending page would be removed from your list`

const msgHello = "Hi there!\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command"
	msgNoSavedPages   = "You don't have any saved pages. Yet (maybe)\nor already.."
	msgSaved          = "Saved!"
	msgAlreadyExists  = "You have already saved this page! Don't you trust my storage?.."
)
