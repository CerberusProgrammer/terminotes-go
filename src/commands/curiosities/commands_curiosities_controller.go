package curiosities

import "math/rand"

var Haikus = []string{
	"\nSilent notes whisper,\nIn the terminal they live,\nThoughts forever saved.\n",
	"\nCode and notes unite,\nIn the quiet terminal,\nIdeas take flight.\n",
	"\nLocal notes reside,\nNo cloud to steal them away,\nPrivacy preserved.\n",
	"\nQuickly jot it down,\nIn the terminal it stays,\nFast and simple way.\n",
	"\nTablet, phone, or key,\nNotes are stored locally,\nFreedom in each line.\n",
	"\nSimple notes are best,\nIn the terminal they rest,\nIdeas at their best.\n",
}

func HaikuCommand(args []string) {
	randIndex := rand.Intn(len(Haikus))
	println(Haikus[randIndex])
}

func LogoCommand(args []string) {
	println(`
  _______  ______  _____   __  __  _____  _   _   ____  _______  ______   _____ 
 |__   __||  ____||  __ \ |  \/  ||_   _|| \ | | / __ \|__   __||  ____| / ____|
    | |   | |__   | |__) || \  / |  | |  |  \| || |  | |  | |   | |__   | (___  
    | |   |  __|  |  _  / | |\/| |  | |  | .   || |  | |  | |   |  __|   \___ \ 
    | |   | |____ | | \ \ | |  | | _| |_ | |\  || |__| |  | |   | |____  ____) |
    |_|   |______||_|  \_\|_|  |_||_____||_| \_| \____/   |_|   |______||_____/ 
	`)
}
