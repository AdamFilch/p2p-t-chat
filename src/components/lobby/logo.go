package components_lobby

import "github.com/rivo/tview"

func LobbyLogo() *tview.TextView {
	// 		logo := `
	// 	s                                                    s
	// 	:8                         .uef^"                    :8
	//    .88                       :d88E                      .88
	//   :888ooo                .   '888E             u       :888ooo
	// -*8888888           .udR88N   888E .z8k     us888u.  -*8888888
	//   8888             <888'888k  888E~?888L .@88 "8888"   8888
	//   8888             9888 'Y"   888E  888E 9888  9888    8888
	//   8888             9888       888E  888E 9888  9888    8888
	//  .8888Lu= 88888888 9888       888E  888E 9888  9888   .8888Lu=
	//  ^%888*   88888888 ?8888u../  888E  888E 9888  9888   ^%888*
	//    'Y"              "8888P'  m888N= 888> "888*""888"    'Y"
	// 					  "P'     'Y"   888   ^Y"   ^Y'
	// 								   J88"
	// 								   @%
	// 								 :"
	// `

	logo2 := `
.----------------.  .----------------.  .----------------.  .----------------.  .----------------.  .----------------.
| .--------------. || .--------------. || .--------------. || .--------------. || .--------------. || .--------------. |
| |  _________   | || |              | || |     ______   | || |  ____  ____  | || |      __      | || |  _________   | |
| | |  _   _  |  | || |              | || |   .' ___  |  | || | |_   ||   _| | || |     /  \     | || | |  _   _  |  | |
| | |_/ | | \_|  | || |    ______    | || |  / .'   \_|  | || |   | |__| |   | || |    / /\ \    | || | |_/ | | \_|  | |
| |     | |      | || |   |______|   | || |  | |         | || |   |  __  |   | || |   / ____ \   | || |     | |      | |
| |    _| |_     | || |              | || |  \ '.___.'\  | || |  _| |  | |_  | || | _/ /    \ \_ | || |    _| |_     | |
| |   |_____|    | || |              | || |   '._____.'  | || | |____||____| | || ||____|  |____|| || |   |_____|    | |
| |              | || |              | || |              | || |              | || |              | || |              | |
| '--------------' || '--------------' || '--------------' || '--------------' || '--------------' || '--------------' |
 '----------------'  '----------------'  '----------------'  '----------------'  '----------------'  '----------------'`

// 	logo3 := `
// _________                   ________      ___  ___      ________      _________   
// |\___   ___\                |\   ____\    |\  \|\  \    |\   __  \    |\___   ___\ 
// \|___ \  \_|  ____________  \ \  \___|    \ \  \\\  \   \ \  \|\  \   \|___ \  \_| 
//      \ \  \  |\____________\ \ \  \        \ \   __  \   \ \   __  \       \ \  \  
//       \ \  \ \|____________|  \ \  \____    \ \  \ \  \   \ \  \ \  \       \ \  \ 
//        \ \__\                  \ \_______\   \ \__\ \__\   \ \__\ \__\       \ \__\
//         \|__|                   \|_______|    \|__|\|__|    \|__|\|__|        \|__|`

	// logo4 := `

	// ████████╗       ██████╗██╗  ██╗ █████╗ ████████╗
	// ╚══██╔══╝      ██╔════╝██║  ██║██╔══██╗╚══██╔══╝
	//    ██║   █████╗██║     ███████║███████║   ██║
	//    ██║   ╚════╝██║     ██╔══██║██╔══██║   ██║
	//    ██║         ╚██████╗██║  ██║██║  ██║   ██║
	//    ╚═╝          ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝`

	// logo5 := `
	// ▄▄▄█████▓ ▄████▄   ██░ ██  ▄▄▄      ▄▄▄█████▓
	// ▓  ██▒ ▓▒▒██▀ ▀█  ▓██░ ██▒▒████▄    ▓  ██▒ ▓▒
	// ▒ ▓██░ ▒░▒▓█    ▄ ▒██▀▀██░▒██  ▀█▄  ▒ ▓██░ ▒░
	// ░ ▓██▓ ░ ▒▓▓▄ ▄██▒░▓█ ░██ ░██▄▄▄▄██ ░ ▓██▓ ░
	//   ▒██▒ ░ ▒ ▓███▀ ░░▓█▒░██▓ ▓█   ▓██▒  ▒██▒ ░
	//   ▒ ░░   ░ ░▒ ▒  ░ ▒ ░░▒░▒ ▒▒   ▓▒█░  ▒ ░░
	//     ░      ░  ▒    ▒ ░▒░ ░  ▒   ▒▒ ░    ░
	//   ░      ░         ░  ░░ ░  ░   ▒     ░
	//          ░ ░       ░  ░  ░      ░  ░
	//          ░                                   `

	textView := tview.NewTextView().SetText(logo2).SetTextAlign(tview.AlignCenter)

	return textView

}
